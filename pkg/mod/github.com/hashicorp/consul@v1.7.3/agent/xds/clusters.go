package xds

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	envoy "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoyauth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	envoycluster "github.com/envoyproxy/go-control-plane/envoy/api/v2/cluster"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyendpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	"github.com/hashicorp/consul/agent/connect"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
)

// clustersFromSnapshot returns the xDS API representation of the "clusters" in the snapshot.
func (s *Server) clustersFromSnapshot(cfgSnap *proxycfg.ConfigSnapshot, token string) ([]proto.Message, error) {
	if cfgSnap == nil {
		return nil, errors.New("nil config given")
	}

	switch cfgSnap.Kind {
	case structs.ServiceKindConnectProxy:
		return s.clustersFromSnapshotConnectProxy(cfgSnap, token)
	case structs.ServiceKindMeshGateway:
		return s.clustersFromSnapshotMeshGateway(cfgSnap, token)
	default:
		return nil, fmt.Errorf("Invalid service kind: %v", cfgSnap.Kind)
	}
}

// clustersFromSnapshot returns the xDS API representation of the "clusters"
// (upstreams) in the snapshot.
func (s *Server) clustersFromSnapshotConnectProxy(cfgSnap *proxycfg.ConfigSnapshot, token string) ([]proto.Message, error) {
	// TODO(rb): this sizing is a low bound.
	clusters := make([]proto.Message, 0, len(cfgSnap.Proxy.Upstreams)+1)

	// Include the "app" cluster for the public listener
	appCluster, err := s.makeAppCluster(cfgSnap, LocalAppClusterName, "", cfgSnap.Proxy.LocalServicePort)
	if err != nil {
		return nil, err
	}

	clusters = append(clusters, appCluster)

	for _, u := range cfgSnap.Proxy.Upstreams {
		id := u.Identifier()

		if u.DestinationType == structs.UpstreamDestTypePreparedQuery {
			upstreamCluster, err := s.makeUpstreamClusterForPreparedQuery(u, cfgSnap)
			if err != nil {
				return nil, err
			}
			clusters = append(clusters, upstreamCluster)

		} else {
			chain := cfgSnap.ConnectProxy.DiscoveryChain[id]
			upstreamClusters, err := s.makeUpstreamClustersForDiscoveryChain(u, chain, cfgSnap)
			if err != nil {
				return nil, err
			}

			for _, cluster := range upstreamClusters {
				clusters = append(clusters, cluster)
			}
		}
	}

	cfgSnap.Proxy.Expose.Finalize()
	paths := cfgSnap.Proxy.Expose.Paths

	// Add service health checks to the list of paths to create clusters for if needed
	if cfgSnap.Proxy.Expose.Checks {
		psid := structs.NewServiceID(cfgSnap.Proxy.DestinationServiceID, &cfgSnap.ProxyID.EnterpriseMeta)
		for _, check := range s.CheckFetcher.ServiceHTTPBasedChecks(psid) {
			p, err := parseCheckPath(check)
			if err != nil {
				s.Logger.Warn("failed to create cluster for", "check", check.CheckID, "error", err)
				continue
			}
			paths = append(paths, p)
		}
	}

	// Create a new cluster if we need to expose a port that is different from the service port
	for _, path := range paths {
		if path.LocalPathPort == cfgSnap.Proxy.LocalServicePort {
			continue
		}
		c, err := s.makeAppCluster(cfgSnap, makeExposeClusterName(path.LocalPathPort), path.Protocol, path.LocalPathPort)
		if err != nil {
			s.Logger.Warn("failed to make local cluster", "path", path.Path, "error", err)
			continue
		}
		clusters = append(clusters, c)
	}
	return clusters, nil
}

func makeExposeClusterName(destinationPort int) string {
	return fmt.Sprintf("exposed_cluster_%d", destinationPort)
}

// clustersFromSnapshotMeshGateway returns the xDS API representation of the "clusters"
// for a mesh gateway. This will include 1 cluster per remote datacenter as well as
// 1 cluster for each service subset.
func (s *Server) clustersFromSnapshotMeshGateway(cfgSnap *proxycfg.ConfigSnapshot, token string) ([]proto.Message, error) {
	// 1 cluster per remote dc + 1 cluster per local service (this is a lower bound - all subset specific clusters will be appended)
	clusters := make([]proto.Message, 0, len(cfgSnap.MeshGateway.GatewayGroups)+len(cfgSnap.MeshGateway.ServiceGroups))

	// generate the remote dc clusters
	for dc, _ := range cfgSnap.MeshGateway.GatewayGroups {
		clusterName := connect.DatacenterSNI(dc, cfgSnap.Roots.TrustDomain)

		cluster, err := s.makeMeshGatewayCluster(clusterName, cfgSnap)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, cluster)
	}

	// generate the per-service clusters
	for svc, _ := range cfgSnap.MeshGateway.ServiceGroups {
		clusterName := connect.ServiceSNI(svc.ID, "", svc.NamespaceOrDefault(), cfgSnap.Datacenter, cfgSnap.Roots.TrustDomain)
		resolver, hasResolver := cfgSnap.MeshGateway.ServiceResolvers[svc]

		// Create the cluster for default/unnamed services
		var cluster *envoy.Cluster
		var err error
		if hasResolver {
			cluster, err = s.makeMeshGatewayClusterWithConnectTimeout(clusterName, cfgSnap, resolver.ConnectTimeout)
		} else {
			cluster, err = s.makeMeshGatewayCluster(clusterName, cfgSnap)
		}
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, cluster)

		// if there is a service-resolver for this service then also setup subset clusters for it
		if hasResolver {
			// generate 1 cluster for each service subset
			for subsetName := range resolver.Subsets {
				clusterName := connect.ServiceSNI(svc.ID, subsetName, svc.NamespaceOrDefault(), cfgSnap.Datacenter, cfgSnap.Roots.TrustDomain)

				cluster, err := s.makeMeshGatewayClusterWithConnectTimeout(clusterName, cfgSnap, resolver.ConnectTimeout)
				if err != nil {
					return nil, err
				}
				clusters = append(clusters, cluster)
			}
		}
	}

	return clusters, nil
}

func (s *Server) makeAppCluster(cfgSnap *proxycfg.ConfigSnapshot, name, pathProtocol string, port int) (*envoy.Cluster, error) {
	var c *envoy.Cluster
	var err error

	cfg, err := ParseProxyConfig(cfgSnap.Proxy.Config)
	if err != nil {
		// Don't hard fail on a config typo, just warn. The parse func returns
		// default config if there is an error so it's safe to continue.
		s.Logger.Warn("failed to parse Connect.Proxy.Config", "error", err)
	}

	// If we have overridden local cluster config try to parse it into an Envoy cluster
	if cfg.LocalClusterJSON != "" {
		return makeClusterFromUserConfig(cfg.LocalClusterJSON)
	}

	addr := cfgSnap.Proxy.LocalServiceAddress
	if addr == "" {
		addr = "127.0.0.1"
	}
	c = &envoy.Cluster{
		Name:                 name,
		ConnectTimeout:       time.Duration(cfg.LocalConnectTimeoutMs) * time.Millisecond,
		ClusterDiscoveryType: &envoy.Cluster_Type{Type: envoy.Cluster_STATIC},
		LoadAssignment: &envoy.ClusterLoadAssignment{
			ClusterName: name,
			Endpoints: []envoyendpoint.LocalityLbEndpoints{
				{
					LbEndpoints: []envoyendpoint.LbEndpoint{
						makeEndpoint(name,
							addr,
							port),
					},
				},
			},
		},
	}
	if cfg.Protocol == "http2" || cfg.Protocol == "grpc" || pathProtocol == "http2" {
		c.Http2ProtocolOptions = &envoycore.Http2ProtocolOptions{}
	}

	return c, err
}

func (s *Server) makeUpstreamClusterForPreparedQuery(upstream structs.Upstream, cfgSnap *proxycfg.ConfigSnapshot) (*envoy.Cluster, error) {
	var c *envoy.Cluster
	var err error

	dc := upstream.Datacenter
	if dc == "" {
		dc = cfgSnap.Datacenter
	}
	sni := connect.UpstreamSNI(&upstream, "", dc, cfgSnap.Roots.TrustDomain)

	cfg, err := ParseUpstreamConfig(upstream.Config)
	if err != nil {
		// Don't hard fail on a config typo, just warn. The parse func returns
		// default config if there is an error so it's safe to continue.
		s.Logger.Warn("failed to parse", "upstream", upstream.Identifier(), "error", err)
	}
	if cfg.ClusterJSON != "" {
		c, err = makeClusterFromUserConfig(cfg.ClusterJSON)
		if err != nil {
			return c, err
		}
		// In the happy path don't return yet as we need to inject TLS config still.
	}

	if c == nil {
		c = &envoy.Cluster{
			Name:                 sni,
			ConnectTimeout:       time.Duration(cfg.ConnectTimeoutMs) * time.Millisecond,
			ClusterDiscoveryType: &envoy.Cluster_Type{Type: envoy.Cluster_EDS},
			EdsClusterConfig: &envoy.Cluster_EdsClusterConfig{
				EdsConfig: &envoycore.ConfigSource{
					ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
						Ads: &envoycore.AggregatedConfigSource{},
					},
				},
			},
			CircuitBreakers: &envoycluster.CircuitBreakers{
				Thresholds: makeThresholdsIfNeeded(cfg.Limits),
			},
			// Having an empty config enables outlier detection with default config.
			OutlierDetection: &envoycluster.OutlierDetection{},
		}
		if cfg.Protocol == "http2" || cfg.Protocol == "grpc" {
			c.Http2ProtocolOptions = &envoycore.Http2ProtocolOptions{}
		}
	}

	// Enable TLS upstream with the configured client certificate.
	c.TlsContext = &envoyauth.UpstreamTlsContext{
		CommonTlsContext: makeCommonTLSContext(cfgSnap),
		Sni:              sni,
	}

	return c, nil
}

func (s *Server) makeUpstreamClustersForDiscoveryChain(
	upstream structs.Upstream,
	chain *structs.CompiledDiscoveryChain,
	cfgSnap *proxycfg.ConfigSnapshot,
) ([]*envoy.Cluster, error) {
	if chain == nil {
		return nil, fmt.Errorf("cannot create upstream cluster without discovery chain for %s", upstream.Identifier())
	}

	cfg, err := ParseUpstreamConfigNoDefaults(upstream.Config)
	if err != nil {
		// Don't hard fail on a config typo, just warn. The parse func returns
		// default config if there is an error so it's safe to continue.
		s.Logger.Warn("failed to parse", "upstream", upstream.Identifier(),
			"error", err)
	}

	var escapeHatchCluster *envoy.Cluster
	if cfg.ClusterJSON != "" {
		if chain.IsDefault() {
			// If you haven't done anything to setup the discovery chain, then
			// you can use the envoy_cluster_json escape hatch.
			escapeHatchCluster, err = makeClusterFromUserConfig(cfg.ClusterJSON)
			if err != nil {
				return nil, err
			}
		} else {
			s.Logger.Warn("ignoring escape hatch setting, because a discovery chain is configued for",
				"discovery chain", chain.ServiceName, "upstream", upstream.Identifier(),
				"envoy_cluster_json", chain.ServiceName)
		}
	}

	id := upstream.Identifier()
	chainEndpointMap, ok := cfgSnap.ConnectProxy.WatchedUpstreamEndpoints[id]
	if !ok {
		// this should not happen
		return nil, fmt.Errorf("no endpoint map for upstream %q", id)
	}

	var out []*envoy.Cluster

	for _, node := range chain.Nodes {
		if node.Type != structs.DiscoveryGraphNodeTypeResolver {
			continue
		}
		failover := node.Resolver.Failover
		targetID := node.Resolver.Target

		target := chain.Targets[targetID]

		// Determine if we have to generate the entire cluster differently.
		failoverThroughMeshGateway := chain.WillFailoverThroughMeshGateway(node)

		sni := target.SNI
		clusterName := CustomizeClusterName(target.Name, chain)

		if failoverThroughMeshGateway {
			actualTargetID := firstHealthyTarget(
				chain.Targets,
				chainEndpointMap,
				targetID,
				failover.Targets,
			)

			if actualTargetID != targetID {
				actualTarget := chain.Targets[actualTargetID]
				sni = actualTarget.SNI
			}
		}

		s.Logger.Debug("generating cluster for", "cluster", clusterName)
		c := &envoy.Cluster{
			Name:                 clusterName,
			AltStatName:          clusterName,
			ConnectTimeout:       node.Resolver.ConnectTimeout,
			ClusterDiscoveryType: &envoy.Cluster_Type{Type: envoy.Cluster_EDS},
			CommonLbConfig: &envoy.Cluster_CommonLbConfig{
				HealthyPanicThreshold: &envoytype.Percent{
					Value: 0, // disable panic threshold
				},
			},
			EdsClusterConfig: &envoy.Cluster_EdsClusterConfig{
				EdsConfig: &envoycore.ConfigSource{
					ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
						Ads: &envoycore.AggregatedConfigSource{},
					},
				},
			},
			CircuitBreakers: &envoycluster.CircuitBreakers{
				Thresholds: makeThresholdsIfNeeded(cfg.Limits),
			},
			// Having an empty config enables outlier detection with default config.
			OutlierDetection: &envoycluster.OutlierDetection{},
		}

		proto := cfg.Protocol
		if proto == "" {
			proto = chain.Protocol
		}

		if proto == "" {
			proto = "tcp"
		}

		if proto == "http2" || proto == "grpc" {
			c.Http2ProtocolOptions = &envoycore.Http2ProtocolOptions{}
		}

		// Enable TLS upstream with the configured client certificate.
		c.TlsContext = &envoyauth.UpstreamTlsContext{
			CommonTlsContext: makeCommonTLSContext(cfgSnap),
			Sni:              sni,
		}

		out = append(out, c)
	}

	if escapeHatchCluster != nil {
		if len(out) != 1 {
			return nil, fmt.Errorf("cannot inject escape hatch cluster when discovery chain had no nodes")
		}
		defaultCluster := out[0]

		// Overlay what the user provided.
		escapeHatchCluster.TlsContext = defaultCluster.TlsContext

		out = []*envoy.Cluster{escapeHatchCluster}
	}

	return out, nil
}

// makeClusterFromUserConfig returns the listener config decoded from an
// arbitrary proto3 json format string or an error if it's invalid.
//
// For now we only support embedding in JSON strings because of the hcl parsing
// pain (see config.go comment above call to PatchSliceOfMaps). Until we
// refactor config parser a _lot_ user's opaque config that contains arrays will
// be mangled. We could actually fix that up in mapstructure which knows the
// type of the target so could resolve the slices to singletons unambiguously
// and it would work for us here... but we still have the problem that the
// config would render incorrectly in general in our HTTP API responses so we
// really need to fix it "properly".
//
// When we do that we can support just nesting the config directly into the
// JSON/hcl naturally but this is a stop-gap that gets us an escape hatch
// immediately. It's also probably not a bad thing to support long-term since
// any config generated by other systems will likely be in canonical protobuf
// from rather than our slight variant in JSON/hcl.
func makeClusterFromUserConfig(configJSON string) (*envoy.Cluster, error) {
	var jsonFields map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(configJSON), &jsonFields); err != nil {
		fmt.Println("Custom error", err, configJSON)
		return nil, err
	}

	var c envoy.Cluster

	if _, ok := jsonFields["@type"]; ok {
		// Type field is present so decode it as a types.Any
		var any types.Any
		err := jsonpb.UnmarshalString(configJSON, &any)
		if err != nil {
			return nil, err
		}
		// And then unmarshal the listener again...
		err = proto.Unmarshal(any.Value, &c)
		if err != nil {
			return nil, err
		}
		return &c, err
	}

	// No @type so try decoding as a straight listener.
	err := jsonpb.UnmarshalString(configJSON, &c)
	return &c, err
}

func (s *Server) makeMeshGatewayCluster(clusterName string, cfgSnap *proxycfg.ConfigSnapshot) (*envoy.Cluster, error) {
	return s.makeMeshGatewayClusterWithConnectTimeout(clusterName, cfgSnap, 0)
}

// makeMeshGatewayClusterWithConnectTimeout initializes a mesh gateway cluster
// with the specified connect timeout. If the timeout is 0, the connect timeout
// defaults to use the mesh gateway timeout.
func (s *Server) makeMeshGatewayClusterWithConnectTimeout(clusterName string, cfgSnap *proxycfg.ConfigSnapshot,
	connectTimeout time.Duration) (*envoy.Cluster, error) {
	cfg, err := ParseMeshGatewayConfig(cfgSnap.Proxy.Config)
	if err != nil {
		// Don't hard fail on a config typo, just warn. The parse func returns
		// default config if there is an error so it's safe to continue.
		s.Logger.Warn("failed to parse mesh gateway config", "error", err)
	}

	if connectTimeout <= 0 {
		connectTimeout = time.Duration(cfg.ConnectTimeoutMs) * time.Millisecond
	}

	return &envoy.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       connectTimeout,
		ClusterDiscoveryType: &envoy.Cluster_Type{Type: envoy.Cluster_EDS},
		EdsClusterConfig: &envoy.Cluster_EdsClusterConfig{
			EdsConfig: &envoycore.ConfigSource{
				ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
					Ads: &envoycore.AggregatedConfigSource{},
				},
			},
		},
		// Having an empty config enables outlier detection with default config.
		OutlierDetection: &envoycluster.OutlierDetection{},
	}, nil
}

func makeThresholdsIfNeeded(limits UpstreamLimits) []*envoycluster.CircuitBreakers_Thresholds {
	var empty UpstreamLimits
	// Make sure to not create any thresholds when passed the zero-value in order
	// to rely on Envoy defaults
	if limits == empty {
		return nil
	}

	threshold := &envoycluster.CircuitBreakers_Thresholds{}
	// Likewise, make sure to not set any threshold values on the zero-value in
	// order to rely on Envoy defaults
	if limits.MaxConnections != nil {
		threshold.MaxConnections = makeUint32Value(*limits.MaxConnections)
	}
	if limits.MaxPendingRequests != nil {
		threshold.MaxPendingRequests = makeUint32Value(*limits.MaxPendingRequests)
	}
	if limits.MaxConcurrentRequests != nil {
		threshold.MaxRequests = makeUint32Value(*limits.MaxConcurrentRequests)
	}

	return []*envoycluster.CircuitBreakers_Thresholds{threshold}
}

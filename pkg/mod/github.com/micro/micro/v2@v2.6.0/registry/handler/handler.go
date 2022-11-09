package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/service"
	pb "github.com/micro/go-micro/v2/registry/service/proto"
	"github.com/micro/micro/v2/internal/namespace"
)

type Registry struct {
	// service id
	Id string
	// the publisher
	Publisher micro.Publisher
	// internal registry
	Registry registry.Registry
	// auth to verify clients
	Auth auth.Auth
}

func ActionToEventType(action string) registry.EventType {
	switch action {
	case "create":
		return registry.Create
	case "delete":
		return registry.Delete
	default:
		return registry.Update
	}
}

func (r *Registry) publishEvent(action string, service *pb.Service) error {
	// TODO: timestamp should be read from received event
	// Right now registry.Result does not contain timestamp
	event := &pb.Event{
		Id:        r.Id,
		Type:      pb.EventType(ActionToEventType(action)),
		Timestamp: time.Now().UnixNano(),
		Service:   service,
	}

	log.Debugf("publishing event %s for action %s", event.Id, action)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Publisher.Publish(ctx, event)
}

func (r *Registry) GetService(ctx context.Context, req *pb.GetRequest, rsp *pb.GetResponse) error {
	// verify the context has access to read the service
	if !r.canReadService(ctx, req.Service) {
		return errors.Forbidden("go.micro.registry", "Cannot read service")
	}

	services, err := r.Registry.GetService(req.Service)
	if err != nil {
		return errors.InternalServerError("go.micro.registry", err.Error())
	}
	for _, srv := range services {
		rsp.Services = append(rsp.Services, service.ToProto(srv))
	}
	return nil
}

func (r *Registry) Register(ctx context.Context, req *pb.Service, rsp *pb.EmptyResponse) error {
	// validate the name is valid
	if _, err := namespace.FromService(req.Name); err != nil {
		return err
	}

	// verify the context has access to register the service
	if !r.canWriteService(ctx, req.Name) {
		return errors.Forbidden("go.micro.registry", "Cannot register service")
	}

	var regOpts []registry.RegisterOption
	if req.Options != nil {
		ttl := time.Duration(req.Options.Ttl) * time.Second
		regOpts = append(regOpts, registry.RegisterTTL(ttl))
	}

	err := r.Registry.Register(service.ToService(req), regOpts...)
	if err != nil {
		return errors.InternalServerError("go.micro.registry", err.Error())
	}

	// publish the event
	go r.publishEvent("create", req)

	return nil
}

func (r *Registry) Deregister(ctx context.Context, req *pb.Service, rsp *pb.EmptyResponse) error {
	// verify the context has access to deregister the service
	if !r.canWriteService(ctx, req.Name) {
		return errors.Forbidden("go.micro.registry", "Cannot deregister service")
	}

	err := r.Registry.Deregister(service.ToService(req))
	if err != nil {
		return errors.InternalServerError("go.micro.registry", err.Error())
	}

	// publish the event
	go r.publishEvent("delete", req)

	return nil
}

func (r *Registry) ListServices(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	services, err := r.Registry.ListServices()
	if err != nil {
		return errors.InternalServerError("go.micro.registry", err.Error())
	}
	for _, srv := range services {
		if r.canReadService(ctx, srv.Name) {
			rsp.Services = append(rsp.Services, service.ToProto(srv))
		}
	}

	return nil
}

func (r *Registry) Watch(ctx context.Context, req *pb.WatchRequest, rsp pb.Registry_WatchStream) error {
	watcher, err := r.Registry.Watch(registry.WatchService(req.Service))
	if err != nil {
		return errors.InternalServerError("go.micro.registry", err.Error())
	}

	for {
		next, err := watcher.Next()
		if err != nil {
			return errors.InternalServerError("go.micro.registry", err.Error())
		}
		if !r.canReadService(ctx, next.Service.Name) {
			continue
		}
		err = rsp.Send(&pb.Result{
			Action:  next.Action,
			Service: service.ToProto(next.Service),
		})
		if err != nil {
			return errors.InternalServerError("go.micro.registry", err.Error())
		}
	}
}

// canReadService returns a boolean indicating is the context has
// permission to read the service
func (r *Registry) canReadService(ctx context.Context, name string) bool {
	// allow all services is no auth is enabled
	if r.Auth.String() == "noop" {
		return true
	}

	ns, err := namespace.FromService(name)
	if err != nil {
		// This should never happen as namespaces will be validated
		log.Warnf("Invalid service name: %v", name)
		return false
	}

	// always allow access to the default namespace (shared), as well
	// as the runtime namespace.
	if ns == namespace.DefaultNamespace || ns == namespace.RuntimeNamespace {
		return true
	}

	// get the namespace from the context and compare this to the services
	// namespace, if they match we allow access. We also allow the runtime
	// services access to read any services (needed for micro web etc).
	ctxNs := namespace.FromContext(ctx)
	return ctxNs == ns || ctxNs == namespace.RuntimeNamespace
}

// canReadService returns a boolean indicating is the context has
// permission to write (amend) a service
func (r *Registry) canWriteService(ctx context.Context, name string) bool {
	// allow all services is no auth is enabled
	if r.Auth.String() == "noop" {
		return true
	}

	ns, err := namespace.FromService(name)

	// the data in the registry is invalid, log an error and don't allow
	// access. This should never happen as namespaces will be validated
	// when services are registered.
	if err != nil {
		log.Warnf("Invalid service name: %v", name)
		return false
	}

	// always allow access to the default namespace (shared)
	if ns == namespace.DefaultNamespace {
		return true
	}

	// get the namespace from the context and compare this to the services
	// namespace, if they match we allow access
	return namespace.FromContext(ctx) == ns
}

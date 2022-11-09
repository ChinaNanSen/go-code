// Package web is a web dashboard
package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-acme/lego/v3/providers/dns/cloudflare"
	"github.com/gorilla/mux"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api/server"
	"github.com/micro/go-micro/api/server/acme"
	"github.com/micro/go-micro/api/server/acme/autocert"
	"github.com/micro/go-micro/api/server/acme/certmagic"
	httpapi "github.com/micro/go-micro/api/server/http"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/cache"
	cfstore "github.com/micro/go-micro/store/cloudflare"
	"github.com/micro/go-micro/sync/lock/memory"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/micro/internal/handler"
	"github.com/micro/micro/internal/helper"
	"github.com/micro/micro/internal/stats"
	"github.com/micro/micro/plugin"
	"github.com/serenize/snaker"
)

var (
	re = regexp.MustCompile("^[a-zA-Z0-9]+([a-zA-Z0-9-]*[a-zA-Z0-9]*)?$")
	// Default server name
	Name = "go.micro.web"
	// Default address to bind to
	Address = ":8082"
	// The namespace to serve
	// Example:
	// Namespace + /[Service]/foo/bar
	// Host: Namespace.Service Endpoint: /foo/bar
	Namespace = "go.micro.web"
	// Base path sent to web service.
	// This is stripped from the request path
	// Allows the web service to define absolute paths
	BasePathHeader        = "X-Micro-Web-Base-Path"
	statsURL              string
	ACMEProvider          = "autocert"
	ACMEChallengeProvider = "cloudflare"
	ACMECA                = acme.LetsEncryptProductionCA

	// A placeholder icon
	DefaultIcon = "https://micro.mu/circle.png"
)

type srv struct {
	*mux.Router
	// registry we use
	registry registry.Registry
}

type reg struct {
	registry.Registry

	sync.Mutex
	lastPull time.Time
	services []*registry.Service
}

func (r *reg) watch() {
Loop:
	for {
		// get a watcher
		w, err := r.Registry.Watch()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		// loop results
		for {
			_, err := w.Next()
			if err != nil {
				w.Stop()
				time.Sleep(time.Second)
				goto Loop
			}

			// next pull will be from the registry
			r.Lock()
			r.lastPull = time.Time{}
			r.Unlock()
		}
	}
}

func (r *reg) ListServices() ([]*registry.Service, error) {
	r.Lock()
	defer r.Unlock()

	if r.lastPull.IsZero() || time.Since(r.lastPull) > time.Minute {
		// pull the services
		s, err := r.Registry.ListServices()
		if err != nil {
			// return stale entries if they exist
			if len(r.services) > 0 {
				return r.services, nil
			}
			// otherwise return an error
			return nil, err
		}
		// collapse the list
		serviceMap := make(map[string]*registry.Service)
		for _, service := range s {
			serviceMap[service.Name] = service
		}
		var services []*registry.Service
		for _, service := range serviceMap {
			services = append(services, service)
		}
		// cache it
		r.services = services
		r.lastPull = time.Now()
		return s, nil
	}

	// return the cached list
	return r.services, nil
}

func (s *srv) proxy() http.Handler {
	sel := selector.NewSelector(
		selector.Registry(s.registry),
	)

	director := func(r *http.Request) {
		kill := func() {
			r.URL.Host = ""
			r.URL.Path = ""
			r.URL.Scheme = ""
			r.Host = ""
			r.RequestURI = ""
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 2 {
			kill()
			return
		}
		if !re.MatchString(parts[1]) {
			kill()
			return
		}
		next, err := sel.Select(Namespace + "." + parts[1])
		if err != nil {
			kill()
			return
		}

		s, err := next()
		if err != nil {
			kill()
			return
		}

		r.Header.Set(BasePathHeader, "/"+parts[1])
		r.URL.Host = s.Address
		r.URL.Path = "/" + strings.Join(parts[2:], "/")
		r.URL.Scheme = "http"
		r.Host = r.URL.Host
	}

	return &proxy{
		Default:  &httputil.ReverseProxy{Director: director},
		Director: director,
	}
}

func format(v *registry.Value) string {
	if v == nil || len(v.Values) == 0 {
		return "{}"
	}
	var f []string
	for _, k := range v.Values {
		f = append(f, formatEndpoint(k, 0))
	}
	return fmt.Sprintf("{\n%s}", strings.Join(f, ""))
}

func formatEndpoint(v *registry.Value, r int) string {
	// default format is tabbed plus the value plus new line
	fparts := []string{"", "%s %s", "\n"}
	for i := 0; i < r+1; i++ {
		fparts[0] += "\t"
	}
	// its just a primitive of sorts so return
	if len(v.Values) == 0 {
		return fmt.Sprintf(strings.Join(fparts, ""), snaker.CamelToSnake(v.Name), v.Type)
	}

	// this thing has more things, it's complex
	fparts[1] += " {"

	vals := []interface{}{snaker.CamelToSnake(v.Name), v.Type}

	for _, val := range v.Values {
		fparts = append(fparts, "%s")
		vals = append(vals, formatEndpoint(val, r+1))
	}

	// at the end
	l := len(fparts) - 1
	for i := 0; i < r+1; i++ {
		fparts[l] += "\t"
	}
	fparts = append(fparts, "}\n")

	return fmt.Sprintf(strings.Join(fparts, ""), vals...)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *srv) cliHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, cliTemplate, nil)
}

func (s *srv) indexHandler(w http.ResponseWriter, r *http.Request) {
	helper.ServeCORS(w, r)

	if r.Method == "OPTIONS" {
		return
	}

	services, err := s.registry.ListServices()
	if err != nil {
		log.Errorf("Error listing services: %v", err)
	}

	type webService struct {
		Name string
		Icon string
	}

	var webServices []webService
	for _, srv := range services {
		if len(srv.Nodes) == 0 {
			srvs, _ := s.registry.GetService(srv.Name)
			if len(srvs) == 0 {
				continue
			}
			srv = srvs[0]
		}

		if len(srv.Nodes) == 0 {
			continue
		}

		if strings.Index(srv.Name, Namespace) == 0 && len(strings.TrimPrefix(srv.Name, Namespace)) > 0 {
			icon := DefaultIcon

			if ico := srv.Nodes[0].Metadata["icon"]; len(ico) > 0 {
				icon = ico
			}

			webServices = append(webServices, webService{
				Name: strings.Replace(srv.Name, Namespace+".", "", 1),
				Icon: icon,
			})
		}
	}

	sort.Slice(webServices, func(i, j int) bool { return webServices[i].Name < webServices[j].Name })

	type templateData struct {
		HasWebServices bool
		WebServices    []webService
	}

	data := templateData{len(webServices) > 0, webServices}
	render(w, r, indexTemplate, data)
}

func (s *srv) registryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	svc := r.Form.Get("service")

	if len(svc) > 0 {
		s, err := s.registry.GetService(svc)
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}

		if len(s) == 0 {
			http.Error(w, "Not found", 404)
			return
		}

		if r.Header.Get("Content-Type") == "application/json" {
			b, err := json.Marshal(map[string]interface{}{
				"services": s,
			})
			if err != nil {
				http.Error(w, "Error occurred:"+err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		render(w, r, serviceTemplate, s)
		return
	}

	services, err := s.registry.ListServices()
	if err != nil {
		log.Errorf("Error listing services: %v", err)
	}

	sort.Sort(sortedServices{services})

	if r.Header.Get("Content-Type") == "application/json" {
		b, err := json.Marshal(map[string]interface{}{
			"services": services,
		})
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	render(w, r, registryTemplate, services)
}

func (s *srv) callHandler(w http.ResponseWriter, r *http.Request) {
	services, err := s.registry.ListServices()
	if err != nil {
		log.Errorf("Error listing services: %v", err)
	}

	sort.Sort(sortedServices{services})

	serviceMap := make(map[string][]*registry.Endpoint)
	for _, service := range services {
		if len(service.Endpoints) > 0 {
			serviceMap[service.Name] = service.Endpoints
			continue
		}
		// lookup the endpoints otherwise
		s, err := s.registry.GetService(service.Name)
		if err != nil {
			continue
		}
		if len(s) == 0 {
			continue
		}
		serviceMap[service.Name] = s[0].Endpoints
	}

	if r.Header.Get("Content-Type") == "application/json" {
		b, err := json.Marshal(map[string]interface{}{
			"services": services,
		})
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	render(w, r, callTemplate, serviceMap)
}

func render(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
	t, err := template.New("template").Funcs(template.FuncMap{
		"format": format,
	}).Parse(layoutTemplate)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}
	t, err = t.Parse(tmpl)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	if err := t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"StatsURL": statsURL,
		"Results":  data,
	}); err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
	}
}

func run(ctx *cli.Context, srvOpts ...micro.Option) {
	log.Name("web")

	if len(ctx.GlobalString("server_name")) > 0 {
		Name = ctx.GlobalString("server_name")
	}
	if len(ctx.String("address")) > 0 {
		Address = ctx.String("address")
	}
	if len(ctx.String("namespace")) > 0 {
		Namespace = ctx.String("namespace")
	}

	// Init plugins
	for _, p := range Plugins() {
		p.Init(ctx)
	}

	// use the caching registry
	cache := cache.New((*cmd.DefaultOptions().Registry))
	reg := &reg{Registry: cache}

	// start the watcher
	go reg.watch()

	var h http.Handler
	s := &srv{
		Router:   mux.NewRouter(),
		registry: reg,
	}
	h = s

	if ctx.GlobalBool("enable_stats") {
		statsURL = "/stats"
		st := stats.New()
		s.HandleFunc("/stats", st.StatsHandler)
		h = st.ServeHTTP(s)
		st.Start()
		defer st.Stop()
	}

	s.HandleFunc("/client", s.callHandler)
	s.HandleFunc("/registry", s.registryHandler)
	s.HandleFunc("/terminal", s.cliHandler)
	s.HandleFunc("/rpc", handler.RPC)
	s.HandleFunc("/favicon.ico", faviconHandler)
	s.PathPrefix("/{service:[a-zA-Z0-9]+}").Handler(s.proxy())
	s.HandleFunc("/", s.indexHandler)

	var opts []server.Option

	if len(ctx.GlobalString("acme_provider")) > 0 {
		ACMEProvider = ctx.GlobalString("acme_provider")
	}
	if ctx.GlobalBool("enable_acme") {
		hosts := helper.ACMEHosts(ctx)
		opts = append(opts, server.EnableACME(true))
		opts = append(opts, server.ACMEHosts(hosts...))
		switch ACMEProvider {
		case "autocert":
			opts = append(opts, server.ACMEProvider(autocert.New()))
		case "certmagic":
			if ACMEChallengeProvider != "cloudflare" {
				log.Fatal("The only implemented DNS challenge provider is cloudflare")
			}
			apiToken, accountID := os.Getenv("CF_API_TOKEN"), os.Getenv("CF_ACCOUNT_ID")
			kvID := os.Getenv("KV_NAMESPACE_ID")
			if len(apiToken) == 0 || len(accountID) == 0 {
				log.Fatal("env variables CF_API_TOKEN and CF_ACCOUNT_ID must be set")
			}
			if len(kvID) == 0 {
				log.Fatal("env var KV_NAMESPACE_ID must be set to your cloudflare workers KV namespace ID")
			}

			cloudflareStore := cfstore.NewStore(
				cfstore.Token(apiToken),
				cfstore.Account(accountID),
				cfstore.Namespace(kvID),
			)
			storage := certmagic.NewStorage(
				memory.NewLock(),
				cloudflareStore,
			)
			config := cloudflare.NewDefaultConfig()
			config.AuthToken = apiToken
			config.ZoneToken = apiToken
			challengeProvider, err := cloudflare.NewDNSProviderConfig(config)
			if err != nil {
				log.Fatal(err.Error())
			}

			opts = append(opts,
				server.ACMEProvider(
					certmagic.New(
						acme.AcceptToS(true),
						acme.CA(ACMECA),
						acme.Cache(storage),
						acme.ChallengeProvider(challengeProvider),
						acme.OnDemand(false),
					),
				),
			)
		default:
			log.Fatalf("%s is not a valid ACME provider\n", ACMEProvider)
		}
	} else if ctx.GlobalBool("enable_tls") {
		config, err := helper.TLSConfig(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		opts = append(opts, server.EnableTLS(true))
		opts = append(opts, server.TLSConfig(config))
	}

	// reverse wrap handler
	plugins := append(Plugins(), plugin.Plugins()...)
	for i := len(plugins); i > 0; i-- {
		h = plugins[i-1].Handler()(h)
	}

	srv := httpapi.NewServer(Address)
	srv.Init(opts...)
	srv.Handle("/", h)

	// service opts
	srvOpts = append(srvOpts, micro.Name(Name))
	if i := time.Duration(ctx.GlobalInt("register_ttl")); i > 0 {
		srvOpts = append(srvOpts, micro.RegisterTTL(i*time.Second))
	}
	if i := time.Duration(ctx.GlobalInt("register_interval")); i > 0 {
		srvOpts = append(srvOpts, micro.RegisterInterval(i*time.Second))
	}

	// Initialise Server
	service := micro.NewService(srvOpts...)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	if err := srv.Stop(); err != nil {
		log.Fatal(err)
	}
}

func Commands(options ...micro.Option) []cli.Command {
	command := cli.Command{
		Name:  "web",
		Usage: "Run the web dashboard",
		Action: func(c *cli.Context) {
			run(c, options...)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "address",
				Usage:  "Set the web UI address e.g 0.0.0.0:8082",
				EnvVar: "MICRO_WEB_ADDRESS",
			},
			cli.StringFlag{
				Name:   "namespace",
				Usage:  "Set the namespace used by the Web proxy e.g. com.example.web",
				EnvVar: "MICRO_WEB_NAMESPACE",
			},
		},
	}

	for _, p := range Plugins() {
		if cmds := p.Commands(); len(cmds) > 0 {
			command.Subcommands = append(command.Subcommands, cmds...)
		}

		if flags := p.Flags(); len(flags) > 0 {
			command.Flags = append(command.Flags, flags...)
		}
	}

	return []cli.Command{command}
}

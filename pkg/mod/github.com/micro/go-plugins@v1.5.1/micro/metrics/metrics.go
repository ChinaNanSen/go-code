// Package metrics provides metrics for micro
package metrics

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/micro/plugin"

	// prometheus metrics
	"github.com/micro/go-plugins/micro/metrics/prometheus"
)

type Metrics struct {
	Path string

	Provider Provider
}

// Provider is a metrics provider
type Provider interface {
	Handler(h http.Handler) http.Handler
}

func (m *Metrics) Handler(h http.Handler) http.Handler {
	return m.Provider.Handler(h)
}

// NewPlugin returns a new metrics plugin
func NewPlugin() plugin.Plugin {
	metrics := new(Metrics)

	return plugin.NewPlugin(
		plugin.WithName("metrics"),
		plugin.WithFlag(
			cli.StringFlag{
				Name:  "metrics",
				Usage: "Specify the type of metrics provider e.g prometheus",
			},
		),
		plugin.WithHandler(metrics.Handler),
		plugin.WithInit(func(ctx *cli.Context) error {
			provider := ctx.String("metrics")

			switch provider {
			case "prometheus":
				metrics.Provider = prometheus.New()
				log.Log("Loaded prometheus metrics at /metrics")
			}

			return nil
		}),
	)
}

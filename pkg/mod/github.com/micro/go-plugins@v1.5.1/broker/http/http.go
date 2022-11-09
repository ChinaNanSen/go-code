// Package http provides a http based message broker
package http

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["http"] = NewBroker
}

// NewBroker returns a new http broker
func NewBroker(opts ...broker.Option) broker.Broker {
	return broker.NewBroker(opts...)
}

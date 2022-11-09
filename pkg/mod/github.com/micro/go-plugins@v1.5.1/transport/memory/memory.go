// Package memory is an in-memory transport
package memory

import (
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/transport/memory"
)

func init() {
	cmd.DefaultTransports["memory"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return memory.NewTransport(opts...)
}

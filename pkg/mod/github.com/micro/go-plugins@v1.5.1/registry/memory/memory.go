// Package memory provides an in-memory registry
package memory

import (
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/memory"
)

func init() {
	cmd.DefaultRegistries["memory"] = NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return memory.NewRegistry(opts...)
}

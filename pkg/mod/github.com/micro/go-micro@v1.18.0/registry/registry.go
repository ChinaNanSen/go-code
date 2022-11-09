// Package registry is an interface for service discovery
package registry

import (
	"errors"
	"github.com/micro/go-micro/v2/registry"
)

var (
	DefaultRegistry = NewRegistry()

	// Not found error when GetService is called
	ErrNotFound = errors.New("service not found")
	// Watcher stopped error when watcher is stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

func (r Registry) Init(...registry.Option) error {
	panic("implement me")
}

func (r Registry) Options() registry.Options {
	panic("implement me")
}

func (r Registry) Register(*registry.Service, ...registry.RegisterOption) error {
	panic("implement me")
}

func (r Registry) Deregister(*registry.Service, ...registry.DeregisterOption) error {
	panic("implement me")
}

func (r Registry) GetService(string, ...registry.GetOption) ([]*registry.Service, error) {
	panic("implement me")
}

func (r Registry) ListServices(...registry.ListOption) ([]*registry.Service, error) {
	panic("implement me")
}

func (r Registry) Watch(...registry.WatchOption) (registry.Watcher, error) {
	panic("implement me")
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)

// Register a service node. Additionally supply options such as TTL.
func Register(s *Service, opts ...RegisterOption) error {
	return DefaultRegistry.Register(s, opts...)
}

// Deregister a service node
func Deregister(s *Service) error {
	return DefaultRegistry.Deregister(s)
}

// Retrieve a service. A slice is returned since we separate Name/Version.
func GetService(name string) ([]*Service, error) {
	return DefaultRegistry.GetService(name)
}

// List the services. Only returns service names
func ListServices() ([]*Service, error) {
	return DefaultRegistry.ListServices()
}

// Watch returns a watcher which allows you to track updates to the registry.
func Watch(opts ...WatchOption) (Watcher, error) {
	return DefaultRegistry.Watch(opts...)
}

func String() string {
	return DefaultRegistry.String()
}

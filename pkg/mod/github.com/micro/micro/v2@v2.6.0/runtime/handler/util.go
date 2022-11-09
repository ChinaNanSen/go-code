package handler

import (
	"context"

	"github.com/micro/go-micro/v2/runtime"
	pb "github.com/micro/go-micro/v2/runtime/service/proto"
	"github.com/micro/micro/v2/internal/namespace"
)

func toProto(s *runtime.Service) *pb.Service {
	return &pb.Service{
		Name:     s.Name,
		Version:  s.Version,
		Source:   s.Source,
		Metadata: s.Metadata,
	}
}

func toService(s *pb.Service) *runtime.Service {
	return &runtime.Service{
		Name:     s.Name,
		Version:  s.Version,
		Source:   s.Source,
		Metadata: s.Metadata,
	}
}

// getNamespace replaces the default auth namespace until we move
// we wil replace go.micro with micro and move our default things there
func getNamespace(ctx context.Context) string {
	ns := namespace.FromContext(ctx)
	if len(ns) == 0 || ns == "go.micro" {
		return "default"
	}
	return ns
}

func toCreateOptions(ctx context.Context, opts *pb.CreateOptions) []runtime.CreateOption {
	options := []runtime.CreateOption{
		runtime.CreateNamespace(getNamespace(ctx)),
	}

	// stop if no options were passed
	if opts == nil {
		return options
	}

	// command options
	if len(opts.Command) > 0 {
		options = append(options, runtime.WithCommand(opts.Command...))
	}

	// args for command
	if len(opts.Args) > 0 {
		options = append(options, runtime.WithArgs(opts.Args...))
	}

	// env options
	if len(opts.Env) > 0 {
		options = append(options, runtime.WithEnv(opts.Env))
	}

	// create specific type of service
	if len(opts.Type) > 0 {
		options = append(options, runtime.CreateType(opts.Type))
	}

	// use specific image
	if len(opts.Image) > 0 {
		options = append(options, runtime.CreateImage(opts.Image))
	}

	// TODO: output options

	return options
}

func toReadOptions(ctx context.Context, opts *pb.ReadOptions) []runtime.ReadOption {
	options := []runtime.ReadOption{
		runtime.ReadNamespace(getNamespace(ctx)),
	}

	// stop if no options were passed
	if opts == nil {
		return options
	}

	if len(opts.Service) > 0 {
		options = append(options, runtime.ReadService(opts.Service))
	}
	if len(opts.Version) > 0 {
		options = append(options, runtime.ReadVersion(opts.Version))
	}
	if len(opts.Type) > 0 {
		options = append(options, runtime.ReadType(opts.Type))
	}

	return options
}

func toUpdateOptions(ctx context.Context) []runtime.UpdateOption {
	return []runtime.UpdateOption{
		runtime.UpdateNamespace(getNamespace(ctx)),
	}
}

func toDeleteOptions(ctx context.Context) []runtime.DeleteOption {
	return []runtime.DeleteOption{
		runtime.DeleteNamespace(getNamespace(ctx)),
	}
}

func toLogsOptions(ctx context.Context) []runtime.LogsOption {
	return []runtime.LogsOption{
		runtime.LogsNamespace(getNamespace(ctx)),
	}
}

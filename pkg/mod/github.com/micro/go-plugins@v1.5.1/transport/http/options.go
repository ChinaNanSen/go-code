package http

import (
	"net/http"

	"github.com/micro/go-micro/transport"
	thttp "github.com/micro/go-micro/transport/http"
)

// Handle registers the handler for the given pattern.
func Handle(pattern string, handler http.Handler) transport.Option {
	return thttp.Handle(pattern, handler)
}

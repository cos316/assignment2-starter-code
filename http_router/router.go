package http_router

import (
	"net/http"
)

// HTTPRouter is an HTTP Router
type HTTPRouter struct {
}

// NewRouter creates a new HTTP Router
func NewRouter() *HTTPRouter {
	return &HTTPRouter {}
}

// AddRoute adds a new route to the router
func (router *HTTPRouter) AddRoute(method string, path string, handler http.HandlerFunc) {
}

func (router *HTTPRouter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
}

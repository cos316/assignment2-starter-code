/*****************************************************************************
 * router.go
 * Name:
 * NetId:
 *****************************************************************************/

package http_router

import (
	"net/http"
)

// Student defined types or constants go here

// HTTPRouter stores the information necessary to route HTTP requests
type HTTPRouter struct {
	// Place anything you'd like here
}

// NewRouter creates a new HTTP Router, with no initial routes
func NewRouter() *HTTPRouter {
	return new(HTTPRouter)
}

// AddRoute adds a new route to the router, associating a given method and path
// pattern with the designated http handler.
func (router *HTTPRouter) AddRoute(method string, pattern string, handler http.HandlerFunc) {
	return
}

// ServeHTTP writes an HTTP response to the provided response writer
// by invoking the handler associated with the route that is appropriate
// for the provided request.
func (router *HTTPRouter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	http.NotFound(response, request)
}

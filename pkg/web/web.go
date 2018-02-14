package web

import "github.com/gorilla/mux"

// NewRouter returns the router created by the internal newRouter function
// Don't expose it directly in case we need to pass extra config that doesn't
// apply to the router.
func NewRouter(config Config) *mux.Router {
	return newRouter(config.Routes)
}

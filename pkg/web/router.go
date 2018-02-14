package web

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
)

func newRouter(routes RouteList) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	router.Use(loggerMiddleWare)
	router.Use(jsonContentTypeMiddleWare)
	router.Use(boom.RecoverHandler)

	router.NotFoundHandler = http.HandlerFunc(boom.NotFoundHandler)

	return router
}

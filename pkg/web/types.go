package web

import (
	"net/http"
)

type Config struct {
	Routes RouteList
}

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type RouteList []Route

package handlers

import (
	"github.com/aerogear/aerogear-metrics-api/pkg/models"
	"github.com/aerogear/aerogear-metrics-api/pkg/web"
)

var app models.App

type handlers struct {
	Routes web.RouteList
}

func BuildHandlers(appInstance models.App) handlers {

	app = appInstance

	routes := web.RouteList{
		web.Route{
			Name:    "Health Check",
			Method:  "GET",
			Path:    "/healthz",
			Handler: healthz,
		},
		web.Route{
			Name:    "Create Metrics",
			Method:  "POST",
			Path:    "/metrics",
			Handler: createMetric,
		},
	}

	return handlers{
		Routes: routes,
	}
}

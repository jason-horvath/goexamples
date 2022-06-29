package router

import "github.com/jason-horvath/goexamples/handlers"

type Route struct {
	Uri     string
	Handler handlers.Handler
}

type RouteCollection struct {
	GET    map[string]Route
	POST   map[string]Route
	PUT    map[string]Route
	PATCH  map[string]Route
	DELETE map[string]Route
}

func Routes() {

}

func (rc *RouteCollection) SetRoute(httpType string, route Route) {

}

// Still hashing out the ideas on this. Just saving here.

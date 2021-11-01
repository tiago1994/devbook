package routes

import (
	"devbook/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI       string
	Method    string
	Function  func(http.ResponseWriter, *http.Request)
	IsPrivate bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes)
	routes = append(routes, publicationRoutes...)

	for _, route := range routes {
		if route.IsPrivate {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}

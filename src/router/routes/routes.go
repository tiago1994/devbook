package routes

import (
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
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}

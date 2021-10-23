package router

import (
	"devbook/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}

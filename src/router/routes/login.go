package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var loginRoutes = Route{
	URI:       "/login",
	Method:    http.MethodPost,
	IsPrivate: false,
	Function:  controllers.Login,
}

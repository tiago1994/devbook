package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:       "/users",
		Method:    http.MethodPost,
		IsPrivate: false,
		Function:  controllers.CreateUser,
	},
	{
		URI:       "/users",
		Method:    http.MethodGet,
		IsPrivate: false,
		Function:  controllers.GetAllUser,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		IsPrivate: false,
		Function:  controllers.GetUser,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodPut,
		IsPrivate: false,
		Function:  controllers.UpdateUser,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodDelete,
		IsPrivate: false,
		Function:  controllers.DeleteUser,
	},
}

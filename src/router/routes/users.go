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
		IsPrivate: true,
		Function:  controllers.GetAllUser,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		IsPrivate: true,
		Function:  controllers.GetUser,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodPut,
		IsPrivate: true,
		Function:  controllers.UpdateUser,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodDelete,
		IsPrivate: true,
		Function:  controllers.DeleteUser,
	},
}

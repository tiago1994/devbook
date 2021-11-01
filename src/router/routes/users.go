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
	{
		URI:       "/users/{userId}/follow",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.FollowUser,
	},
	{
		URI:       "/users/{userId}/unfollow",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.UnFollowUser,
	},
	{
		URI:       "/users/{userId}/followers",
		Method:    http.MethodGet,
		IsPrivate: true,
		Function:  controllers.GetFollowers,
	},
	{
		URI:       "/users/{userId}/following",
		Method:    http.MethodGet,
		IsPrivate: true,
		Function:  controllers.GetFollowing,
	},
	{
		URI:       "/users/{userId}/updatePassword",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.UpdatePassword,
	},
}

package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:       "/publications",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.CreatePublication,
	},
	{
		URI:       "/publications",
		Method:    http.MethodGet,
		IsPrivate: true,
		Function:  controllers.GetAllPublication,
	},
	{
		URI:       "/publications/{publicationId}",
		Method:    http.MethodGet,
		IsPrivate: true,
		Function:  controllers.GetPublication,
	},
	{
		URI:       "/publications/{publicationId}",
		Method:    http.MethodPut,
		IsPrivate: true,
		Function:  controllers.UpdatePublication,
	},
	{
		URI:       "/publications/{publicationId}",
		Method:    http.MethodDelete,
		IsPrivate: true,
		Function:  controllers.DeletePublication,
	},
	{
		URI:       "/publications/{publicationId}/like",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.Like,
	},
	{
		URI:       "/publications/{publicationId}/dislike",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.Like,
	},
	{
		URI:       "/users/{userId}/publications",
		Method:    http.MethodPost,
		IsPrivate: true,
		Function:  controllers.GetPublicationsByUser,
	},
}

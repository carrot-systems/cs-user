package rest

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.POST("/users", routesHandler.CreateUserHandler)
	authenticatedRoutes := r.Use(routesHandler.GetUserMiddleware)
	authenticatedRoutes.GET("/users/:handle", routesHandler.GetProfileHandler)
	authenticatedRoutes.PUT("/users/:handle", routesHandler.EditProfileHandler)
	authenticatedRoutes.DELETE("/users/:handle", routesHandler.RemoveUserHandler)
}

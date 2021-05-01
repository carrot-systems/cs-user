package rest

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.POST("/users", routesHandler.CreateUserHandler)
	authenticatedRoutes := r.Use(routesHandler.GetUserMiddleware)
	authenticatedRoutes.GET("/me", routesHandler.GetProfileHandler)
	authenticatedRoutes.PUT("/me", routesHandler.EditProfileHandler)
	authenticatedRoutes.DELETE("/me", routesHandler.RemoveUserHandler)
}

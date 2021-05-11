package rest

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	group := r.Group("/users")
	group.POST("/", routesHandler.CreateUserHandler)
	authenticatedRoutes := group.Use(routesHandler.GetUserMiddleware)
	authenticatedRoutes.GET("/:handle", routesHandler.GetProfileHandler)
	authenticatedRoutes.PUT("/:handle", routesHandler.EditProfileHandler)
	authenticatedRoutes.DELETE("/:handle", routesHandler.RemoveUserHandler)
}

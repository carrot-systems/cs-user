package rest

import (
	"github.com/gin-gonic/gin"
)

func (rH RoutesHandler) GetUserMiddleware(c *gin.Context) {
	c.Next()
}

func (rH RoutesHandler) CreateUserHandler(c *gin.Context) {

}

func (rH RoutesHandler) RemoveUserHandler(c *gin.Context) {

}

func (rH RoutesHandler) GetProfileHandler(c *gin.Context) {

}

func (rH RoutesHandler) EditProfileHandler(c *gin.Context) {

}

package rest

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) GetUserMiddleware(c *gin.Context) {
	c.Next()
}

func (rH RoutesHandler) CreateUserHandler(c *gin.Context) {
	var creationRequest domain.UserCreationRequest

	err := c.ShouldBindJSON(&creationRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = rH.Usecases.CreateUser(creationRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.ResponseUserCreated,
	})
}

func (rH RoutesHandler) RemoveUserHandler(c *gin.Context) {
	userHandler := "" //TODO: get handler
	userToRemove := c.Param("handler")

	err := rH.Usecases.RemoveUser(userHandler, userToRemove)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.ResponseUserDeleted,
	})
}

func (rH RoutesHandler) GetProfileHandler(c *gin.Context) {
	userHandler := "" //TODO: get handler
	userToGet := c.Param("handler")

	user, err := rH.Usecases.GetProfile(userHandler, userToGet)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.UserResponse{
		Status: domain.Status{
			Success: true,
		},
		User: user,
	})
}

func (rH RoutesHandler) EditProfileHandler(c *gin.Context) {
	auth, exists := c.Get("authenticatedUser")

	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: domain.ErrFailedToGetUser.Error(),
		})
		return
	}

	authenticatedUser := auth.(domain.User)

	userToEdit := c.Param("handler")
	var editRequest domain.User

	err := c.ShouldBind(&editRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = rH.Usecases.EditProfile(userHandler, userToEdit, &editRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Status{ //TODO: Translate error to correct code
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.ResponseUserCreated,
	})
}

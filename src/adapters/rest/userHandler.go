package rest

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) handleError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(codeForError(err), domain.Status{
		Success: false,
		Message: err.Error(),
	})
}

func (rH RoutesHandler) GetUserMiddleware(c *gin.Context) {
	var user domain.AuthenticatedUserResponse

	err := c.ShouldBindJSON(&user)
	if err != nil {
		rH.handleError(c, domain.ErrNotAuthenticated)
		return
	}

	authenticatedUser, err := rH.UserRepo.FindId(user.ID.String())

	if authenticatedUser == nil || err != nil {
		rH.handleError(c, domain.ErrNotAuthenticated)
		return
	}

	c.Set("authenticatedUser", authenticatedUser)

	c.Next()
}

func (rH RoutesHandler) getAuthenticatedUser(c *gin.Context) *domain.User {
	auth, exists := c.Get("authenticatedUser")

	if !exists {
		rH.handleError(c, domain.ErrFailedToGetUser)
		return nil
	}

	authenticatedUser := auth.(*domain.User)

	return authenticatedUser
}

func (rH RoutesHandler) CreateUserHandler(c *gin.Context) {
	var creationRequest domain.UserCreationRequest

	err := c.ShouldBindJSON(&creationRequest)

	if err != nil {
		rH.handleError(c, ErrFormValidation)
		return
	}

	err = rH.Usecases.CreateUser(creationRequest)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.ResponseUserCreated,
	})
}

func (rH RoutesHandler) RemoveUserHandler(c *gin.Context) {
	authenticatedUser := rH.getAuthenticatedUser(c)

	if authenticatedUser == nil {
		return
	}

	userToRemove := c.Param("handle")

	err := rH.Usecases.RemoveUser(authenticatedUser, userToRemove)

	if err != nil {
		c.AbortWithStatusJSON(codeForError(err), domain.Status{
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
	authenticatedUser := rH.getAuthenticatedUser(c)

	if authenticatedUser == nil {
		return
	}

	userToGet := c.Param("handle")

	user, err := rH.Usecases.GetProfile(authenticatedUser, userToGet)

	if err != nil {
		rH.handleError(c, err)
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
	authenticatedUser := rH.getAuthenticatedUser(c)

	if authenticatedUser == nil {
		return
	}

	userToEdit := c.Param("handle")
	var editRequest domain.User

	err := c.ShouldBind(&editRequest)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	err = rH.Usecases.EditProfile(authenticatedUser, userToEdit, &editRequest)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.ResponseUserCreated,
	})
}

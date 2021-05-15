package rest

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (rH RoutesHandler) handleError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(codeForError(err), domain.Status{
		Success: false,
		Message: err.Error(),
	})
}

func (rH RoutesHandler) handleOk(c *gin.Context, message string, data interface{}) {
	c.JSON(codeForOkStatus(message), domain.Status{
		Success: true,
		Message: message,
		Data:    data,
	})
}

type Claims struct {
	jwt.StandardClaims
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
}

var jwtKey = []byte("my_secret_key")

func (rH RoutesHandler) GetUserMiddleware(c *gin.Context) {
	va := c.GetHeader("Authorization")

	spew.Dump(va)

	//////////////

	tknStr := va

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	spew.Dump(tkn)
	spew.Dump(claims)

	////////////

	authenticatedUser, err := rH.UserRepo.FindIdWithoutCredentials(claims.UserId)

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

	rH.handleOk(c, domain.StatusUserCreated, nil)
}

func (rH RoutesHandler) RemoveUserHandler(c *gin.Context) {
	authenticatedUser := rH.getAuthenticatedUser(c)

	if authenticatedUser == nil {
		return
	}

	userToRemove := c.Param("handle")

	err := rH.Usecases.RemoveUser(authenticatedUser, userToRemove)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	rH.handleOk(c, domain.StatusUserDeleted, nil)
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

	rH.handleOk(c, domain.StatusUserCreated, user)
}

func (rH RoutesHandler) GetMyProfile(c *gin.Context) {
	authenticatedUser := rH.getAuthenticatedUser(c)

	if authenticatedUser == nil {
		return
	}

	userToGet := authenticatedUser.Handle

	user, err := rH.Usecases.GetProfile(authenticatedUser, userToGet)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	rH.handleOk(c, domain.StatusUserCreated, user)
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

	rH.handleOk(c, domain.StatusUserCreated, nil) //TODO: better update code
}

func (rH RoutesHandler) GetProfileIdHandler(c *gin.Context) {
	profileToGet := c.Param("handle")
	var cred domain.Credentials

	err := c.BindQuery(&cred)

	if err != nil {
		rH.handleError(c, ErrFormValidation)
		return
	}

	id, err := rH.Usecases.GetProfileId(profileToGet, cred)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	rH.handleOk(c, domain.StatusCredentialsOk, id)
}

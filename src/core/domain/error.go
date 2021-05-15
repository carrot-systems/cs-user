package domain

import "errors"

var (
	ErrHandleAlreadyUsed     = errors.New("this handle already exists")
	ErrNotAuthenticated      = errors.New("you need to be authenticated to perform this action")
	ErrForbidden             = errors.New("you don't have the necessary permissions to perform this action")
	ErrConnectedUserNotFound = errors.New("connected user not found")
	ErrUserCreation          = errors.New("cannot create user")

	//TODO: merge those two ?
	ErrFailedToGetUser = errors.New("failed to get user")
	ErrUserNotFound    = errors.New("user not found") //Keeping this one
)

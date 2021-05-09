package domain

import "errors"

var (
	ErrHandleAlreadyUsed     = errors.New("this handle already exists.")
	ErrConnectedUserNotFound = errors.New("connected user not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserCreation          = errors.New("cannot create user")
)

package rest

import (
	"errors"
	"github.com/carrot-systems/cs-user/src/core/domain"
	"net/http"
)

var (
	ErrFormValidation = errors.New("failed to validate form")
)

func codeForError(err error) int {
	switch err {
	case domain.ErrHandleAlreadyUsed:
		return http.StatusConflict
	case domain.ErrNotAuthenticated:
		return http.StatusUnauthorized
	case domain.ErrForbidden:
		return http.StatusForbidden
	case domain.ErrConnectedUserNotFound, domain.ErrFailedToGetUser, domain.ErrUserNotFound:
		return http.StatusNotFound
	case domain.ErrUserCreation:
		return http.StatusInternalServerError
	case ErrFormValidation:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

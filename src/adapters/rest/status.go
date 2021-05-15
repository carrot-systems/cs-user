package rest

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"net/http"
)

func codeForOkStatus(status string) int {
	switch status {
	case domain.StatusUserCreated:
		return http.StatusCreated
	case domain.StatusUserDeleted:
		return http.StatusAccepted
	}
	return http.StatusOK
}

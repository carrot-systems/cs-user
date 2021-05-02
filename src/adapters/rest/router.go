package rest

import (
	"github.com/carrot-systems/cs-user/src/core/usecases"
)

type RoutesHandler struct {
	Usecases usecases.Usecases
	UserRepo usecases.UserRepo
}

func NewRouter(ucHandler usecases.Usecases, u usecases.UserRepo) RoutesHandler {
	return RoutesHandler{
		Usecases: ucHandler,
		UserRepo: u,
	}
}

package usecases

import "github.com/carrot-systems/cs-user/src/core/domain"

type UserRepo interface {
	FindUser(handler string) (*domain.User, error)
	CreateUser(user domain.UserCreationRequest) error
	DeleteUser(handler string) error
	UpdateUser(handler string, user *domain.User) error
}

type interactor struct {
	userRepo UserRepo
}

func NewInteractor() interactor {
	return interactor{}
}

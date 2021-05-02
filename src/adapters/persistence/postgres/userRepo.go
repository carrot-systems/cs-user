package postgres

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/carrot-systems/cs-user/src/core/usecases"
)

type userRepo struct {
}

func (u userRepo) FindUser(handler string) (*domain.User, error) {
	//TODO: implement
	panic("implement me")
}

func (u userRepo) CreateUser(user domain.UserCreationRequest) error {
	//TODO: implement
	panic("implement me")
}

func (u userRepo) DeleteUser(handler string) error {
	//TODO: implement
	panic("implement me")
}

func (u userRepo) UpdateUser(handler string, user *domain.User) error {
	//TODO: implement
	panic("implement me")
}

func NewUserRepo() usecases.UserRepo {
	return &userRepo{}
}

package usecases

import "github.com/carrot-systems/cs-user/src/core/domain"

type UserRepo interface {
	FindHandle(handle string) (*domain.User, error)
	FindId(id string) (*domain.User, error)
	CreateUser(user domain.UserCreationRequest) error
	DeleteUser(handle string) error
	UpdateUser(handle string, user *domain.User) error
}

type PermissionRepo interface {
	FindPermissions(id string, permission string) int
	SetPermissions(id string, permission string, flag int) error
}

type interactor struct {
	userRepo       UserRepo
	permissionRepo PermissionRepo
}

func NewInteractor(uR UserRepo, pR PermissionRepo) interactor {
	return interactor{
		userRepo:       uR,
		permissionRepo: pR,
	}
}

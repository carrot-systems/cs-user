package usecases

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/google/uuid"
)

type UserRepo interface {
	FindHandle(handle string) (*domain.User, error)
	FindIdWithoutCredentials(id string) (*domain.User, error)
	FindId(handle string, credentials domain.Credentials) (*domain.User, error)
	CreateUser(user domain.UserCreationRequest) error
	DeleteUser(handle string) error
	UpdateUser(handle string, user *domain.User) error
}

type PermissionRepo interface {
	FindPermissions(id uuid.UUID, permission string) int
	SetPermissions(id uuid.UUID, permission string, flag int) error
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

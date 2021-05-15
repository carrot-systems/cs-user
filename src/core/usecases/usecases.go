package usecases

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/google/uuid"
)

type Usecases interface {
	CreateUser(user domain.UserCreationRequest) error
	RemoveUser(connectedUser *domain.User, handle string) error
	GetProfile(connectedUser *domain.User, handle string) (*domain.User, error)
	EditProfile(connectedUser *domain.User, handle string, user *domain.User) error

	GetProfileId(handle string, credentials domain.Credentials) (uuid.UUID, error)
}

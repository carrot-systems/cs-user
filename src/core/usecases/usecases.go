package usecases

import "github.com/carrot-systems/cs-user/src/core/domain"

type Usecases interface {
	CreateUser(user domain.UserCreationRequest) error
	RemoveUser(connectedUser *domain.User, handle string) error
	GetProfile(connectedUser *domain.User, handle string) (*domain.User, error)
	EditProfile(connectedUser *domain.User, handle string, user *domain.User) error
}

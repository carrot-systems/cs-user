package usecases

import "github.com/carrot-systems/cs-user/src/core/domain"

type Usecases interface {
	CreateUser(user domain.UserCreationRequest) error
	RemoveUser(connectedUserHandle string, handle string) error
	GetProfile(connectedUserHandle string, handle string) (*domain.User, error)
	EditProfile(connectedUserHandle string, handle string, user *domain.User) error
}

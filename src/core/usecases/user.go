package usecases

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/google/uuid"
)

func (i interactor) hasPermissions(user *domain.User, permission int) bool {
	permissions := i.permissionRepo.FindPermissions(user.ID, domain.UserPermissionIdentifier)

	return permissions&permission != 0
}

func (i interactor) CreateUser(user domain.UserCreationRequest) error {
	userWithThisHandler, err := i.userRepo.FindHandle(user.User.Handle)

	if err == nil && userWithThisHandler != nil {
		return domain.ErrHandleAlreadyUsed
	}

	err = i.userRepo.CreateUser(user)

	if err != nil {
		return domain.ErrUserCreation
	}

	return nil
}

func (i interactor) RemoveUser(connectedUser *domain.User, handle string) error {
	var userToRemove *domain.User

	if handle == connectedUser.Handle {
		userToRemove = connectedUser
	} else {
		if !i.hasPermissions(connectedUser, domain.PermDeleteUser) {
			return domain.ErrForbidden
		}

		user, err := i.userRepo.FindHandle(handle)
		if err != nil {
			return domain.ErrFailedToGetUser
		}

		userToRemove = user
	}

	return i.userRepo.DeleteUser(userToRemove.Handle)
}

func (i interactor) GetProfile(connectedUser *domain.User, handle string) (*domain.User, error) {
	var userToFetch *domain.User

	if handle == connectedUser.Handle {
		return connectedUser, nil
	} else {
		if !i.hasPermissions(connectedUser, domain.PermReadAllUser) {
			return nil, domain.ErrForbidden
		}

		user, err := i.userRepo.FindHandle(handle)
		if err != nil {
			return nil, domain.ErrFailedToGetUser
		}

		userToFetch = user
	}

	return i.userRepo.FindHandle(userToFetch.Handle)
}

//TODO: implement permission management
func (i interactor) EditProfile(connectedUser *domain.User, handle string, user *domain.User) error {
	/*connectedUser, err := i.userRepo.FindHandle(connectedUserHandle)
	var userToEdit *domain.User

	if err != nil {
		return domain.ErrConnectedUserNotFound
	}

	if handle == connectedUserHandle {
		userToEdit = connectedUser
	}

	return i.userRepo.UpdateUser(userToEdit.Handle, user)*/
	return nil

}

func (i interactor) GetProfileId(handle string, credentials domain.Credentials) (uuid.UUID, error) {
	user, err := i.userRepo.FindId(handle, credentials)

	if err != nil || user == nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}

package usecases

import "github.com/carrot-systems/cs-user/src/core/domain"

func (i interactor) CreateUser(user domain.UserCreationRequest) error {
	userWithThisHandler, err := i.userRepo.FindUser(user.User.Handle)

	if err == nil && userWithThisHandler != nil {
		return domain.ErrHandleAlreadyUsed
	}

	err = i.userRepo.CreateUser(user)

	if err != nil {
		return domain.ErrUserCreation
	}

	return nil
}

//TODO: implement permission management
func (i interactor) RemoveUser(connectedUserHandle string, handle string) error {
	connectedUser, err := i.userRepo.FindUser(connectedUserHandle)
	var userToRemove *domain.User

	if err != nil {
		return domain.ErrConnectedUserNotFound
	}

	if handle == connectedUserHandle {
		userToRemove = connectedUser
	}

	return i.userRepo.DeleteUser(userToRemove.Handle)
}

//TODO: implement permission management
func (i interactor) GetProfile(connectedUserHandle string, handle string) (*domain.User, error) {
	connectedUser, err := i.userRepo.FindUser(connectedUserHandle)
	var userToFetch *domain.User

	if err != nil {
		return nil, domain.ErrConnectedUserNotFound
	}

	if handle == connectedUserHandle {
		return connectedUser, nil
	}

	return i.userRepo.FindUser(userToFetch.Handle)
}

//TODO: implement permission management
func (i interactor) EditProfile(connectedUserHandle string, handle string, user *domain.User) error {
	connectedUser, err := i.userRepo.FindUser(connectedUserHandle)
	var userToEdit *domain.User

	if err != nil {
		return domain.ErrConnectedUserNotFound
	}

	if handle == connectedUserHandle {
		userToEdit = connectedUser
	}

	return i.userRepo.UpdateUser(userToEdit.Handle, user)

}

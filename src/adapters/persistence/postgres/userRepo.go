package postgres

import (
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/carrot-systems/cs-user/src/core/usecases"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//TODO: move password to another repo
//TODO: salt it also
//TODO(bug): a deleted user do not release its handle (because of soft deletion but unique keyword don't care for that)
type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Handle      string
	DisplayName string
	Mail        string
	Password    string
}

type userRepo struct {
	db *gorm.DB
}

func (u userRepo) FindIdWithoutCredentials(id string) (*domain.User, error) {
	var user *User

	result := u.db.Where("id = ?", id).First(&user)

	if user == nil || result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return user.toDomain(), nil
}

func (u userRepo) FindHandle(handle string) (*domain.User, error) {
	var user *User

	result := u.db.Where("handle = ?", handle).First(&user)

	if user == nil || result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return user.toDomain(), nil
}

func (u userRepo) FindId(handle string, credentials domain.Credentials) (*domain.User, error) {
	var user *User

	result := u.db.Where("handle = ? AND password = ?", handle, credentials.Password).First(&user)

	if user == nil || result.Error != nil {
		return nil, domain.ErrUserNotFound
	}

	return user.toDomain(), nil
}

func (u userRepo) CreateUser(user domain.UserCreationRequest) error {
	var userPersisted = fromCreationRequest(user)
	id := uuid.New()
	userPersisted.ID = id

	result := u.db.Create(&userPersisted)

	return result.Error
}

func (u userRepo) DeleteUser(handle string) error {
	result := u.db.Where("handle = ?", handle).Delete(&User{})

	return result.Error
}

func (u userRepo) UpdateUser(handle string, user *domain.User) error {
	//TODO: implement
	panic("implement me")
}

func NewUserRepo(db *gorm.DB) usecases.UserRepo {
	return &userRepo{
		db,
	}
}

func (u User) toDomain() *domain.User {
	return &domain.User{
		ID:          u.ID,
		DisplayName: u.DisplayName,
		Handle:      u.Handle,
		Mail:        u.Mail,
	}
}

func fromDomain(f domain.User) User {
	return User{
		DisplayName: f.DisplayName,
		Handle:      f.Handle,
		Mail:        f.Mail,
		ID:          f.ID,
	}
}

func fromCreationRequest(f domain.UserCreationRequest) User {
	user := fromDomain(f.User)
	user.Password = f.Credentials.Password

	return user
}

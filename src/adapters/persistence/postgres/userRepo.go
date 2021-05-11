package postgres

import (
	"errors"
	"github.com/carrot-systems/cs-user/src/core/domain"
	"github.com/carrot-systems/cs-user/src/core/usecases"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string `gorm:"type:uuid;primary_key"`
	Handle      string
	DisplayName string
	Mail        string
	Password    string
}

type userRepo struct {
	db *gorm.DB
}

func (u userRepo) FindHandle(handle string) (*domain.User, error) {
	var user *domain.User

	result := u.db.Where("handle = ?", handle).First(&user)

	if user == nil || result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u userRepo) FindId(id string) (*domain.User, error) {
	var user *domain.User

	result := u.db.Where("id = ?", id).First(&user)

	if user == nil || result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u userRepo) CreateUser(user domain.UserCreationRequest) error {
	var userPersisted = fromCreationRequest(user)
	id := uuid.New().String()
	userPersisted.ID = id

	result := u.db.Create(&userPersisted)

	return result.Error
}

func (u userRepo) DeleteUser(handle string) error {
	u.db.Delete("handle = ?", handle)
	return nil
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

func (u User) toDomain() domain.User {
	return domain.User{
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
		Password:    "",
	}
}

func fromCreationRequest(f domain.UserCreationRequest) User {
	user := fromDomain(f.User)
	user.Password = f.Credentials.Password

	return user
}

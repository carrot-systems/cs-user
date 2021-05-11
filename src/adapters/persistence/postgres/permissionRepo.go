package postgres

import (
	"github.com/carrot-systems/cs-user/src/core/usecases"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID         string `gorm:"type:uuid;primary_key"`
	User       string
	Permission string
	Flag       int
}

type permissionRepo struct {
	db *gorm.DB
}

func (u permissionRepo) FindPermissions(id string) (int, error) {
	panic("implement me")
}

func (u permissionRepo) SetPermissions(id string, permission int) error {
	panic("implement me")
}

func NewPermissionRepo(db *gorm.DB) usecases.PermissionRepo {
	return &permissionRepo{
		db,
	}
}

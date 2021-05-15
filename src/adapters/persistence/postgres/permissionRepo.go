package postgres

import (
	"github.com/carrot-systems/cs-user/src/core/usecases"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID         string `gorm:"type:uuid;primary_key"`
	UsersId    string
	Permission string
	Flag       int
}

type permissionRepo struct {
	db *gorm.DB
}

func (u permissionRepo) FindPermissions(id uuid.UUID, permission string) int {
	var foundPermission Permission

	result := u.db.Where("users_id = ? AND permission = ?", id, permission).FirstOrCreate(&foundPermission)

	//This is a silent fail, we don't want errors when no permissions are set but just saying there are no permissions set
	if result.Error != nil {
		return 0
	}

	return foundPermission.Flag
}

func (u permissionRepo) SetPermissions(id uuid.UUID, permission string, flag int) error {
	//TODO: implem this
	panic("implement me")
}

func NewPermissionRepo(db *gorm.DB) usecases.PermissionRepo {
	return &permissionRepo{
		db,
	}
}

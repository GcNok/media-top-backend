package database

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) dbRepo.LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) GetUserName(id string) []dbEntity.Admin {
	var admins []dbEntity.Admin
	result := r.db.First(&admins, id)
	spew.Dump(result)
	return admins
}

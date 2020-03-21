package database

import (
	dbEntity "github.com/SmartShopping/special-management-backend/domain/entity/database"
	dbRepo "github.com/SmartShopping/special-management-backend/domain/repository/database"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
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

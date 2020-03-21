package database

import (
	dbEntity "github.com/SmartShopping/special-management-backend/domain/entity/database"
)

type LoginRepository interface {
	GetUserName(id string) ([]dbEntity.Admin)
}

package database

import (
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
)

type LoginRepository interface {
	GetUserName(id string) []dbEntity.Admin
}

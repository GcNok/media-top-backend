package database

import (
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
)

type SpecialRepository interface {
	GetArticle() []dbEntity.Special
}

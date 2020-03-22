package database

import (
	dbEntity "github.com/shiji-naoki/media-top-appli-back/domain/entity/database"
)

type SpecialRepository interface {
	GetArticle() []dbEntity.Special
}

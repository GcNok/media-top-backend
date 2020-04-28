package database

import (
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
)

type SpecialRepository interface {
	GetPopularArticles() ([]dbEntity.Special, error)
	GetNewArticles() ([]dbEntity.Special, error)
	GetRecommendArticles() ([]dbEntity.Special, error)
}

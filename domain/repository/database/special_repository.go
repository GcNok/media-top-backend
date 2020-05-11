package database

import (
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
)

type SpecialRepository interface {
	GetPopularArticles(offset int) ([]dbEntity.Special, error)
	GetNewArticles(offset int) ([]dbEntity.Special, error)
	GetRecommendArticles() ([]dbEntity.Special, error)
	GetComparisonArticles(offset int) ([]dbEntity.Special, error)
}

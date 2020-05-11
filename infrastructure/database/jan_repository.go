package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type janRepository struct {
	db *gorm.DB
}

// NewJanRepository コンストラクタ
func NewJanRepository(db *gorm.DB) dbRepo.JanRepository {
	return &janRepository{db: db}
}

func (r *janRepository) GetJanEntity(jan string) (dbEntity.Jan, error) {
	var janEntity dbEntity.Jan
	err := r.db.Where("JAN = ?", jan).First(&janEntity).Error
	if err != nil {
		if err.Error() == "record not found" {
			return dbEntity.Jan{}, nil
		}
		return dbEntity.Jan{}, err
	}
	return janEntity, nil
}

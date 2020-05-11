package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type specialJanRankRepository struct {
	db *gorm.DB
}

// NewSpecialJanRankRepository コンストラクタ
func NewSpecialJanRankRepository(db *gorm.DB) dbRepo.SpecialJanRankRepository {
	return &specialJanRankRepository{db: db}
}

func (r *specialJanRankRepository) GetSpecialJanRanks(id uint64) ([]dbEntity.SpecialJanRank, error) {
	var specialJanRanks []dbEntity.SpecialJanRank
	err := r.db.Where("special_id = ?", id).Find(&specialJanRanks).Error
	if err != nil {
		return []dbEntity.SpecialJanRank{}, err
	}
	return specialJanRanks, nil
}

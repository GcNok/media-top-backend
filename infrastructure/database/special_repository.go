package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type specialRepository struct {
	db *gorm.DB
}

// NewSpecialRepository コンストラクタ
func NewSpecialRepository(db *gorm.DB) dbRepo.SpecialRepository {
	return &specialRepository{db: db}
}

func (r *specialRepository) GetPopularArticles(offset int) ([]dbEntity.Special, error) {
	var specials []dbEntity.Special
	//result := r.db.First(&specials, id)
	err := r.db.
		Order("last_30days_pv desc").
		Offset(offset).
		Limit(10).
		Find(&specials).Error
	if err != nil {
		return []dbEntity.Special{}, err
	}
	return specials, nil
}

func (r *specialRepository) GetNewArticles(offset int) ([]dbEntity.Special, error) {
	var specials []dbEntity.Special
	err := r.db.Order("created desc").
		Offset(offset).
		Limit(10).
		Find(&specials).Error
	if err != nil {
		return []dbEntity.Special{}, err
	}
	return specials, nil
}

func (r *specialRepository) GetRecommendArticles() ([]dbEntity.Special, error) {
	var specials []dbEntity.Special
	err := r.db.Where("id IN ?",
		r.db.Table("special_recommends").
			Select("special_id").
			SubQuery()).
		Find(&specials).Error
	if err != nil {
		return []dbEntity.Special{}, err
	}
	return specials, nil
}

func (r *specialRepository) GetComparisonArticles(offset int) ([]dbEntity.Special, error) {
	var specials []dbEntity.Special
	err := r.db.Where("id IN ?",
		r.db.Table("special_jan_ranks").
			Select("distinct(special_id)").
			SubQuery()).
		Order("last_30days_pv desc").
		Offset(offset).
		Limit(10).
		Find(&specials).Error
	if err != nil {
		return []dbEntity.Special{}, err
	}
	return specials, nil
}
package database

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type specialRepository struct {
	db *gorm.DB
}

func NewSpecialRepository(db *gorm.DB) dbRepo.SpecialRepository {
	return &specialRepository{db: db}
}

func (r *specialRepository) GetArticle() []dbEntity.Special {
	var specials []dbEntity.Special
	//result := r.db.First(&specials, id)
	result := r.db.Limit(3).Find(&specials)
	spew.Dump(result)
	return specials
}

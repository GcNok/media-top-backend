package database

import (
	"github.com/jinzhu/gorm"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type virtualWriterRepository struct {
	db *gorm.DB
}

// NewVirtualWriterRepository コンストラクタ
func NewVirtualWriterRepository(db *gorm.DB) dbRepo.VirtualWriterRepository {
	return &virtualWriterRepository{db: db}
}

func (r *virtualWriterRepository) GetWriterById(id uint64) (dbEntity.VirtualWriter, error) {
	var virtualWriter dbEntity.VirtualWriter
	err := r.db.Where("id = ?", id).First(&virtualWriter).Error
	if err != nil {
		return dbEntity.VirtualWriter{}, err
	}
	return virtualWriter, nil
}

package registry

import (
	"github.com/jinzhu/gorm"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
	"github.com/shiji-naoki/media-top-backend/infrastructure/database"
)

type (
	Repository interface {
		JanRepository() dbRepo.JanRepository
		SpecialRepository() dbRepo.SpecialRepository
		SpecialJanRankRepository() dbRepo.SpecialJanRankRepository
		VirtualWriterRepository() dbRepo.VirtualWriterRepository
	}

	repositoryImpl struct {
		dbConn *gorm.DB
	}
)

func NewRepository() Repository {
	return &repositoryImpl{dbConn: database.MysqlConnection()}
}

func (r *repositoryImpl) JanRepository() dbRepo.JanRepository {
	return database.NewJanRepository(r.dbConn)
}

func (r *repositoryImpl) SpecialRepository() dbRepo.SpecialRepository {
	return database.NewSpecialRepository(r.dbConn)
}

func (r *repositoryImpl) SpecialJanRankRepository() dbRepo.SpecialJanRankRepository {
	return database.NewSpecialJanRankRepository(r.dbConn)
}

func (r *repositoryImpl) VirtualWriterRepository() dbRepo.VirtualWriterRepository {
	return database.NewVirtualWriterRepository(r.dbConn)
}

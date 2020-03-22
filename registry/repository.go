package registry

import (
	"github.com/jinzhu/gorm"
	dbRepo "github.com/shiji-naoki/media-top-appli-back/domain/repository/database"
	"github.com/shiji-naoki/media-top-appli-back/infrastructure/database"
)

type (
	Repository interface {
		LoginRepository() dbRepo.LoginRepository
		SpecialRepository() dbRepo.SpecialRepository
	}

	repositoryImpl struct {
		dbConn *gorm.DB
	}
)

func NewRepository() Repository {
	return &repositoryImpl{dbConn: database.MysqlConnection()}
}

func (r *repositoryImpl) LoginRepository() dbRepo.LoginRepository {
	return database.NewLoginRepository(r.dbConn)
}

func (r *repositoryImpl) SpecialRepository() dbRepo.SpecialRepository {
	return database.NewSpecialRepository(r.dbConn)
}

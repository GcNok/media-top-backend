package registry

import (
	dbRepo "github.com/SmartShopping/special-management-backend/domain/repository/database"
	"github.com/SmartShopping/special-management-backend/infrastructure/database"
	"github.com/jinzhu/gorm"
)

type (
	Repository interface {
		LoginRepository() dbRepo.LoginRepository
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

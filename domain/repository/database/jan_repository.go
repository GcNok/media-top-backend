package database

import dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"

type JanRepository interface {
	GetJanEntity(jan string) (dbEntity.Jan, error)
}

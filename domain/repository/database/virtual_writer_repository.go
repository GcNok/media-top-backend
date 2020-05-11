package database

import dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"

type (
	VirtualWriterRepository interface {
		GetWriterById(id uint64) (dbEntity.VirtualWriter, error)
	}
)

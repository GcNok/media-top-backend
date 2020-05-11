package database

import dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"

type SpecialJanRankRepository interface {
	GetSpecialJanRanks(id uint64) ([]dbEntity.SpecialJanRank, error)
}

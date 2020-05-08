package registry

import (
	"github.com/shiji-naoki/media-top-backend/domain/service"
)

type (
	Service interface {
		ArticleService() service.ArticleService
	}
	serviceImpl struct {
		repo Repository
	}
)

func NewService() Service {
	return &serviceImpl{
		repo: NewRepository(),
	}
}

func (r *serviceImpl) ArticleService() service.ArticleService {
	return service.NewArticleService(
		r.repo.VirtualWriterRepository(),
		r.repo.SpecialJanRankRepository(),
		r.repo.JanRepository())
}

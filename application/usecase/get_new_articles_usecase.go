package usecase

import (
	"context"

	"github.com/shiji-naoki/media-top-backend/domain/entity/api/response"
	"github.com/shiji-naoki/media-top-backend/domain/service"

	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type (
	// GetNewArticlesUseCase このクラスで定義するメソッドを定義
	GetNewArticlesUseCase interface {
		Do(ctx context.Context) ([]response.GetArticlesResponse, error)
	}

	// このクラスで使用するクラスを定義
	getNewArticlesUseCase struct {
		specialRepository dbRepo.SpecialRepository
		articleService    service.ArticleService
	}
)

// NewGetNewArticlesUseCase このクラスのインスタンスを返す
func NewGetNewArticlesUseCase(lr dbRepo.SpecialRepository, as service.ArticleService) GetNewArticlesUseCase {
	return &getNewArticlesUseCase{specialRepository: lr, articleService: as}
}

// このクラスのメイン処理
func (r *getNewArticlesUseCase) Do(ctx context.Context) ([]response.GetArticlesResponse, error) {
	specials, err := r.specialRepository.GetNewArticles()
	if err != nil {
		return []response.GetArticlesResponse{}, err
	}
	var articles []response.GetArticlesResponse
	articles, err = r.articleService.CreateResponse(specials)
	if err != nil {
		return []response.GetArticlesResponse{}, err
	}
	return articles, nil
}

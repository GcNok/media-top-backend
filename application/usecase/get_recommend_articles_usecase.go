package usecase

import (
	"context"

	"github.com/shiji-naoki/media-top-backend/domain/entity/api/response"
	"github.com/shiji-naoki/media-top-backend/domain/service"

	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type (
	// GetRecommendArticlesUseCase このクラスで定義するメソッドを定義
	GetRecommendArticlesUseCase interface {
		Do(ctx context.Context) ([]response.GetArticlesResponse, error)
	}

	// このクラスで使用するクラスを定義
	getRecommendArticlesUseCase struct {
		specialRepository dbRepo.SpecialRepository
		articleService    service.ArticleService
	}
)

// NewGetRecommendArticlesUseCase このクラスのインスタンスを返す
func NewGetRecommendArticlesUseCase(sr dbRepo.SpecialRepository, as service.ArticleService) GetRecommendArticlesUseCase {
	return &getRecommendArticlesUseCase{specialRepository: sr, articleService: as}
}

// このクラスのメイン処理
func (r *getRecommendArticlesUseCase) Do(ctx context.Context) ([]response.GetArticlesResponse, error) {
	specials, err := r.specialRepository.GetRecommendArticles()
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

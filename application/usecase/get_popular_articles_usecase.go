package usecase

import (
	"context"

	"github.com/shiji-naoki/media-top-backend/domain/entity/api/response"
	"github.com/shiji-naoki/media-top-backend/domain/service"

	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type (
	// GetPopularArticlesUseCase このクラスで定義するメソッドを定義
	GetPopularArticlesUseCase interface {
		Do(ctx context.Context, offset int) ([]response.GetArticlesResponse, error)
	}

	// このクラスで使用するクラスを定義
	getPopularArticlesUseCase struct {
		specialRepository dbRepo.SpecialRepository
		articleService    service.ArticleService
	}
)

// NewGetPopularArticlesUseCase このクラスのインスタンスを返す
func NewGetPopularArticlesUseCase(sr dbRepo.SpecialRepository, as service.ArticleService) GetPopularArticlesUseCase {
	return &getPopularArticlesUseCase{specialRepository: sr, articleService: as}
}

// このクラスのメイン処理
func (r *getPopularArticlesUseCase) Do(ctx context.Context, offset int) ([]response.GetArticlesResponse, error) {
	specials, err := r.specialRepository.GetPopularArticles(offset)
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

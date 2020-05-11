package usecase

import (
	"context"

	"github.com/shiji-naoki/media-top-backend/domain/entity/api/response"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
	"github.com/shiji-naoki/media-top-backend/domain/service"
)

type (
	// GetComparisonArticlesUseCase このクラスで定義するメソッドを定義
	GetComparisonArticlesUseCase interface {
		Do(ctx context.Context, offset int) ([]response.GetComparisonArticlesResponse, error)
	}

	// このクラスで使用するクラスを定義
	getComparisonArticlesUseCase struct {
		specialRepository dbRepo.SpecialRepository
		articleService    service.ArticleService
	}
)

// NewGetComparisonArticlesUseCase このクラスのインスタンスを返す
func NewGetComparisonArticlesUseCase(sr dbRepo.SpecialRepository, as service.ArticleService) GetComparisonArticlesUseCase {
	return &getComparisonArticlesUseCase{specialRepository: sr, articleService: as}
}

// このクラスのメイン処理
func (r *getComparisonArticlesUseCase) Do(ctx context.Context, offset int) ([]response.GetComparisonArticlesResponse, error) {
	specials, err := r.specialRepository.GetComparisonArticles(offset)
	if err != nil {
		return []response.GetComparisonArticlesResponse{}, err
	}
	var articles []response.GetComparisonArticlesResponse
	articles, err = r.articleService.CreateComparisonArticlesResponse(specials)
	if err != nil {
		return []response.GetComparisonArticlesResponse{}, err
	}
	return articles, nil
}

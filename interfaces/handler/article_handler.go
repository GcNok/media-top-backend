package handler

import (
	"net/http"

	"github.com/shiji-naoki/media-top-backend/application/usecase"
	"github.com/shiji-naoki/media-top-backend/domain/entity/api/response/common"
	"github.com/shiji-naoki/media-top-backend/registry"

	"github.com/gin-gonic/gin"
	log "github.com/shiji-naoki/media-top-backend/infrastructure/logger"
)

type (
	// ArticleHandler このクラスに定義するメソッドを定義する
	ArticleHandler interface {
		GetPopularArticles(c *gin.Context)
		GetNewArticles(c *gin.Context)
		GetRecommendArticles(c *gin.Context)
		GetComparisonArticles(c *gin.Context)
	}

	articleHandler struct {
		repo registry.Repository
		svc  registry.Service
	}
)

// NewArticleHandler このクラスのインスタンスを返す
func NewArticleHandler(repo registry.Repository, svc registry.Service) ArticleHandler {
	return &articleHandler{repo, svc}
}

// 人気記事一覧取得
func (r *articleHandler) GetPopularArticles(c *gin.Context) {
	ctx := c.Request.Context()
	var (
		requestParams = struct {
			Offset int `form:"offset"`
		}{}
	)
	if err := c.ShouldBindQuery(&requestParams); err != nil {
		log.Critf(ctx, "%s", err.Error())
		c.JSON(http.StatusInternalServerError, common.HttpResponse{
			ResultCode: 500,
			Message:    "リクエストに問題があるため更新に失敗しました。",
		})
		return
	}
	articles, err := usecase.
		NewGetPopularArticlesUseCase(
			r.repo.SpecialRepository(),
			r.svc.ArticleService()).
		Do(ctx, requestParams.Offset)
	if err != nil {
		log.Crit(ctx, err.Error())
		c.JSON(http.StatusInternalServerError, common.HttpResponse{
			ResultCode: http.StatusInternalServerError,
			Message:    "人気記事の取得に失敗しました。",
		})
	}
	c.JSON(http.StatusOK, articles)
	return
}

// 新着記事一覧取得
func (r *articleHandler) GetNewArticles(c *gin.Context) {
	ctx := c.Request.Context()
	articles, err := usecase.
		NewGetNewArticlesUseCase(
			r.repo.SpecialRepository(),
			r.svc.ArticleService()).
		Do(ctx)
	if err != nil {
		log.Crit(ctx, err.Error())
		c.JSON(http.StatusInternalServerError, common.HttpResponse{
			ResultCode: http.StatusInternalServerError,
			Message:    "新着記事の取得に失敗しました。",
		})
	}
	c.JSON(http.StatusOK, articles)
	return
}

// 話題記事一覧取得
func (r *articleHandler) GetRecommendArticles(c *gin.Context) {
	ctx := c.Request.Context()
	articles, err := usecase.
		NewGetRecommendArticlesUseCase(
			r.repo.SpecialRepository(),
			r.svc.ArticleService()).
		Do(ctx)
	if err != nil {
		log.Crit(ctx, err.Error())
		c.JSON(http.StatusInternalServerError, common.HttpResponse{
			ResultCode: http.StatusInternalServerError,
			Message:    "話題記事の取得に失敗しました。",
		})
	}
	c.JSON(http.StatusOK, articles)
	return
}

// 徹底比較記事一覧取得
func (r *articleHandler) GetComparisonArticles(c *gin.Context) {
	ctx := c.Request.Context()
	articles, err := usecase.
		NewGetComparisonArticlesUseCase(
			r.repo.SpecialRepository(),
			r.svc.ArticleService()).
		Do(ctx)
	if err != nil {
		log.Crit(ctx, err.Error())
		c.JSON(http.StatusInternalServerError, common.HttpResponse{
			ResultCode: http.StatusInternalServerError,
			Message:    "話題記事の取得に失敗しました。",
		})
	}
	c.JSON(http.StatusOK, articles)
	return
}

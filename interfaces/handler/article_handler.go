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
		GetArticles(c *gin.Context)
	}

	articleHandler struct {
		repo registry.Repository
	}
)

// NewArticleHandler このクラスのインスタンスを返す
func NewArticleHandler(repo registry.Repository) ArticleHandler {
	return &articleHandler{repo}
}

// Article リクエストパラメータに付与されたidの値を元にusecaseを呼び出し、返された値をフロントに返す
func (l *articleHandler) GetArticles(c *gin.Context) {
	ctx := c.Request.Context()
	articles, err := usecase.NewGetPopularArticlesUseCase(l.repo.SpecialRepository()).Do(ctx)
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

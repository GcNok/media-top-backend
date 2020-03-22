package handler

import (
	"net/http"

	"github.com/shiji-naoki/media-top-appli-back/application/usecase"
	"github.com/shiji-naoki/media-top-appli-back/registry"

	"github.com/gin-gonic/gin"
)

type (
	// SpecialHandler このクラスに定義するメソッドを定義する
	SpecialHandler interface {
		GetArticles(c *gin.Context)
	}

	specialHandler struct {
		repo registry.Repository
	}
)

// NewSpecialHandler このクラスのインスタンスを返す
func NewSpecialHandler(repo registry.Repository) SpecialHandler {
	return &specialHandler{repo}
}

// Special リクエストパラメータに付与されたidの値を元にusecaseを呼び出し、返された値をフロントに返す
func (l *specialHandler) GetArticles(c *gin.Context) {
	ctx := c.Request.Context()
	articles := usecase.NewSpecialUseCase(l.repo.SpecialRepository()).Do(ctx)
	c.JSON(http.StatusOK, articles)
	return
}

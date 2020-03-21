package auth

import (
	
	"github.com/SmartShopping/special-management-backend/registry"
	"github.com/SmartShopping/special-management-backend/application/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// LoginHandler このクラスに定義するメソッドを定義する
	LoginHandler interface {
		Login(c *gin.Context)
	}

	loginHandler struct {
		repo registry.Repository
	}
)

// NewLoginHandler このクラスのインスタンスを返す
func NewLoginHandler(repo registry.Repository) LoginHandler {
	return &loginHandler{repo}
}

// Login リクエストパラメータに付与されたidの値を元にusecaseを呼び出し、返された値をフロントに返す
func (l *loginHandler) Login(c *gin.Context) {
	id := c.Query("id")
	ctx := c.Request.Context()
	name := auth.NewLoginUseCase(l.repo.LoginRepository()).Do(ctx, id)
	c.JSON(http.StatusOK, name)
	return
}

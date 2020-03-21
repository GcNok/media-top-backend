package server

import (
	"github.com/SmartShopping/special-management-backend/interfaces/handler/auth"
	"github.com/SmartShopping/special-management-backend/registry"
	"github.com/gin-gonic/gin"
)

// V1 APIのルーティングを行う
func V1(e *gin.Engine) {
	var (
		// register配下のrepository.goのインスタンスを取得
		repo         = registry.NewRepository()
		// LoginHandlerのインスタンスを取得
		loginHandler = auth.NewLoginHandler(repo)
		// ルーティングするAPIのURLをv1配下にグルーピング
		v1Group      = e.Group("/v1")
	)
	// URLが/v1/loginの場合、login_handler.goのLoginメソッドを実行する
	v1Group.GET("/login", loginHandler.Login)
}

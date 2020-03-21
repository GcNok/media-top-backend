package main

import (
	"github.com/SmartShopping/special-management-backend/interfaces/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// ginのインスタンスを取得
	router := gin.Default()
	// ルーティングを設定
	server.V1(router)
	// サーバー実行
	router.Run()
}

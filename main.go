package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shiji-naoki/media-top-appli-back/interfaces/server"
)

func main() {
	// ginのインスタンスを取得
	router := gin.Default()
	// ルーティングを設定
	server.V1(router)
	// サーバー実行
	router.Run()
}

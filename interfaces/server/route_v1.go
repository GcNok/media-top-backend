package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shiji-naoki/media-top-backend/interfaces/handler"
	"github.com/shiji-naoki/media-top-backend/registry"
)

// V1 APIのルーティングを行う
func V1(e *gin.Engine) {
	var (
		// register配下のrepository.goのインスタンスを取得
		repo           = registry.NewRepository()
		svc            = registry.NewService()
		articleHandler = handler.NewArticleHandler(repo, svc)
		// ルーティングするAPIのURLをv1配下にグルーピング
		v1Group = e.Group("/v1")
	)

	v1Group.GET("/get_recommend_articles", articleHandler.GetRecommendArticles)
	v1Group.GET("/get_popular_articles", articleHandler.GetPopularArticles)
	v1Group.GET("/get_new_articles", articleHandler.GetNewArticles)
	v1Group.GET("/get_comparison_articles", articleHandler.GetComparisonArticles)
}

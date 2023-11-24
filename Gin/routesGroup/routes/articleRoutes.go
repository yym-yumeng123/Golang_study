package routes

import (
	"gin/routesGroup/controllers/article"
	"github.com/gin-gonic/gin"
)

func ArticleRoutesInit(r *gin.Engine) {
	article := article.ArticleController{}
	apiRoutes := r.Group("/article")
	{
		apiRoutes.GET("/", article.Index)
	}
}

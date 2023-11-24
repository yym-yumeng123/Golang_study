package article

import (
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
}

func (a ArticleController) Index(c *gin.Context) {
	articleList := []models.Article{}
	// 查询文章获取对应的分类
	models.DB.Preload("ArticleCate").Find(&articleList)

	c.JSON(http.StatusOK, gin.H{
		"result": articleList,
	})
}

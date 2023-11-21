package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutesInit(r *gin.Engine) {
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个api接口")
		})
		apiRoutes.GET("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个api接口-list")
		})
	}
}

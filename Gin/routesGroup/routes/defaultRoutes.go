package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRoutesInit(r *gin.Engine) {
	defaultRoutes := r.Group("/")
	{
		defaultRoutes.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})
		defaultRoutes.GET("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "列表")
		})
	}
}

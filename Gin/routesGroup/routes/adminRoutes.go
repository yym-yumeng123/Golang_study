package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutesInit(r *gin.Engine) {
	adminRoutes := r.Group("/admin")
	{
		adminRoutes.GET("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "用户列表")
		})
	}
}

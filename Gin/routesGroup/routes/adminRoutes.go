package routes

import (
	"fmt"
	"gin/routesGroup/controllers/admin"
	"gin/routesGroup/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutesInit(r *gin.Engine) {
	user := admin.UserController{}
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middlewares.InitMiddleware)
	{
		// 当前路由中间件
		adminRoutes.GET("/user", func(context *gin.Context) {
			fmt.Println("aaaa")
		}, user.Index)
		adminRoutes.GET("/user/add", user.Add)
		adminRoutes.GET("/user/edit", user.Edit)
		adminRoutes.GET("/user/delete", user.Delete)
		adminRoutes.POST("/user/doUpload", user.DoUpload)
	}
}

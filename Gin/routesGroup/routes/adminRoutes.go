package routes

import (
	"gin/routesGroup/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutesInit(r *gin.Engine) {
	user := admin.UserController{}
	adminRoutes := r.Group("/admin")
	{
		adminRoutes.GET("/user", user.Index)
		adminRoutes.GET("/user/add", user.Add)
		adminRoutes.GET("/user/edit", user.Edit)
	}
}

package routes

import (
	"gin/routesGroup/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutesInit(r *gin.Engine) {
	api := api.AController{}
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/", api.Index)
		apiRoutes.GET("/list", api.List)
	}
}

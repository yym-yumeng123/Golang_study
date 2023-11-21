package routes

import (
	"gin/routesGroup/controllers/itying"
	"github.com/gin-gonic/gin"
)

func DefaultRoutesInit(r *gin.Engine) {
	ity := itying.DefaultController{}
	defaultRoutes := r.Group("/")
	{
		defaultRoutes.GET("/", ity.Index)
		defaultRoutes.GET("/list", ity.List)
	}
}

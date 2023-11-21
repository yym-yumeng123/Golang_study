package main

import (
	"gin/routesGroup/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.AdminRoutesInit(r)
	routes.ApiRoutesInit(r)
	routes.DefaultRoutesInit(r)

	r.Run()
}

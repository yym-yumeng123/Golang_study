package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"success": true,
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "你好, 我是新闻")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "post....")
	})

	r.Run() // 启动一个web 服务
}

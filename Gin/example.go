package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title string
	Desc  string
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLGlob("template/*") // 配置模板路径
	// 配置路由
	r.GET("/ping", func(c *gin.Context) {
		// 响应一个 JSON 类型
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"success": true,
		})
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"msg":     "你好, 我是 json",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		a := &Article{
			Title: "我是一个标题",
			Desc:  "描述",
		}
		c.JSON(http.StatusOK, a)
	})

	// 响应 jsonp 请求 /jsonp?callback=test 解决跨域
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title: "我是一个标题-jsonp",
			Desc:  "描述",
		}
		c.JSONP(http.StatusOK, a)
	})

	r.GET("/news", func(c *gin.Context) {
		// 响应一个 String 类型的
		c.String(200, "你好, 我是新闻")
	})

	r.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "我是一个 xml",
		})
	})

	r.GET("/html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "post....")
	})

	r.Run() // 启动一个web 服务
}

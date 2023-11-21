package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// 自定义时间戳转换日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
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
		news := &Article{
			Title: "标题1",
			Desc:  "我是详情1",
		}
		context.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
			"news":  news,
			"date":  time.Now(),
		})
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "post....")
	})

	r.Run() // 启动一个web 服务
}

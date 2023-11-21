package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `json:"userName" form:"name"`
	Pwd      string `json:"pwd" form:"pwd"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")

	// Get 请求传值
	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
			"page": page,
		})
	})

	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	// 动态路由传值
	r.GET("/list/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, id)
	})

	// post
	// 1. 使用 GET 把页面显示
	r.GET("/user1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{})
	})

	// 2. 使用 post
	r.POST("/add", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"pwd":      pwd,
		})
	})

	// Get 传值绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(user); err == nil {
			// 成功返回结构体
			c.JSON(http.StatusOK, user)
			fmt.Println(user)
		} else {
			// 失败返回错误
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/xml", func(c *gin.Context) {
		xmlSliceData, _ := c.GetRawData() // 获取 c.Request.Body 请求数据
		fmt.Println(xmlSliceData)
		//xml.Unmarshal(xmlSliceData, &article)
	})

	r.Run()
}

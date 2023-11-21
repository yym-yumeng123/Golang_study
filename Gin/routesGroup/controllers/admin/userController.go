package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (u *UserController) Index(c *gin.Context) {
	val, _ := c.Get("username")
	fmt.Println(val, "得到的值")
	v, ok := val.(string)

	if ok {
		c.String(http.StatusOK, "用户列表--"+v)
	} else {
		c.String(http.StatusBadRequest, "获取用户失败")
	}
}

func (u *UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "useradd.html", gin.H{})
}

func (u *UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	// 单文件
	file, err := c.FormFile("face")
	//log.Panicln(file.Filename, "我是log")
	dst := "./static/" + file.Filename
	if err == nil {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename), username)
}

package admin

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type UserController struct {
}

func (u *UserController) Index(c *gin.Context) {
	//val, _ := c.Get("username")
	userList := []models.User{}
	models.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})
}

func (u *UserController) Add(c *gin.Context) {
	user := models.User{
		Username: "张三",
		Age:      19,
		Email:    "2@qq.com",
		AddTime:  12,
	}

	result := models.DB.Create(&user)
	fmt.Println(result)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

// 查询, 修改,  保存
func (u *UserController) Edit(c *gin.Context) {
	user := models.User{Id: 1}
	models.DB.Find(&user)
	user.Username = "yym_编辑"
	models.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 删除
func (u *UserController) Delete(c *gin.Context) {
	user := models.User{}
	models.DB.First(&user).Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "删除成功",
	})
}

func (u *UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	// 单文件
	file, err := c.FormFile("face")
	//log.Panicln(file.Filename, "我是log")
	dst := "./static/" + file.Filename
	if err == nil {
		// 获取后缀名
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg": true,
		}
		fmt.Println(extName, allowExtMap)
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename), username)
}

package itying

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (a *DefaultController) Index(c *gin.Context) {
	// maxAge 参数为 -1 < 0 , 删除 cookie, > 0 表示秒数
	c.SetCookie("username", "李四", 10, "/", "localhost", false, false)
	c.String(http.StatusOK, "首页")
}

func (a *DefaultController) List(c *gin.Context) {
	str, err := c.Cookie("username")
	if err == nil {
		c.String(http.StatusOK, "列表"+str)
	}
}

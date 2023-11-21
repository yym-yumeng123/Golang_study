package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (u *UserController) Index(c *gin.Context) {
	//c.String(http.StatusOK, "用户列表")
	u.success(c)
}

func (u *UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "用户列表-add")
}

func (u *UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "用户列表-edit")
}

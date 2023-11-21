package itying

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (a *DefaultController) Index(c *gin.Context) {
	c.String(http.StatusOK, "首页")
}

func (a *DefaultController) List(c *gin.Context) {
	c.String(http.StatusOK, "列表")
}

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AController struct {
}

func (a *AController) Index(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api接口")
}

func (a *AController) List(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api接口-list")
}

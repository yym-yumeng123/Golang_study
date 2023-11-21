package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseController struct{}

func (b *baseController) success(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}

func (b *baseController) error(c *gin.Context) {
	c.String(http.StatusNotFound, "失败")
}

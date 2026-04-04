package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (bc *BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}

func (bc *BaseController) Fail(c *gin.Context) {
	c.String(http.StatusOK, "失败")
}

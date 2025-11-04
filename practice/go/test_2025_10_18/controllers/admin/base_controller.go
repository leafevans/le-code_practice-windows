package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (*BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}

func (*BaseController) Fail(c *gin.Context) {
	c.String(http.StatusOK, "失败")
}

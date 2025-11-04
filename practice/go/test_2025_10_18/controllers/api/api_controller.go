package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIController struct{}

func (APIController) Index(c *gin.Context) {
	c.String(http.StatusOK, "我是一个 API 接口")
}

func (APIController) UserList(c *gin.Context) {
	c.String(http.StatusOK, "我是一个 API-userlist 接口")
}

func (APIController) PList(c *gin.Context) {
	c.String(http.StatusOK, "我是一个 API-plist 接口")
}

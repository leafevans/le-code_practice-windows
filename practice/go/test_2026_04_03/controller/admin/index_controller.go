package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (ic IndexController) Index(c *gin.Context) {
	c.String(http.StatusOK, "后台首页")
	username, _ := c.Get("username")
	fmt.Println(username)

	v, ok := username.(string)
	if ok {
		c.String(http.StatusOK, "用户列表——"+v)
	} else {
		c.String(http.StatusOK, "用户列表——获取用户失败")
	}
}

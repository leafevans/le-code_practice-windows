package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (uc UserController) Index(c *gin.Context) {
	c.String(http.StatusOK, "用户列表")
}

func (uc UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "用户列表-add")
}

func (uc UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "用户列表-edit")
}

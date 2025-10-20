package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (ic IndexController) Index(c *gin.Context) {
	c.String(http.StatusOK, "后台首页")
}

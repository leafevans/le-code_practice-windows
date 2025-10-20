package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebController struct{}

func (wc WebController) Index(c *gin.Context) {
	c.String(http.StatusOK, "首页")
}

func (wc WebController) News(c *gin.Context) {
	c.String(http.StatusOK, "新闻")
}

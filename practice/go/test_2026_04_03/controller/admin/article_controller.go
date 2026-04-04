package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (ac ArticleController) Index(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表")
}

func (ac ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-add")
}

func (ac ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-edit")
}

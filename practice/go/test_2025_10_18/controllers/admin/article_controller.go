package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (ArticleController) Index(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表")
}

func (ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-add")
}

func (ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-edit")
}

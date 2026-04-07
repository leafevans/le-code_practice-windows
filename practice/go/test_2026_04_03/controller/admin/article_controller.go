package admin

import (
	"net/http"
	"test_2026_04_03/models"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (ac ArticleController) Index(c *gin.Context) {
	var articleCateList []models.ArticleCate
	models.DB.Preload("Articles").Where("id > 0").Offset(1).Limit(1).Find(&articleCateList)
	c.JSON(http.StatusOK, gin.H{
		"result": articleCateList,
	})
}

func (ac ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-add")
}

func (ac ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-edit")
}

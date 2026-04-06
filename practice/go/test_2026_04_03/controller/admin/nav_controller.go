package admin

import (
	"net/http"
	"test_2026_04_03/models"

	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (nc NavController) Index(c *gin.Context) {
	navs := []models.Nav{}
	models.DB.Distinct("status").Find(&navs)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"navs":    navs,
	})
}

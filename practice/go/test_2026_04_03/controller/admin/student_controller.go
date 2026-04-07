package admin

import (
	"net/http"
	"test_2026_04_03/models"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	BaseController
}

func (sc StudentController) Index(c *gin.Context) {
	students := []models.Student{}
	models.DB.Preload("Lessons").Where("id = ?", 1).Find(&students)
	c.JSON(http.StatusOK, students)
}

package admin

import (
	"net/http"
	"test_2026_04_03/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LessonController struct {
	BaseController
}

func (lc LessonController) Index(c *gin.Context) {
	lessons := []models.Lesson{}
	models.DB.Preload("Students", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).Find(&lessons)
	c.JSON(http.StatusOK, lessons)
}

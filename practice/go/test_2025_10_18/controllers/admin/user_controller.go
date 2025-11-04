package admin

import (
	"context"
	"fmt"
	"net/http"
	"test_2025_10_18/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	BaseController
}

func (UserController) Index(c *gin.Context) {
	ctx := context.Background()
	users, err := gorm.G[models.User](models.DB).Find(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (UserController) Add(c *gin.Context) {
	user := models.User{
		Username: "大狗叫",
		Age:      22,
		Email:    "gin@gorm.com",
		AddTime:  int(models.GetUnix()),
	}
	ctx := context.Background()
	result := gorm.WithResult()
	err := gorm.G[models.User](models.DB, result).Create(ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)
	c.String(http.StatusOK, "增加用户")
}

func (UserController) Edit(c *gin.Context) {
	ctx := context.Background()
	res, err := gorm.G[models.User](models.DB).Where("id = ?", 1).Update(ctx, "age", 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println(res)
	c.String(http.StatusOK, "编辑用户")
}

func (UserController) Delete(c *gin.Context) {
	user := models.User{}
	models.DB.Where("username = ?", "叮咚鸡").Delete(&user)
	c.String(http.StatusOK, "删除用户")
}

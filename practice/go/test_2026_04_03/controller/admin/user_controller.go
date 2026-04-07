package admin

import (
	"net/http"
	"path"
	"test_2026_04_03/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

type Result struct {
	Name string
	Age  int
}

func (uc UserController) Index(c *gin.Context) {
	var count int
	models.DB.Raw("SELECT COUNT(*) FROM users").Scan(&count)
	c.JSON(http.StatusOK, gin.H{
		"result": count,
	})
}

func (uc UserController) Edit(c *gin.Context) {
	res := models.DB.Exec("UPDATE users SET username = ? WHERE id = 2", "叮咚鸡")
	c.JSON(http.StatusOK, gin.H{
		"result": res.RowsAffected,
	})
}

func (uc UserController) Add(c *gin.Context) {
	user := models.User{
		Username: "哈基米",
		Age:      18,
		Email:    "hachimi@foxmail.com",
		AddTime:  int(time.Now().Unix()),
	}

	tx := models.DB.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "add 失败")
	}

	tx.Commit()

	c.String(http.StatusOK, "add 成功")
}

func (uc UserController) DoAdd(c *gin.Context) {
	username := c.PostForm("username")
	form, _ := c.MultipartForm()
	files := form.File["avatar[]"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":      "文件上传成功",
		"username": username,
	})
}

func (uc UserController) Delete(c *gin.Context) {
	res := models.DB.Exec("DELETE FROM users WHERE id = ?", 3)
	c.JSON(http.StatusOK, gin.H{
		"result": res.RowsAffected,
	})
}

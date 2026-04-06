package admin

import (
	"fmt"
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
	var res Result
	models.DB.Raw("SELECT username, age FROM users WHERE username = ?", "曼波").Scan(&res)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  res,
	})
}

func (uc UserController) Edit(c *gin.Context) {
	user := models.User{ID: 7}
	models.DB.Find(&user)
	user.Username = "曼波"
	user.Age = 1
	models.DB.Save(&user)
	c.String(http.StatusOK, "Edit")
}

func (uc UserController) Add(c *gin.Context) {
	user := models.User{
		Username: "哈基米",
		Age:      18,
		Email:    "hachimi@foxmail.com",
		AddTime:  int(time.Now().Unix()),
	}

	res := models.DB.Create(&user)
	if res.Error != nil {
		fmt.Println("添加用户失败", res.Error)
		c.String(http.StatusInternalServerError, "添加失败")
		return
	}

	if res.RowsAffected == 1 {
		fmt.Println(user.ID)
	}
	fmt.Println(res.RowsAffected)
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
	models.DB.Where("id > 9").Delete(&models.User{})
	c.String(http.StatusOK, "DeleteAll")
}

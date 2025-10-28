package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (uc UserController) Index(c *gin.Context) {
	c.String(http.StatusOK, "用户列表")
}

func (uc UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/user/add.html", gin.H{})
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
		"message":  "文件上传成功",
		"username": username,
	})
}

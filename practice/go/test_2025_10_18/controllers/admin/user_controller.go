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
	c.HTML(http.StatusOK, "admin/user_add.html", gin.H{})
}

func (uc UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")

	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}

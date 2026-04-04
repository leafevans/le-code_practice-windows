package admin

import (
	"fmt"
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

func (uc UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "用户列表-edit")
}

func (uc UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/user/add.html", gin.H{})
}

func (uc UserController) DoAdd(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	dst := path.Join("./static/upload", file.Filename)
	fmt.Println(dst)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message":  fmt.Sprintf("'%s' uploaded!", file.Filename),
		"username": username,
	})
}

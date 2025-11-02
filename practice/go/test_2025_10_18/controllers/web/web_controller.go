package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WebController struct{}

func (wc WebController) Index(c *gin.Context) {
	c.SetCookie("username", "曼波", 3600, "/", "localhost", false, false)
	c.HTML(http.StatusOK, "web/index.html", gin.H{
		"msg":  "我是一个 msg",
		"time": time.Now().Unix(),
	})
}

func (wc WebController) News(c *gin.Context) {
	username, _ := c.Cookie("username")
	c.String(http.StatusOK, "Username = %v", username)
}

func (wc WebController) Shop(c *gin.Context) {
	username, _ := c.Cookie("username")
	c.String(http.StatusOK, "Cookie = %v", username)
}

func (wc WebController) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "曼波", -1, "/", "localhost", false, false)
	c.String(http.StatusOK, "删除成功")
}

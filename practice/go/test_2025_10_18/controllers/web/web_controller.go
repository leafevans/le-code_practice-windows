package web

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type WebController struct{}

func (wc WebController) Index(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("username", "曼波")
	session.Options(sessions.Options{MaxAge: 3600 * 6})
	session.Save()
	c.HTML(http.StatusOK, "web/index.html", gin.H{
		"msg":  "我是一个 msg",
		"time": time.Now().Unix(),
	})
}

func (wc WebController) News(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(http.StatusOK, "Username = %v", username)
}

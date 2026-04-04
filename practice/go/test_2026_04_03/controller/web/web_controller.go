package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WebController struct{}

func (wc WebController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg":  "我是一个 msg",
		"time": time.Now().Unix(),
	})
}

func (wc WebController) News(c *gin.Context) {
	c.String(http.StatusOK, "新闻")
}

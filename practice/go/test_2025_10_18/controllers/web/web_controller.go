package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WebController struct{}

func (wc WebController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"message": "我是一个 message",
		"time":    time.Now().Unix(),
	})
}

func (wc WebController) News(c *gin.Context) {
	c.String(http.StatusOK, "新闻")
}

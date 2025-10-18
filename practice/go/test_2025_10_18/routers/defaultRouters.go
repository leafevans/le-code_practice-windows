package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻")
		})
	}
}

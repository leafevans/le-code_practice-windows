package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个 API 接口")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个 API-userlist 接口")
		})
		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个 API-plist 接口")
		})
	}
}

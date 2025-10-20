package routers

import (
	"test_2025_10_18/controllers/web"

	"github.com/gin-gonic/gin"
)

func WebRoutersInit(r *gin.Engine) {
	webRouters := r.Group("/")
	{
		webRouters.GET("/", web.WebController{}.Index)
		webRouters.GET("/news", web.WebController{}.News)
	}
}

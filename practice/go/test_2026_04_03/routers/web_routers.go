package routers

import (
	"test_2026_04_03/controller/web"
	"test_2026_04_03/middlewares"

	"github.com/gin-gonic/gin"
)

func WebRoutersInit(r *gin.Engine) {
	webRouters := r.Group("/", middlewares.InitMiddleware)
	{
		webRouters.GET("/", web.WebController{}.Index)
		webRouters.GET("/news", web.WebController{}.News)
	}
}

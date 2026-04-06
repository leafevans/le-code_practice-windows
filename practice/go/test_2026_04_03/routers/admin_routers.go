package routers

import (
	"test_2026_04_03/controller/admin"
	"test_2026_04_03/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.POST("/user/add", admin.UserController{}.Add)
		adminRouters.PUT("/user/edit", admin.UserController{}.Edit)
		adminRouters.POST("/user/doAdd", admin.UserController{}.DoAdd)
		adminRouters.DELETE("/user/delete", admin.UserController{}.Delete)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)

		adminRouters.GET("/nav", admin.NavController{}.Index)
	}
}

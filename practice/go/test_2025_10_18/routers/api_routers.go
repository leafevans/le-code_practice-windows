package routers

import (
	"test_2025_10_18/controllers/api"

	"github.com/gin-gonic/gin"
)

func APIRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.APIController{}.Index)
		apiRouters.GET("/userlist", api.APIController{}.UserList)
		apiRouters.GET("/plist", api.APIController{}.PList)
	}
}

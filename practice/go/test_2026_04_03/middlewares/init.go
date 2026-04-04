package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println("路由分组中间件")
	c.Set("username", "哈基米")
	c.Next()
}

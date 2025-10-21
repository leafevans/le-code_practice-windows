package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "哈基米")

	cCp := c.Copy()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

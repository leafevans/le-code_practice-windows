package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个路由引擎
	r := gin.Default()

	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "Hello Gin!")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "我是测试页面")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是 Post 请求，主要用于增加数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "我是 Put 请求，主要用于编辑数据")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "我是 Delete 请求，主要用于删除数据")
	})

	// 启动一个 Web 服务，默认在 0.0.0.0:8080 启动服务
	r.Run()
}

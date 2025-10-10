package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认路由引擎
	r := gin.Default()
	r.GET("/")
}

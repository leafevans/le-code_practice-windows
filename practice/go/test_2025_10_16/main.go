// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// 创建一个路由引擎
// 	r := gin.Default()

// 	// 配置路由
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "值：%v", "Hello Gin!")
// 	})

// 	r.GET("/news", func(c *gin.Context) {
// 		c.String(200, "我是新闻页面")
// 	})

// 	r.GET("/test", func(c *gin.Context) {
// 		c.String(200, "我是测试页面")
// 	})

// 	r.POST("/add", func(c *gin.Context) {
// 		c.String(200, "我是 Post 请求，主要用于增加数据")
// 	})

// 	r.PUT("/edit", func(c *gin.Context) {
// 		c.String(200, "我是 Put 请求，主要用于编辑数据")
// 	})

// 	r.DELETE("/delete", func(c *gin.Context) {
// 		c.String(200, "我是 Delete 请求，主要用于删除数据")
// 	})

//		// 启动一个 Web 服务，默认在 0.0.0.0:8080 启动服务
//		r.Run()
//	}
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		// 创建一个默认路由引擎
//		r := gin.Default()
//		// 配置路由
//		r.GET("/select", func(c *gin.Context) {
//			c.String(http.StatusOK, "Get")
//		})
//		r.POST("/create", func(c *gin.Context) {
//			c.String(http.StatusOK, "Post")
//		})
//		r.PUT("/update", func(c *gin.Context) {
//			c.String(http.StatusOK, "Put")
//		})
//		r.DELETE("/delete", func(c *gin.Context) {
//			c.String(http.StatusOK, "Delete")
//		})
//		r.GET("/news", func(c *gin.Context) {
//			aid := c.Query("aid")
//			c.String(200, "aid = %s", aid)
//		})
//		r.GET("/user/:uid", func(c *gin.Context) {
//			uid := c.Param("uid")
//			c.String(200, "userID = %s", uid)
//		})
//		// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
//		r.Run()
//	}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是后台数据",
			"price": 20,
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "首页")
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello Gin!",
		})
	})
	r.GET("/struct", func(c *gin.Context) {
		a := &Article{
			"title", "desc", "content",
		}
		c.JSON(http.StatusOK, a)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			"title-jsonp", "desc", "content",
		}
		c.JSONP(http.StatusOK, a)
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello XML!",
		})
	})
	r.GET("/struct-xml", func(c *gin.Context) {
		a := &Article{
			"title", "desc", "content",
		}
		c.XML(http.StatusOK, a)
	})

	r.Run()
}

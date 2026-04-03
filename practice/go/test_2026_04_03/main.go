// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Article struct {
// 	Title   string `json:"title" xml:"title"`
// 	Desc    string `json:"desc" xml:"desc"`
// 	Content string `json:"content" xml:"content"`
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "值：%v", "首页")
// 	})

// 	r.GET("/select", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Get")
// 	})

// 	r.POST("/create", func(c *gin.Context) {
// 		c.String(http.StatusCreated, "Post")
// 	})

// 	r.DELETE("/delete", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Delete")
// 	})

// 	r.GET("/news", func(c *gin.Context) {
// 		aid := c.Query("aid")
// 		c.String(http.StatusOK, "aid = %s", aid)
// 	})

// 	r.GET("/user/:uid", func(c *gin.Context) {
// 		uid := c.Param("uid")
// 		c.String(http.StatusOK, "userID = %s", uid)
// 	})

// 	r.GET("/json", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": true,
// 			"message": "Hello Gin!",
// 		})
// 	})

// 	r.GET("/struct", func(c *gin.Context) {
// 		a := &Article{
// 			Title:   "title",
// 			Desc:    "desc",
// 			Content: "content",
// 		}
// 		c.JSON(http.StatusOK, a)
// 	})

// 	r.GET("/jsonp", func(c *gin.Context) {
// 		a := &Article{
// 			Title:   "title",
// 			Desc:    "desc",
// 			Content: "content",
// 		}
// 		c.JSONP(http.StatusOK, a)
// 	})

// 	r.GET("/xml", func(c *gin.Context) {
// 		c.XML(http.StatusOK, gin.H{
// 			"success": true,
// 			"message": "Hello XML!",
// 		})
// 	})

// 	r.GET("/struct-xml", func(c *gin.Context) {
// 		a := &Article{
// 			"title", "desc", "content",
// 		}
// 		c.XML(http.StatusOK, a)
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		r := gin.Default()
//		r.LoadHTMLGlob("templates/*")
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "index.html", gin.H{
//				"title": "Main Website",
//			})
//		})
//		r.Run()
//	}
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		r := gin.Default()
//		r.LoadHTMLGlob("templates/**/*")
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "default/index.html", gin.H{
//				"title": "前台首页",
//			})
//		})
//		r.GET("/admin", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "admin/index.html", gin.H{
//				"titel": "前台首页",
//			})
//		})
//	}
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type UserInfo struct {
// 	Name   string
// 	Gender string
// 	Age    int
// }

//	func main() {
//		r := gin.Default()
//		r.LoadHTMLGlob("templates/**/*")
//		user := &UserInfo{
//			Name:   "哈基米",
//			Gender: "女",
//			Age:    18,
//		}
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "default/index.html", gin.H{
//				"title": "前台首页",
//				"user":  user,
//			})
//		})
//		r.Run()
//	}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	user := &UserInfo{
		Name:   "哈基米",
		Gender: "女",
		Age:    18,
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "前台首页",
			"user":  user,
			"hobby": []string{"吃饭", "睡觉", "敲代码"},
		})
	})
	r.Run()
}

// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Article struct {
// 	Title   string
// 	Content string
// }

//	func main() {
//		r := gin.Default()
//		r.LoadHTMLGlob("templates/**/*")
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "default/index.html", gin.H{
//				"title":   "首页",
//				"message": "信息",
//				"score":   89,
//				"hobby":   []string{"吃饭", "睡觉", "写代码"},
//				"newslist": []any{
//					&Article{
//						Title:   "新闻标题1",
//						Content: "新闻详情1",
//					}, &Article{
//						Title:   "新闻标题2",
//						Content: "新闻详情2",
//					},
//				},
//				"testslice": []string{},
//				"news": &Article{
//					Title:   "新闻标题",
//					Content: "新闻详情",
//				},
//			})
//		})
//		r.GET("/news", func(c *gin.Context) {
//			news := &Article{
//				Title:   "新闻标题",
//				Content: "新闻详情",
//			}
//			c.HTML(http.StatusOK, "default/news.html", gin.H{
//				"title": "新闻页面",
//				"news":  news,
//			})
//		})
//		r.GET("/admin", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "admin/news.html", gin.H{
//				"title": "后台首页",
//			})
//		})
//		r.GET("/admin/news", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "admin/news.html", gin.H{
//				"title": "后台新闻页面",
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
//				"title": "后台首页",
//			})
//		})
//		r.Run()
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
//		// user := &UserInfo{
//		// 	Name:   "哈基米",
//		// 	Gender: "女",
//		// 	Age:    18,
//		// }
//		// r.GET("/", func(c *gin.Context) {
//		// 	c.HTML(http.StatusOK, "default/index.html", gin.H{
//		// 		"title": "前台首页",
//		// 		"user":  user,
//		// 	})
//		// })
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "default/index.html", gin.H{
//				"hobby": []string{"吃饭", "睡觉", "敲代码"},
//			})
//		})
//		r.Run()
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"text/template"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type UserInfo struct {
// 	Name   string
// 	Gender string
// 	Age    int
// }

// func UnixToTime(timestamp int64) string {
// 	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
// }

// func Println(str1 string, str2 string) string {
// 	fmt.Println(str1, str2)
// 	return str1 + str2
// }

//	func main() {
//		r := gin.Default()
//		r.Static("/static", "./static")
//		r.SetFuncMap(template.FuncMap{
//			"UnixToTime": UnixToTime,
//			"Println":    Println,
//		})
//		r.LoadHTMLGlob("templates/**/*")
//		user := &UserInfo{
//			Name:   "曼波",
//			Gender: "男",
//			Age:    18,
//		}
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "default/index.html", gin.H{
//				"title": "前台页面",
//				"user":  user,
//				"date":  time.Now().Unix(),
//				"hobby": []string{"吃饭", "睡觉", "敲代码"},
//			})
//		})
//		r.Run()
//	}
// package main

// import (
// 	"net/http"
// 	"text/template"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func UnixToTime(timestamp int64) string {
// 	return time.Unix(timestamp, 0).Format("2006-01-02 15:03:04")
// }

// func main() {
// 	r := gin.Default()

// 	r.SetFuncMap(template.FuncMap{
// 		"UnixToTime": UnixToTime,
// 	})

// 	r.LoadHTMLGlob("templates/**/*")

// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "default/index.html", gin.H{
// 			"title": "前台首页",
// 			"now":   time.Now().Unix(),
// 		})
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.Static("/xxxx", "./static")

// 	r.LoadHTMLGlob("templates/**/*")

// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "default/index.html", gin.H{
// 			"title": "前台首页",
// 		})
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type UserInfo struct {
// 	Username string `json:"username" form:"username"`
// 	Password string `json:"password" form:"password"`
// }

// type Article struct {
// 	Title   string `json:"title" xml:"title"`
// 	Content string `json:"content" xml:"content"`
// }

// func main() {
// 	r := gin.Default()

// 	r.LoadHTMLGlob("templates/**/*")

// 	r.GET("/article", func(c *gin.Context) {
// 		id := c.DefaultQuery("id", "1")

// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": "新闻详情",
// 			"id":  id,
// 		})
// 	})

// 	r.GET("/user", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "default/user.html", gin.H{})
// 	})

// 	r.POST("/doAddUser1", func(c *gin.Context) {
// 		username := c.PostForm("username")
// 		password := c.PostForm("password")
// 		age := c.DefaultPostForm("age", "18")
// 		c.JSON(http.StatusOK, gin.H{
// 			"username": username,
// 			"password": password,
// 			"age":      age,
// 		})
// 	})

// 	r.GET("/getUser", func(c *gin.Context) {
// 		user := &UserInfo{}
// 		if err := c.ShouldBind(user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"err": err.Error(),
// 			})
// 		} else {
// 			fmt.Printf("%#v", user)
// 			c.JSON(http.StatusOK, user)
// 		}
// 	})

// 	r.POST("/doAddUser2", func(c *gin.Context) {
// 		user := &UserInfo{}
// 		if err := c.ShouldBind(user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"err": err.Error(),
// 			})
// 		} else {
// 			c.JSON(http.StatusOK, user)
// 		}
// 	})

// 	r.POST("/xml", func(c *gin.Context) {
// 		xmlSliceDate, _ := c.GetRawData()

// 		article := &Article{}

// 		if err := xml.Unmarshal(xmlSliceDate, article); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"err": err.Error(),
// 			})
// 		} else {
// 			c.JSON(http.StatusOK, article)
// 		}
// 	})

//		r.Run()
//	}
package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:03:04")
}

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	r.LoadHTMLGlob("templates/**/*")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		fmt.Println("aaa")
	}, func(c *gin.Context) {
		c.String(http.StatusOK, "首页")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "新闻页面")
	})

	r.Run()
}

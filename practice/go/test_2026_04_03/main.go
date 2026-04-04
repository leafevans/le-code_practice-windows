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
//				"hobby": []string{"吃饭", "睡觉", "敲代码"},
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
//		user := &UserInfo{
//			Name:   "曼波",
//			Gender: "男",
//			Age:    18,
//		}
//		r.GET("/", func(c *gin.Context) {
//			c.HTML(http.StatusOK, "default/index.html", gin.H{
//				"user":  user,
//				"hobby": []string{"吃饭", "睡觉", "敲代码"},
//			})
//		})
//		r.Run()
//	}
// package main

// import (
// 	"html/template"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func UnixToTome(timestamp int64) string {
// 	return time.Unix(timestamp, 0).Format("2006-01-02 15:03:04")
// }

// func main() {
// 	r := gin.Default()

// 	r.SetFuncMap(template.FuncMap{
// 		"UnixToTime": UnixToTome,
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

// 	r.Static("/static", "./static")

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
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.Static("/static", "./static")

// 	r.LoadHTMLGlob("templates/**/*")

// 	r.GET("/user", func(c *gin.Context) {
// 		uid := c.Query("uid")
// 		page := c.DefaultQuery("page", "0")
// 		c.String(http.StatusOK, "uid = %v, page = %v", uid, page)
// 	})

// 	r.GET("/user/:uid", func(c *gin.Context) {
// 		uid := c.Param("uid")
// 		c.String(http.StatusOK, "userID = %s", uid)
// 	})

// 	r.GET("/addUser", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "default/add_user.html", gin.H{})
// 	})

// 	r.POST("/doAddUser", func(c *gin.Context) {
// 		username := c.PostForm("username")
// 		password := c.PostForm("password")
// 		age := c.DefaultPostForm("age", "20")

// 		c.JSON(http.StatusOK, gin.H{
// 			"username": username,
// 			"password": password,
// 			"age":      age,
// 		})
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type UserInfo struct {
// 	Username string `form:"username" json:"username"`
// 	Password string `form:"password" json:"password"`
// }

// type Article struct {
// 	Title   string `xml:"title"`
// 	Content string `xml:"content"`
// }

// func main() {
// 	r := gin.Default()

// 	r.LoadHTMLGlob("templates/**/*")

// 	r.Static("/static", "./static")

// 	r.GET("/", func(c *gin.Context) {
// 		userInfo := &UserInfo{}
// 		if err := c.ShouldBind(userInfo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 		} else {
// 			c.JSON(http.StatusOK, userInfo)
// 		}
// 	})

// 	r.GET("/addUser", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "default/add_user.html", gin.H{})
// 	})

// 	r.POST("/doLogin", func(c *gin.Context) {
// 		userInfo := &UserInfo{}
// 		if err := c.ShouldBind(userInfo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		} else {
// 			c.JSON(http.StatusOK, userInfo)
// 		}
// 	})

// 	r.GET("/xml", func(c *gin.Context) {
// 		a := &Article{}
// 		if err := c.ShouldBindXML(a); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		} else {
// 			c.JSON(http.StatusOK, a)
// 		}
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

// 	defaultRouters := r.Group("/")
// 	{
// 		defaultRouters.GET("/", func(c *gin.Context) {
// 			c.String(http.StatusOK, "首页")
// 		})
// 		defaultRouters.GET("/news", func(c *gin.Context) {
// 			c.String(http.StatusOK, "新闻")
// 		})
// 	}

// 	apiRouters := r.Group("/api")
// 	{
// 		apiRouters.GET("/", func(c *gin.Context) {
// 			c.String(http.StatusOK, "我是一个 API 接口")
// 		})
// 		apiRouters.GET("/userlist", func(c *gin.Context) {
// 			c.String(http.StatusOK, "我是一个 API-userlist 接口")
// 		})
// 		apiRouters.GET("/plist", func(c *gin.Context) {
// 			c.String(http.StatusOK, "我是一个 API-plist 接口·")
// 		})
// 	}

// 	adminRouters := r.Group("/admin")
// 	{
// 		adminRouters.GET("/", func(c *gin.Context) {
// 			c.String(http.StatusOK, "后台首页")
// 		})
// 		adminRouters.GET("/user", func(c *gin.Context) {
// 			c.String(http.StatusOK, "用户列表")
// 		})
// 		adminRouters.GET("/article", func(c *gin.Context) {
// 			c.String(http.StatusOK, "新闻列表")
// 		})
// 	}

//		r.Run()
//	}
// package main

// import (
// 	"test_2026_04_03/routers"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	routers.WebRoutersInit(r)
// 	routers.APIRoutersInit(r)
// 	routers.AdminRoutersInit(r)

//		r.Run()
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func initMiddleware(c *gin.Context) {
// 	fmt.Println("我是中间件")
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/", initMiddleware, func(c *gin.Context) {
// 		c.String(http.StatusOK, "首页——中间件演示")
// 	})

// 	r.GET("/news", initMiddleware, func(c *gin.Context) {
// 		c.String(http.StatusOK, "首页——中间件演示")
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func initMiddleware(c *gin.Context) {
// 	start := time.Now().UnixNano()
// 	fmt.Println("1-执行中，中间件")
// 	c.Next()
// 	end := time.Now().UnixNano()
// 	fmt.Printf("3-程序执行完成，计算时间：%d\n", end-start)
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/", initMiddleware, func(c *gin.Context) {
// 		fmt.Println("2-执行首页返回数据")
// 		c.String(http.StatusOK, "首页——中间件演示")
// 	})

// 	r.GET("/news", initMiddleware, func(c *gin.Context) {
// 		c.String(http.StatusOK, "新闻页面——中间件演示")
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func initMiddlewareOne(c *gin.Context) {
// 	fmt.Println("1-执行中，中间件-One")
// 	c.Next()
// 	fmt.Println("3-执行中，中间件-One")
// }

// func initMiddlewareTwo(c *gin.Context) {
// 	fmt.Println("1-执行中，中间件-Two")
// 	c.Abort()
// 	fmt.Println("3-执行中，中间件-Two")
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
// 		fmt.Println("2-执行路由中的程序")
// 		c.String(http.StatusOK, "首页——中间件演示")
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func initMiddleware(c *gin.Context) {
// 	fmt.Println("全局中间件，通过 r.Use 配置")
// }

// func main() {
// 	r := gin.Default()

// 	r.Use(initMiddleware)

// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "首页——中间件演示")
// 	})

// 	r.GET("/news", func(c *gin.Context) {
// 		c.String(http.StatusOK, "新闻页面——中间件演示")
// 	})

//		r.Run()
//	}
// package main

// import (
// 	"html/template"
// 	"test_2026_04_03/models"
// 	"test_2026_04_03/routers"

// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/zap"
// )

// func main() {
// 	logger, _ := zap.NewProduction()
// 	defer logger.Sync()
// 	zap.ReplaceGlobals(logger)

// 	r := gin.Default()

// 	r.SetFuncMap(template.FuncMap{
// 		"UnixToTime": models.UnixToTime,
// 	})

// 	r.LoadHTMLGlob("templates/**/*")

// 	routers.AdminRoutersInit(r)
// 	routers.WebRoutersInit(r)

//		r.Run()
//	}
package main

import (
	"test_2026_04_03/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")

	r.Static("/static", "./static")

	routers.WebRoutersInit(r)

	routers.APIRoutersInit(r)

	routers.AdminRoutersInit(r)

	r.Run()
}

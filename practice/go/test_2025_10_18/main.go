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

// 	r.LoadHTMLGlob("templates/*")

// 	r.GET("/user", func(c *gin.Context) {
// 		uid := c.Query("uid")
// 		page := c.DefaultQuery("page", "0")
// 		c.String(http.StatusOK, "uid = %v, page = %v", uid, page)
// 	})

// 	r.GET("/user/:uid", func(c *gin.Context) {
// 		uid := c.Param("uid")
// 		c.String(http.StatusOK, "userID = %v", uid)
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

// 	r.POST("/doLogin", func(c *gin.Context) {
// 		userInfo := &UserInfo{}
// 		if err := c.ShouldBind(userInfo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		} else {
// 			c.JSON(http.StatusOK, userInfo)
// 		}
// 	})

// 	r.POST("/xml", func(c *gin.Context) {
// 		a := &Article{}
// 		if err := c.ShouldBindXML(a); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		} else {
// 			c.JSON(http.StatusOK, a)
// 		}
// 	})

// 	r.Run()
// }

// package main

// import (
// 	"test_2025_10_18/routers"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	routers.DefaultRoutersInit(r)
// 	routers.APIRoutersInit(r)
// 	routers.AdminRoutersInit(r)

// 	r.Run()
// }

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
// 			c.String(http.StatusOK, "我是一个 API-plist 接口")
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
// 	r.Run()
// }

// package main

// import (
// 	"test_2025_10_18/routers"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.Static("/static", "./static")

// 	routers.WebRoutersInit(r)
// 	routers.APIRoutersInit(r)
// 	routers.AdminRoutersInit(r)

// 	r.Run()
// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initMiddlewareOne(c *gin.Context) {
	fmt.Println("1-执行中，中间件-One")
	c.Next()
	fmt.Println("3-执行中，中间件-One")
}

func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("1-执行中，中间件-Two")
	c.Next()
	fmt.Println("3-执行中，中间件-Two")
}

func main() {
	r := gin.Default()

	r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
		fmt.Println("2-执行路由中的程序")
		c.String(http.StatusOK, "首页——中间件演示")
	})

	r.Run()
}

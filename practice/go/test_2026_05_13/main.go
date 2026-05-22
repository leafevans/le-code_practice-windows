// package main

// import (
// 	"context"
// 	"time"

// 	"github.com/cloudwego/hertz/pkg/app"
// 	"github.com/cloudwego/hertz/pkg/app/server"
// 	"github.com/cloudwego/hertz/pkg/common/utils"
// 	"github.com/cloudwego/hertz/pkg/network/standard"
// 	"github.com/cloudwego/hertz/pkg/protocol/consts"
// )

// func authMiddleware(c context.Context, ctx *app.RequestContext) {
// 	token := ctx.GetHeader("Authorization")
// 	if len(token) == 0 {
// 		ctx.JSON(consts.StatusUnauthorized, utils.H{"error": "unauthorized"})
// 		ctx.Abort()
// 		return
// 	}
// 	ctx.Next(c)
// }

// func main() {
// 	h := server.Default(
// 		server.WithHostPorts(":8080"),
// 		server.WithReadTimeout(10*time.Second),
// 		server.WithWriteTimeout(10*time.Second),
// 		server.WithTransport(standard.NewTransporter),
// 	)

// 	h.GET("/user/:id", func(c context.Context, ctx *app.RequestContext) {
// 		userID := ctx.Param("id")
// 		ctx.JSON(consts.StatusOK, utils.H{
// 			"user_id": userID,
// 		})
// 	})

// 	h.POST("/user", func(c context.Context, ctx *app.RequestContext) {
// 		ctx.JSON(consts.StatusOK, utils.H{
// 			"message": "User created",
// 		})
// 	})

// 	h.PUT("/user/:id", func(c context.Context, ctx *app.RequestContext) {
// 		ctx.JSON(consts.StatusOK, utils.H{
// 			"message": "User updated",
// 			"id":      ctx.Param("id"),
// 		})
// 	})

// 	h.DELETE("/user/:id", func(c context.Context, ctx *app.RequestContext) {
// 		ctx.JSON(consts.StatusOK, utils.H{
// 			"message": "User deleted",
// 		})
// 	})

//		h.Spin()
//	}
// package main

// import (
// 	"context"
// 	"strconv"

// 	"github.com/cloudwego/hertz/pkg/app"
// 	"github.com/cloudwego/hertz/pkg/app/server"
// 	"github.com/cloudwego/hertz/pkg/protocol/consts"
// )

// type User struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

//	func main() {
//		h := server.Default()
//		h.GET("/users/:id", func(c context.Context, ctx *app.RequestContext) {
//			id, _ := strconv.Atoi(ctx.Param("id"))
//			user := User{
//				ID:    id,
//				Name:  "John Doe",
//				Email: "john@example.com",
//			}
//			ctx.JSON(consts.StatusOK, user)
//		})
//	}
// package main

// import (
// 	"context"

// 	"github.com/cloudwego/hertz/pkg/app"
// 	"github.com/cloudwego/hertz/pkg/app/server"
// 	"github.com/cloudwego/hertz/pkg/common/utils"
// 	"github.com/cloudwego/hertz/pkg/protocol/consts"
// )

//	func main() {
//		h := server.Default(server.WithHandleMethodNotAllowed(true))
//		h.GET("/api/resume/:id", func(ctx context.Context, c *app.RequestContext) {
//			c.JSON(consts.StatusOK, utils.H{
//				"message": "GET 请求成功",
//				"id":      c.Param("id"),
//			})
//		})
//		h.POST("/api/resume", func(ctx context.Context, c *app.RequestContext) {
//			c.JSON(consts.StatusOK, utils.H{
//				"message": "POST 请求成功（创建简历）",
//			})
//		})
//		h.NoRoute(func(ctx context.Context, c *app.RequestContext) {
//			c.JSON(consts.StatusNotFound, utils.H{
//				"code":    404,
//				"message": "接口路径不存在",
//				"path":    string(c.Request.URI().Path()),
//			})
//		})
//		h.NoMethod(func(ctx context.Context, c *app.RequestContext) {
//			c.JSON(consts.StatusMethodNotAllowed, utils.H{
//				"code":    405,
//				"message": "不支持的请求方法",
//				"path":    string(c.Request.URI().Path()),
//				"method":  string(c.Request.Method()),
//			})
//		})
//		h.Spin()
//	}
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
	h := server.Default()
	h.GET("/questions/tags", func(ctx context.Context, c *app.RequestContext) {
		tagBytes := c.QueryArgs().PeekAll("tag")
		tags := make([]string, len(tagBytes))
		for i, b := range tagBytes {
			tags[i] = string(b)
		}
		questions := mockQuestionsByTags(tags)
		c.JSON(200, utils.H{
			"tags":      tags,
			"questions": questions,
		})
	})
	h.Spin()
}

func mockQuestionsByTags(tags []string) []string {
	if len(tags) == 0 {
		return []string{"请提供至少一个 tag 参数"}
	}
	result := []string{}
	for _, t := range tags {
		result = append(result, "关于 "+t+" 的面试题示例")
	}
	return result
}

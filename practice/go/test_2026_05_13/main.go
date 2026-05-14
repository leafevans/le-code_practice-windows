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

// 	h.Spin()
// }
package main

func main() {
	
}

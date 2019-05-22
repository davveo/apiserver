package router

import (
	"apiserver/handler/user"
	"net/http"

	"apiserver/handler/sd"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The user handlers, requiring authentication
	g.POST("/login", user.Login)
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		// 创建用户
		u.POST("", user.Create)
		// 删除用户
		u.DELETE("/:id", user.Delete)
		// 更新用户
		u.PUT("/:id", user.Update)
		// 查询用户
		// TODO: 为什么查看用户要通过用户名???
		u.GET("/:username", user.Get)
		// 查询用户列表
		u.GET("", user.List)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

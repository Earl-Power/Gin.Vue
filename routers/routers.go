package routers

import (
	"github.com/Earl-Power/Gin.Vue/controllers"
	"github.com/Earl-Power/Gin.Vue/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
	// 通过中间件授权保护用户信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controllers.Info)
	return r
}

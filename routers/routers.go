package routers

import (
	"github.com/Earl-Power/Gin.Vue/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controllers.Register)
	return r
}

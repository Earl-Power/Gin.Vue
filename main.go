package main

import (
	"github.com/Earl-Power/Gin.Vue/common"
	"github.com/Earl-Power/Gin.Vue/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	common.InitConfig()
	db := common.InitDB()
	defer db.Clauses()
	//gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	//x := gin.DebugMode
	gin.SetMode(common.SerMode())
	r := gin.Default()
	r = routers.CollectRoute(r)
	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}

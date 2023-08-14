package main

import (
	"github.com/Earl-Power/Gin.Vue/common"
	"github.com/Earl-Power/Gin.Vue/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Clauses()
	r := gin.Default()
	r = routers.CollectRoute(r)
	panic(r.Run())
}

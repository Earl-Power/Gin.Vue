package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func SerMode() string {
	decode := viper.GetString("server.Mode")
	if decode == "0" {
		return gin.ReleaseMode
	} else if decode == "1" {
		return gin.DebugMode
	} else {
		print("Server mode error !")
		os.Exit(-1)
		return ""
	}
	// gin.SetMode(gin.ReleaseMode)

}

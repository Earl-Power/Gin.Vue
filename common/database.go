package common

import (
	"fmt"
	"github.com/Earl-Power/Gin.Vue/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 数据库连接
func InitDB() *gorm.DB {
	//driverType := viper.GetString("datasource.driverType")
	host := viper.GetString("datasource.host")
	dbPort := viper.GetString("datasource.db_port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	//args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
	//	username,
	//	password,
	//	host,
	//	dbPort,
	//	database,
	//	charset,
	//)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username, password, host, dbPort, database, charset)

	//dsn := "root:root@tcp(localhost:3306)/gin_vue?charset=utf8&parseTime=True"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect databases, err:" + err.Error())
	}
	// 当返回错误时，直接忽略错误信息
	_ = db.AutoMigrate(&models.User{})
	DB = db
	return db
}

func GetDb() *gorm.DB {
	return DB
}

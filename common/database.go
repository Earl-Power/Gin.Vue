package common

import (
	"github.com/Earl-Power/Gin.Vue/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 数据库连接
func InitDB() *gorm.DB {
	/*
		driverName := "mysql"
		host := "localhost"
		port := "3306"
		username := "root"
		password := "root"
		database := "gin_vue"
		charset := "utf8"
		args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			username,
			password,
			host,
			port,
			database,
			charset,
		)
	*/
	dsn := "root:root@tcp(localhost:3306)/gin_vue?charset=utf8&parseTime=True"
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

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Name      string
	Telephone string
	Password  string
	ID        int
}

func main() {
	db := InitDB()
	defer db.Clauses()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		// 验证数据

		if len(name) == 0 {
			name = RandomString(10)
		}
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位。"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位。"})
			return
		}
		// 判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在！"})
			return
		}
		// 创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		// 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功！",
		})
		log.Panicln(name, telephone, password)

	})

	panic(r.Run())
}

// RandomString 随机用户名（10位）
func RandomString(n int) string {
	var letters = []byte("abcdefghijklnmopqrstuvwxyz1234567890ABCDEFGHIJKLNMOPQRSTUVWXYZ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
	dsn := "root:root@tcp(localhost:3306)/gin_vue?charset=utf8&parseTime=True&loc=local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect databases, err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

// isTelephoneExist
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

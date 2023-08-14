package controllers

import (
	"github.com/Earl-Power/Gin.Vue/common"
	"github.com/Earl-Power/Gin.Vue/models"
	"github.com/Earl-Power/Gin.Vue/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDb()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 验证数据

	if len(name) == 0 {
		name = util.RandomString(10)
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
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在！"})
		return
	}
	// 创建用户
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功！",
	})
	//log.Panicln(name, telephone, password)

}

// isTelephoneExist
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

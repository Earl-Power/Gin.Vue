package controllers

import (
	"github.com/Earl-Power/Gin.Vue/common"
	"github.com/Earl-Power/Gin.Vue/dto"
	"github.com/Earl-Power/Gin.Vue/models"
	"github.com/Earl-Power/Gin.Vue/responses"
	"github.com/Earl-Power/Gin.Vue/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// isTelephoneExist 手机号查询验证
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// Register 用户注册逻辑
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
		//ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位。"})
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位！"})
		return
	}
	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位。"})
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位！"})
		return
	}
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在！"})
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 200, "msg": "用户已经存在！"})
		return
	}
	// 用户密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误！"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误！"})
		return
	}
	// 创建用户
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	//ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功！"})
	responses.Success(ctx, nil, "注册成功！")
	// log.Panicln(name, telephone, password)
}

// Login 用户登录逻辑
func Login(ctx *gin.Context) {
	DB := common.GetDb()
	var user models.User
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		responses.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位！")
		return
	}

	if len(password) > 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码只可6位"})
		responses.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码只可6位！")
		return
	}
	// 判断手机号是否存在
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		responses.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在!")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		responses.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误!")
		return
	}
	// 返回token
	token, err := common.ReleaseToken(user)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		responses.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常!")
		log.Printf("token generate error: %v", err)
		return
	}
	// 返回结果
	// ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"Token": token}, "msg": "登录成功"})
	responses.Success(ctx, gin.H{"Telephone": user.Telephone, "Token": token}, "登录成功!")

}

// Info 用户信息获取（登录后）
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(models.User))}})
}

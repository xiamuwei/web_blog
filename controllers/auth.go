package controllers

import (
	"net/http"
	"web_blog/global"
	"web_blog/model"
	"web_blog/utils"

	"github.com/gin-gonic/gin"
)

// 注册：信息实例化-加密-写入数据库
func Register(ctx *gin.Context) {
	// 模型绑定
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 密码加密
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error()},
		)
		return
	}

	// 用户密码加密存到数据库
	user.Password = hashedPwd

	// 生成jwt令牌
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

	// 数据库操作
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func Login(ctx *gin.Context) {
	// 获取前端数据
	var info struct {
		Username string `gorm: username`
		Password string `gorm: password`
	}
	if err := ctx.ShouldBindBodyWithJSON(&info); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 比较用户名，密码
	var user model.User
	if err := global.Db.Where("username = ?", info.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if !utils.CheckPassword(user.Password, info.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong Authorization"})
		return
	}

	// 生成jwt
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

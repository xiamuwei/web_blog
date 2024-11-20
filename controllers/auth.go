package controllers

import (
	"net/http"
	"web_blog/model"
	"web_blog/utils"

	"github.com/gin-gonic/gin"
)

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

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

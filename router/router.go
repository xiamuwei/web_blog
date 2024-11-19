package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatus(http.StatusOK)
		})
		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatus(http.StatusOK)
		})
	}
	return r
}

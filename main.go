package main

import (
	"web_blog/config"
	"web_blog/router"
)

func main() {
	// 初始化Config结构体
	config.InitConfig()
	// fmt.Println(config.AppConfig.App.Name)

	// 服务器
	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
}

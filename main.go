package main

import (
	"fmt"
	"web_blog/config"
)

func main() {
	// 初始化Config结构体
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Name)
}

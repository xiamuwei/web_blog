package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host         string
		Port         string
		User         string
		Password     string
		Name         string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	// 读取配置文件
	file, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatalf("Can not Open File %v\n", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("can not read all %v\n", err)
	}
	fmt.Println(string(content))

	// 反序列化
	AppConfig = &Config{}
	yaml.Unmarshal(content, AppConfig)
}

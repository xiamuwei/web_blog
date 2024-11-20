package config

import (
	"fmt"
	"log"
	"time"
	"web_blog/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@@tcp(%v%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can not open gorm engine :%v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("database happened :%v\n", err)
	}
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.Db = db
}

package database

import (
	"fmt"

	"gonews/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMysqlDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("Connected Successfully to Database (Mysql)")

	return db
}

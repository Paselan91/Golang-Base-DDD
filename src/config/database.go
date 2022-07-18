package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func ConnectDB() (*gorm.DB, error) {
	config := ConfigDB{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_CONTAINER_NAME"), // コンテナ名を指定する
		Port:     os.Getenv("MYSQL_PORT"),
		Dbname:   os.Getenv("MYSQL_DATABASE"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)
	log.Println("dsn")
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		hoge := fmt.Sprintf("failed to connect DB. err : %s", err)
		log.Println(hoge)
		return nil, err
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	return db, nil
}

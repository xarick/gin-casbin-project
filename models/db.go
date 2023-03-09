package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {

	DBHost := "localhost"
	DBUserName := "root"
	DBUserPassword := "root"
	DBName := "gin-casib-pr-db"
	DBPort := "5432"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tashkent", DBHost, DBUserName, DBUserPassword, DBName, DBPort)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Panic("Failed to connect to the Database")
	// }

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Role{})
	// DB = db
}

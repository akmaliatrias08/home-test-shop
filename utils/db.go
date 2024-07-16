package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var dbHost, dbUsername, dbPassword, dbName, dbPort = os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func AutoMigrate(from string, dst ...interface{}) {
	err := DB.AutoMigrate(dst...)

	if err != nil {
		panic("[" + from + "] failed to migrate database")
	}

	fmt.Println("[" + from + "] database migrated")
}

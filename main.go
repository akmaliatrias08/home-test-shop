package main

import (
	"log"
	"online-shop-home-test/modules"
	"online-shop-home-test/modules/utils"

	"github.com/joho/godotenv"
)

// @title         Online Shop
// @version       1.0
// @description   A Online Shop Backend Service

// @host          localhost:5000
// @BasePath      /
func main() {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//to init and run the database
	utils.InitDatabase()

	//run routes from main modules
	modules.Run()
}

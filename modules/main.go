package modules

import (
	"log"
	_ "online-shop-home-test/docs"
	"online-shop-home-test/modules/health"
	"online-shop-home-test/modules/role"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router = gin.Default()

func Run() {
	port := os.Getenv("APP_PORT")

	//routes for swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	getRoutes()
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func getRoutes() {
	router.RedirectTrailingSlash = false
	rootRoute := router.Group("/")

	//router for app
	health.Init(rootRoute)
	role.Init(rootRoute)
}

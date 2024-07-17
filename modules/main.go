package modules

import (
	"log"
	_ "online-shop-home-test/docs"
	cartItems "online-shop-home-test/modules/cart_items"
	"online-shop-home-test/modules/health"
	"online-shop-home-test/modules/product"
	productCategory "online-shop-home-test/modules/product_category"
	"online-shop-home-test/modules/role"
	"online-shop-home-test/modules/users"
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
	users.Init(rootRoute)
	productCategory.Init(rootRoute)
	product.Init(rootRoute)
	cartItems.Init(rootRoute)
}

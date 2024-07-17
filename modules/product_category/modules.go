package product_category

import (
	models "online-shop-home-test/modules/product_category/models"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	//list you model here
	utils.AutoMigrate("productCategories", &models.ProductCategories{})
}

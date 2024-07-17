package cart_items

import (
	models "online-shop-home-test/modules/cart_items/models"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	//list you model here
	utils.AutoMigrate("cartItems", &models.CartItems{})
}

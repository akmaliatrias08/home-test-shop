package transaction

import (
	models "online-shop-home-test/modules/transaction/models"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	//list you model here
	utils.AutoMigrate("transaction", &models.Transaction{})
	utils.AutoMigrate("transaction_cart_items", &models.TransactionCartItems{})
}

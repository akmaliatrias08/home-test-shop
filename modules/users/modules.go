package users

import (
	"online-shop-home-test/modules/users/models"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	//list you model here
	utils.AutoMigrate("user", &models.User{})
}

package role

import (
	"online-shop-home-test/modules/role/models"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	//list you model here
	utils.AutoMigrate("role", &models.Role{})
}

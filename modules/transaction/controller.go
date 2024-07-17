package transaction

import (
	"net/http"
	transaction "online-shop-home-test/modules/transaction/dto"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	transactionRoute := rg.Group("/transaction")
	transactionProtectedRoute := transactionRoute.Use(utils.JwtAuthMiddleware())

	transactionProtectedRoute.POST("", checkedOut)

}

// Transaction     				godoc
// @Summary     				Create transaction
// @Description  				Create transaction
// @Param						transaction body transaction.CreateTransactionDTO true  "Create transaction"
// @Param 						Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				transaction
// @Produce      				json
// @Success      				201  {object} transaction.SuccessCreateTransactionResponse
// @Router       				/transaction [post]
func checkedOut(ctx *gin.Context) {
	var input transaction.CreateTransactionDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDAny, _ := ctx.Get("user_id")
	userID, _ := userIDAny.(string)
	data, err := CheckedOut(userID, input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "success",
	})
}

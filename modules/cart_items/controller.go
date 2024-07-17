package cart_items

import (
	"net/http"
	cartItem "online-shop-home-test/modules/cart_items/dto"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	cartItemRoute := rg.Group("/cart-item")
	cartItemProtectedRoute := cartItemRoute.Use(utils.JwtAuthMiddleware())

	cartItemProtectedRoute.POST("", createCartItem)
	cartItemProtectedRoute.GET("", getAllCustomerCartItem)
	cartItemProtectedRoute.DELETE("/:id", deleteCustomerCartItem)
}

// Cart Item     				godoc
// @Summary     				Create cart item product
// @Description  				Create new cart item
// @Param						cart-item body cart_items.CreateCartItemsDTO true  "Create cart item"
// @Param 						Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				cart-item
// @Produce      				json
// @Success      				200  {object} cart_items.SuccessCreateCartItemResponse
// @Failure      				400  {object} cart_items.BadRequestCreateCartItemResponse
// @Router       				/cart-item [post]
func createCartItem(ctx *gin.Context) {
	var input cartItem.CreateCartItemsDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDAny, _ := ctx.Get("user_id")
	userID, _ := userIDAny.(string)
	data, err := CreateCartItems(userID, input)
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

// Cart Item      				godoc
// @Summary     				Get all customer cart item
// @Description  				Get all product from customer cart item
// @Param        				page    	query     string  false  "Number of page to load"
// @Param        				pageSize    query     string  false  "Number page size or limit to display data"
// @Param 						Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				cart-item
// @Produce      				json
// @Success      				200  {object} cart_items.SuccessGetCustomerCartItemResponse
// @Failure      				400  {object} cart_items.BadRequestGetCustomerCartItemResponse
// @Router       				/cart-item [get]
func getAllCustomerCartItem(ctx *gin.Context) {
	userIDAny, _ := ctx.Get("user_id")
	userID, _ := userIDAny.(string)
	count, data, err := GetCartItemByCustomer(userID, ctx.Query("page"), ctx.Query("pageSize"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"count":   count,
		"message": "success",
	})
}

// Cart Item      	godoc
// @Summary     	Delete customer cart item by ID
// @Description  	Delete customer cart item by ID
// @Param			id path string true "cart item ID"
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         	cart-item
// @Produce      	json
// @Success      	200  {object} cart_items.SuccessDeleteCustomerCartItemResponse
// @Failed      	400  {object} cart_items.BadRequestDeleteCustomerCartItemResponse
// @Router       	/cart-item/{id} [delete]
func deleteCustomerCartItem(ctx *gin.Context) {
	userIDAny, _ := ctx.Get("user_id")
	userID, _ := userIDAny.(string)
	err := DeleteCartItem(userID, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

package product

import (
	"net/http"
	product "online-shop-home-test/modules/product/dto"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	productRoute := rg.Group("/product")
	productProtectedRoute := productRoute.Use(utils.JwtAuthMiddleware())

	productProtectedRoute.POST("", createProduct)
	productProtectedRoute.GET("/:product_category_id", getAllProductByCategory)

}

// Product     					godoc
// @Summary     				Create product
// @Description  				Create new product
// @Param						product body product.CreateProductDTO true  "Create product"
// @Param 						Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				product
// @Produce      				json
// @Success      				201  {object} product.SuccessCreateProductResponse
// @Failure      				400  {object} product.BadRequestCreateProductResponse
// @Router       				/product [post]
func createProduct(ctx *gin.Context) {
	var input product.CreateProductDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := CreateProduct(input)
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

// Product      				godoc
// @Summary     				Get all product by product category
// @Description  				Get all product by product category
// @Param						product_category_id path string true "product category id from product category"
// @Param        				page    	query     string  false  "Number of page to load"
// @Param        				pageSize    query     string  false  "Number page size or limit to display data"
// @Param 						Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				product
// @Produce      				json
// @Success      				200  {object} product.SuccessGetProductByCategoryResponse
// @Failure      				400  {object} product.BadRequestGetProductByCategoryResponse
// @Router       				/product/{product_category_id} [get]
func getAllProductByCategory(ctx *gin.Context) {
	count, data, err := GetProductByProductCategoryId(ctx.Param("product_category_id"), ctx.Query("page"), ctx.Query("pageSize"))
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

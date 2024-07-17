package product_category

import (
	"net/http"
	productCategory "online-shop-home-test/modules/product_category/dto"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	productCategoryRoute := rg.Group("/product-category")
	productCategoryProtectedRoute := productCategoryRoute.Use(utils.JwtAuthMiddleware())

	productCategoryProtectedRoute.POST("", createProductCategory)
	productCategoryProtectedRoute.GET("", getAllProductCategory)
}

// Product Category      		godoc
// @Summary     				Create product category
// @Description  				Create new product category
// @Param						product-category body productCategory.CreateProductCategoryDTO true  "Create product category"
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				product-category
// @Produce      				json
// @Success      				201  {object} product_category.SuccessCreateProductCategoryResponse
// @Failure      				400  {object} product_category.BadRequestCreateProductCategoryResponse
// @Router       				/product-category [post]
func createProductCategory(ctx *gin.Context) {
	var input productCategory.CreateProductCategoryDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := CreateProductCategory(input)
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

// Product Category      		godoc
// @Summary     				Get all product category
// @Description  				Get all product category
// @Param        				page    	query     string  false  "Number of page to load"
// @Param        				pageSize    query     string  false  "Number page size or limit to display data"
// @Param 						Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         				product-category
// @Produce      				json
// @Success      				200  {object} product_category.SuccessGetAllProductCategoryResponse
// @Router       				/product-category [get]
func getAllProductCategory(ctx *gin.Context) {
	count, data, err := GetAllProductCategory(ctx.Query("page"), ctx.Query("pageSize"))
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

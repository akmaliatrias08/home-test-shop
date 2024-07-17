package product_category

import models "online-shop-home-test/modules/product_category/models"

type SuccessGetAllProductCategoryResponse struct {
	Count   int64                      `json:"count" example:"1"`
	Data    []models.ProductCategories `json:"data"`
	Message string                     `json:"message" example:"success"`
}

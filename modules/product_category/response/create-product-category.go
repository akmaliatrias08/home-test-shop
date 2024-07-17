package product_category

import models "online-shop-home-test/modules/product_category/models"

type SuccessCreateProductCategoryResponse struct {
	Data    models.ProductCategories `json:"data"`
	Message string                   `json:"message" example:"success"`
}

type BadRequestCreateProductCategoryResponse struct {
	Error string `json:"message" example:"error message"`
}

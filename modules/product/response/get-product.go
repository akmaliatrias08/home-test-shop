package product

import models "online-shop-home-test/modules/product/models"

type SuccessGetProductByCategoryResponse struct {
	Count   int64            `json:"count" example:"1"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message" example:"success"`
}

type BadRequestGetProductByCategoryResponse struct {
	Error string `json:"error" example:"record not found"`
}

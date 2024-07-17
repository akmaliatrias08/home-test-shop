package product

import models "online-shop-home-test/modules/product/models"

type SuccessCreateProductResponse struct {
	Data    models.Product `json:"data"`
	Message string         `json:"message" example:"success"`
}

type BadRequestCreateProductResponse struct {
	Error string `json:"error" example:"error message"`
}

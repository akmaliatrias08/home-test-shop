package cart_items

import models "online-shop-home-test/modules/cart_items/models"

type SuccessGetCustomerCartItemResponse struct {
	Count   int64              `json:"count" example:"1"`
	Data    []models.CartItems `json:"data"`
	Message string             `json:"message" example:"success"`
}

type BadRequestGetCustomerCartItemResponse struct {
	Error string `json:"error" example:"record not found"`
}

package cart_items

import models "online-shop-home-test/modules/cart_items/models"

type SuccessCreateCartItemResponse struct {
	Data    models.CartItems `json:"data"`
	Message string           `json:"message" example:"success"`
}

type BadRequestCreateCartItemResponse struct {
	Error string `json:"error" example:"record not found"`
}

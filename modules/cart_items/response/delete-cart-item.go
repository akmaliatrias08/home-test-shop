package cart_items

type SuccessDeleteCustomerCartItemResponse struct {
	Message string `json:"message" example:"success"`
}

type BadRequestDeleteCustomerCartItemResponse struct {
	Error string `json:"error" example:"cannot delete cart items because user are not created"`
}

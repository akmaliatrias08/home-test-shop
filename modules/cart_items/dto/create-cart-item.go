package cart_items

type CreateCartItemsDTO struct {
	ProductID string `json:"product_id" binding:"required" example:"0723074e-5c4b-4c4c-9073-e4f028cbd2f8"`
	Quantity  int    `json:"quantity" binding:"required" example:"1"`
}

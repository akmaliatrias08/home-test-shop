package transaction

import "github.com/google/uuid"

type CreateTransactionDTO struct {
	CartItems []CreateTransactionCartItemsDTO `json:"cart_items" binding:"required"`
}

type CreateTransactionCartItemsDTO struct {
	ID        uuid.UUID `json:"id" binding:"required" example:"769dee9f-1745-4419-87d3-bd3608ac36db"`
	ProductID string    `json:"product_id" binding:"required" example:"0723074e-5c4b-4c4c-9073-e4f028cbd2f8"`
	Price     float64   `json:"price" example:"100000" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required" example:"1"`
}

package transaction

import (
	cartItems "online-shop-home-test/modules/cart_items/models"
	"online-shop-home-test/modules/utils"

	"github.com/google/uuid"
)

type TransactionCartItems struct {
	utils.Base    `gorm:"embedded"`
	TransactionID uuid.UUID `json:"transaction_id" gorm:"foreignKey:id;references:TransactionID" example:"7c08adb6-a702-45c5-99b7-1b33868c0e01"`
	CartItemID    uuid.UUID `json:"cart_item_id" gorm:"foreignKey:id;references:CartItemID" example:"7c08adb6-a702-45c5-99b7-1b33868c0e02"`
	Price         float64   `json:"price" example:"100000"`
	Quantity      int       `json:"quantity" example:"1"`

	CartItem *cartItems.CartItems `json:"cart_item,omitempty"`
}

package transaction

import (
	users "online-shop-home-test/modules/users/models"
	"online-shop-home-test/modules/utils"

	"github.com/google/uuid"
)

type Transaction struct {
	utils.Base    `gorm:"embedded"`
	CustomerID    uuid.UUID `json:"customer_id" gorm:"foreignKey:id;references:CustomerID" example:"7c08adb6-a702-45c5-99b7-1b33868c0e01"`
	TotalPrice    float64   `json:"total_price"`
	PaymentStatus string    `json:"payment_status"`

	Customer             *users.User            `json:"customer,omitempty"`
	TransactionCartItems []TransactionCartItems `json:"transaction_cart_items,omitempty"`
}

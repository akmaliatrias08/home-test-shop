package cart_items

import (
	product "online-shop-home-test/modules/product/models"
	users "online-shop-home-test/modules/users/models"

	"online-shop-home-test/utils"

	"github.com/google/uuid"
)

type CartItems struct {
	utils.Base   `gorm:"embedded"`
	ProductID    uuid.UUID `json:"product_id" gorm:"foreignKey:id;references:ProductID" example:"0723074e-5c4b-4c4c-9073-e4f028cbd2f8"`
	CustomerID   uuid.UUID `json:"customer_id" gorm:"foreignKey:id;references:CustomerID" example:"7c08adb6-a702-45c5-99b7-1b33868c0e01"`
	Quantity     int       `json:"quantity"`
	IsCheckedOut bool      `json:"is_checked_out"`

	Product  *product.Product `json:"product,omitempty"`
	Customer *users.User      `json:"customer,omitempty"`
}

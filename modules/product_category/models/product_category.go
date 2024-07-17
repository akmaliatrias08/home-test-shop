package product_category

import "online-shop-home-test/modules/utils"

type ProductCategories struct {
	utils.Base `gorm:"embedded"`
	Name       string `json:"name" example:"electronic"`
}

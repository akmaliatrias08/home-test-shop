package product

import (
	models "online-shop-home-test/modules/product_category/models"
	"online-shop-home-test/utils"

	"github.com/google/uuid"
)

type Product struct {
	utils.Base        `gorm:"embedded"`
	Name              string    `json:"name" example:"Couple Ring"`
	Price             float64   `json:"price" example:"15000"`
	ProductCategoryID uuid.UUID `json:"product_category_id" gorm:"foreignKey:id;references:ProductCategoryID" example:"ab7ac1cb-17c6-4e9a-8cd8-d51d8988c5ec"`

	ProductCategory *models.ProductCategories `json:"product_category,omitempty"`
}

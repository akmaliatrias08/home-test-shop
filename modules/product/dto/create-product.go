package product

type CreateProductDTO struct {
	Name              string  `json:"name" example:"Love Gold Earing" binding:"required"`
	Price             float64 `json:"price" example:"200000" binding:"required"`
	ProductCategoryID string  `json:"product_category_id" example:"34407d74-2ab1-4933-bbe7-1240b0be9702" binding:"required"`
}

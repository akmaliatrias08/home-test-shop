package product

import (
	"online-shop-home-test/modules/helpers"
	product "online-shop-home-test/modules/product/dto"
	models "online-shop-home-test/modules/product/models"
	productCategoryModule "online-shop-home-test/modules/product_category"
	"online-shop-home-test/modules/utils"
)

func CreateProduct(createProductDTO product.CreateProductDTO) (models.Product, error) {
	var product models.Product

	//validate product category id
	productCategory, err := productCategoryModule.GetProductCategoryById(createProductDTO.ProductCategoryID)
	if err != nil {
		return product, err
	}

	//create
	product.Name = createProductDTO.Name
	product.Price = createProductDTO.Price
	product.ProductCategory = &productCategory
	product.ProductCategoryID = productCategory.ID

	if err := utils.DB.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func GetProductByProductCategoryId(productCategoryId, page, pageSize string) (int64, []models.Product, error) {
	var productsBycategory []models.Product
	var count int64

	//handle pagination
	limit, offset, err := helpers.Pagination(page, pageSize)
	if err != nil {
		return count, productsBycategory, err
	}

	//validate product category id
	_, err = productCategoryModule.GetProductCategoryById(productCategoryId)
	if err != nil {
		return count, productsBycategory, err
	}

	//count all product by categories
	if err := utils.DB.Model(&productsBycategory).Where("product_category_id", productCategoryId).Count(&count).Error; err != nil {
		return count, productsBycategory, err
	}

	//get all product categories
	if err := utils.DB.Model(&productsBycategory).Where("product_category_id", productCategoryId).Preload("ProductCategory").Limit(limit).Offset(offset).Find(&productsBycategory).Error; err != nil {
		return count, productsBycategory, err
	}

	return count, productsBycategory, nil
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product

	if err := utils.DB.Model(&product).Where("id", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

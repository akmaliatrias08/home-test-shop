package product_category

import (
	"online-shop-home-test/modules/helpers"
	productCategory "online-shop-home-test/modules/product_category/dto"
	models "online-shop-home-test/modules/product_category/models"
	"online-shop-home-test/modules/utils"
)

func GetAllProductCategory(page, pageSize string) (int64, []models.ProductCategories, error) {
	var count int64
	var productCategories []models.ProductCategories

	//handle pagination
	limit, offset, err := helpers.Pagination(page, pageSize)
	if err != nil {
		return count, productCategories, err
	}

	//count all product categories
	if err := utils.DB.Model(&productCategories).Count(&count).Error; err != nil {
		return count, productCategories, err
	}

	//get all product categories
	if err := utils.DB.Model(&productCategories).Limit(limit).Offset(offset).Find(&productCategories).Error; err != nil {
		return count, productCategories, err
	}

	return count, productCategories, nil
}

func GetProductCategoryById(id string) (models.ProductCategories, error) {
	var productCategory models.ProductCategories
	if err := utils.DB.Where("id", id).First(&productCategory).Error; err != nil {
		return productCategory, err
	}

	return productCategory, nil
}

func CreateProductCategory(createProductCategoryDTO productCategory.CreateProductCategoryDTO) (models.ProductCategories, error) {
	var productCategory models.ProductCategories

	productCategory.Name = createProductCategoryDTO.Name
	if err := utils.DB.Create(&productCategory).Error; err != nil {
		return productCategory, err
	}

	return productCategory, nil
}

package cart_items

import (
	"errors"
	"fmt"
	cartItems "online-shop-home-test/modules/cart_items/dto"
	models "online-shop-home-test/modules/cart_items/models"
	"online-shop-home-test/modules/helpers"
	product "online-shop-home-test/modules/product"
	"online-shop-home-test/modules/users"
	"online-shop-home-test/modules/utils"

	"gorm.io/gorm"
)

func CreateCartItems(userID string, createCartItemsDTO cartItems.CreateCartItemsDTO) (models.CartItems, error) {
	var cartItem models.CartItems

	//validate the user id
	user, err := users.GetUserById(userID)
	if err != nil {
		return cartItem, err
	}

	//validate product id
	product, err := product.GetProductById(createCartItemsDTO.ProductID)
	if err != nil {
		return cartItem, err
	}

	//validate cart item.product_id with customer_id is not exist
	productCartItem, err := GetCustomerCartItemByProductID(userID, createCartItemsDTO.ProductID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return cartItem, err
	}

	//validate same product id in one customer cart cannot be created
	if product.ID == productCartItem.ProductID {
		return cartItem, errors.New("same product cannot be added as new cart item")
	}

	//create
	cartItem.Customer = &user
	cartItem.CustomerID = user.ID
	cartItem.Product = &product
	cartItem.ProductID = product.ID
	cartItem.Quantity = createCartItemsDTO.Quantity

	if err := utils.DB.Create(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func GetCartItemByCustomer(userID, page, pageSize string) (int64, []models.CartItems, error) {
	var customerCartItems []models.CartItems
	var count int64

	//validate the user id
	_, err := users.GetUserById(userID)
	if err != nil {
		return count, customerCartItems, err
	}

	//handle pagination
	limit, offset, err := helpers.Pagination(page, pageSize)
	if err != nil {
		return count, customerCartItems, err
	}

	//get where cart item is_checked_out == false && customer_id == userID
	whereDefault := fmt.Sprintf("is_checked_out = false AND customer_id = '%s'", userID)

	//count cart item with where default
	if err := utils.DB.Model(&customerCartItems).Where(whereDefault).Count(&count).Error; err != nil {
		return count, customerCartItems, err
	}

	//get cart item with where default
	if err := utils.DB.Model(&customerCartItems).Where(whereDefault).Preload("Product").Limit(limit).Offset(offset).Find(&customerCartItems).Error; err != nil {
		return count, customerCartItems, err
	}

	return count, customerCartItems, nil
}

func GetCartItemById(id string) (models.CartItems, error) {
	var cartItem models.CartItems

	if err := utils.DB.Model(&cartItem).Where("id", id).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func GetCustomerCartItemByProductID(userID, productID string) (models.CartItems, error) {
	var cartItem models.CartItems

	if err := utils.DB.Model(&cartItem).Where("product_id", productID).Where("customer_id", userID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func DeleteCartItem(userID, id string) error {
	var cartItem models.CartItems

	//validation cart item
	_, err := GetCartItemById(id)
	if err != nil {
		return err
	}

	//validation user
	userDeleteCartItem, err := users.GetUserById(userID)
	if err != nil {
		return err
	}

	//validation the cartItems.customer_id != userId (to validate the user create cart items are the same with user delete cart items)
	if cartItem.CustomerID != userDeleteCartItem.ID {
		return errors.New("cannot delete cart items because user are not created")
	}

	if err := utils.DB.Model(&cartItem).Where("id", id).Delete(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

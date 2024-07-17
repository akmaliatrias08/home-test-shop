package transaction

import (
	cartItems "online-shop-home-test/modules/cart_items"
	transaction "online-shop-home-test/modules/transaction/dto"
	models "online-shop-home-test/modules/transaction/models"
	"online-shop-home-test/modules/users"
	"online-shop-home-test/modules/utils"
)

func CheckedOut(userID string, createTransactionDTO transaction.CreateTransactionDTO) (models.Transaction, error) {
	var transaction models.Transaction
	var transactionCartItems []models.TransactionCartItems
	var cartItemIDs []string
	var totalPrice float64

	//validate the user id
	customer, err := users.GetUserById(userID)
	if err != nil {
		return transaction, err
	}

	//calculate the total price (quantity * total price)
	for _, cartItem := range createTransactionDTO.CartItems {
		totalPrice += (cartItem.Price * float64(cartItem.Quantity))

		trxCartItem := models.TransactionCartItems{
			CartItemID: cartItem.ID,
			Price:      cartItem.Price,
			Quantity:   cartItem.Quantity,
		}

		cartItemIDs = append(cartItemIDs, cartItem.ID.String())
		transactionCartItems = append(transactionCartItems, trxCartItem)
	}

	//create transaction
	transaction.CustomerID = customer.ID
	transaction.Customer = &customer
	transaction.TotalPrice = totalPrice
	transaction.TransactionCartItems = transactionCartItems
	transaction.PaymentStatus = "PENDING"

	if err := utils.DB.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	//update is_checked_out from cart items
	err = cartItems.UpdateCustomerCartItemCheckedOut(userID, cartItemIDs)
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

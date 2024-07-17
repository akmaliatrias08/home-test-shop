package transaction

import models "online-shop-home-test/modules/transaction/models"

type SuccessCreateTransactionResponse struct {
	Data    models.Transaction `json:"data"`
	Message string             `json:"message" example:"success"`
}

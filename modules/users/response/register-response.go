package users

import (
	"online-shop-home-test/modules/users/models"
)

type SuccessRegisterResponse struct {
	Data    models.User `json:"data"`
	Message string      `json:"message" example:"success"`
}

type BadRequestRegisterResponse struct {
	Error string `json:"error" example:"username is exist"`
}

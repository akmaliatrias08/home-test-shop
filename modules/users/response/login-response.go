package users

import users "online-shop-home-test/modules/users/dto"

type SuccessLoginResponse struct {
	Data    users.TokenResponse `json:"data"`
	Message string              `json:"message" example:"success"`
}

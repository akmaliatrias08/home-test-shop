package users

type SuccessAuthorizeToken struct {
	Message string `json:"message" example:"token is valid"`
}

type UnauthorizedToken struct {
	Message string `json:"message" example:"unauthorized"`
}

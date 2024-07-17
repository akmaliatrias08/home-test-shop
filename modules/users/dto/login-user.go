package users

type LoginDTO struct {
	Username string `json:"username" example:"adminshop" binding:"required"`
	Password string `json:"password" example:"adminshop123" binding:"required"`
}

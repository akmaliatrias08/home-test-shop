package users

type CreateUserDTO struct {
	Username string `json:"username" example:"adminshop" binding:"required"`
	Password string `json:"password" example:"adminshop123" binding:"required"`
	Name     string `json:"name" example:"Admin Shop" binding:"required"`
	RoleID   string `json:"role_id" example:"8cb17fe3-3552-4bfa-852f-a9a7d91fd00d" binding:"required"`
}

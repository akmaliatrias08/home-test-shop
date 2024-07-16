package role

type CreateRoleDTO struct {
	Name string `json:"name" example:"Admin" binding:"required"`
}

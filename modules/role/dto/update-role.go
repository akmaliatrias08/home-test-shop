package role

type UpdateRoleDTO struct {
	Name string `json:"name" example:"Shop" binding:"required"`
}

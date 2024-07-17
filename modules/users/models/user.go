package models

import (
	"online-shop-home-test/modules/role/models"
	"online-shop-home-test/modules/utils"

	"github.com/google/uuid"
)

type User struct {
	utils.Base `gorm:"embedded"`
	Username   string    `json:"username" gorm:"type:varchar(100);unique;not null"`
	Salt       string    `json:"salt"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	RoleID     uuid.UUID `json:"role_id" gorm:"foreignKey:id;references:RoleID"`

	Role *models.Role `json:"role,omitempty"`
}

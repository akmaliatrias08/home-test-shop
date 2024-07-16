package models

import "online-shop-home-test/modules/utils"

type Role struct {
	utils.Base `gorm:"embedded"`
	Name       string `json:"name" example:"Admin" gorm:"type:varchar(100)"`
}

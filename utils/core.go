package utils

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `json:"id" example:"ab7ac1cb-17c6-4e9a-8cd8-d51d8988c5ec" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now();autoCreateTime" example:"2024-07-16T23:13:03.115483+07:00"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now();autoUpdateTime" example:"2024-07-16T23:13:03.115483+07:00"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at" gorm:"nullable:true" example:"null"`
}

func (base *Base) BeforeCreate(_ *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}

	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	base.DeletedAt = nil
	return
}

func (base *Base) BeforeUpdate(_ *gorm.DB) (err error) {
	base.UpdatedAt = time.Now()
	return
}

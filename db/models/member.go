package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	// gorm.Model
	ID        string `gorm:"type:uuid; primaryKey; column:id" json:"id"`
	FirstName string `gorm:"type:varchar(200); column:firstName" json:"firstName"`
	LastName  string `gorm:"type:varchar(200); column:lastName" json:"lastName"`

	// CreatedAt time.Time      `gorm:"column:createdAt" json:"createdAt"`
	// UpdatedAt time.Time      `gorm:"column:updatedAt" json:"updatedAt"`
	// DeletedAt gorm.DeletedAt `gorm:"index; column:deletedAt" json:"deletedAt"`

	CreatedAt time.Time      `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index; column:deletedAt" json:"deletedAt"`
}

func (u Member) TableName() string {
	return "member"
}

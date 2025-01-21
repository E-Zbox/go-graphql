package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Age       int       `gorm:"not null;size:3" json:"age"`
	Country   string    `gorm:"not null" json:"country"`
	Email     string    `gorm:"not null;unique" json:"email"`
	FirstName string    `gorm:"uniqueIndex" json:"firstName"`
	LastName  string    `gorm:"uniqueIndex" json:"lastName"`
	Role      string    `gorm:"not null" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt
}

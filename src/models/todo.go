package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Completed bool           `gorm:"not null" json:"completed"`
	Started   bool           `gorm:"not null" json:"started"`
	Text      string         `gorm:"not null" json:"text"`
	UserId    uint           `gorm:"not null" json:"userId"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (todo *Todo) User() *User {
	var user User

	if err := DBInstance.Where("id=?", todo.UserId).First(&user).Error; err != nil {
		DBInstance.AddError(fmt.Errorf("could not retrieve User `%d` for Post", todo.UserId))
	}

	return &user
}

func (todo *Todo) Validate(db *gorm.DB) {
	// check if UserId exists
	var user User
	if err := DBInstance.Where("id=?", todo.UserId).First(&user).Error; err != nil {
		DBInstance.AddError(fmt.Errorf("User with id: `%d` does not exist", todo.UserId))
	}
}

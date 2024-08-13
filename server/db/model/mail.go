package model

import (
	"time"

	"gorm.io/gorm"
)

type Mail struct {
	ID        uint `gorm:"primaryKey,autoIncrement,unique"`
	UserID    uint
	Type      string `gorm:"type:varchar(20)"` // OTP, NOTIFICATION
	Body      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

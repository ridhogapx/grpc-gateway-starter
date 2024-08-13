package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey,autoIncrement,unique"`
	Name        string `gorm:"type:varchar(100)"`
	Email       string `gorm:"type:varchar(30)"`
	Username    string `gorm:"type:varchar(50)"`
	Password    string `gorm:"type:varchar(255)"`
	Status      string `gorm:"type:varchar(15)"`
	ProfilePict string
	AllowNSFW   bool
	Bio         string
	SocialMedia string `gorm:"default:{}"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

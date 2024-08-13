package model

import (
	"time"

	"gorm.io/gorm"
)

type Ilustration struct {
	ID            uint
	Title         string `gorm:"type:varchar(250)"`
	Image         string
	Description   string
	IsNSFW        bool
	IsAIGenerated bool
	UserID        uint
	Username      string `gorm:"type:varchar(50)"`
	ProfilePict   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type FavouriteIlustration struct {
	ID            uint
	UserID        uint
	IlustrationID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

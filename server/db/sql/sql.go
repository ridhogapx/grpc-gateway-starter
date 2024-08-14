package sql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewRepository(db *gorm.DB) *Repository {

	logger := zap.Must(zap.NewProduction())

	return &Repository{
		db:  db,
		log: logger,
	}
}

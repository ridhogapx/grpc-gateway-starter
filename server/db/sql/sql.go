package sql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type repository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewRepository(db *gorm.DB) *repository {

	logger := zap.Must(zap.NewProduction())

	return &repository{
		db:  db,
		log: logger,
	}
}

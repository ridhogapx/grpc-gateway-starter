package config

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(dsn string) (*gorm.DB, error) {

	log := zap.Must(zap.NewProduction())

	defer log.Sync()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Debug("Couldn't initiate dsn connection", zap.Error(err))
		return nil, err
	}

	log.Info("GORM has been successfully initiated...")

	return db, nil

}

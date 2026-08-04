package database

import (
	"gorm/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Product{},
	)
}

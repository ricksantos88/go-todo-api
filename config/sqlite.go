package config

import (
	"github.com/ricksantos88/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Create db and connect
	db, err := gorm.Open(sqlite.Open("./db/openings.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// return DB
	return db, nil
}

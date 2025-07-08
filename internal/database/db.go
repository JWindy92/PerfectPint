package database

import (
	"github.com/JWindy92/PerfectPint/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("perfectpint.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&models.User{}, &models.Review{}, &models.Location{})
	db.AutoMigrate(&models.User{})

	DB = db
}

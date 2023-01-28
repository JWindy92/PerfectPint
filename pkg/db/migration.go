package db

import (
	"github.com/JWindy92/PerfectPint/pkg/models"
	_ "github.com/lib/pq"
)

func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)

	connection.AutoMigrate(models.User{})
}

package db

//* -------------------------------------------------------------------------- */
//*                                  DATABASE                                  */
//* -------------------------------------------------------------------------- */
//* docker run --name postgresdb -p 1111:5432 -e POSTGRES_PASSWORD=dbpass -d postgres

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

const (
	host     = "localhost"
	port     = 1111
	user     = "postgres"
	password = "dbpass"
	dbname   = "postgres"
)

func GetDatabase() *gorm.DB {
	database := dbname
	databaseurl := "postgres://postgres:" + password + "@localhost:1111/" + dbname + "?sslmode=disable"
	connection, err := gorm.Open(database, databaseurl)
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}

func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(User{})
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}

package db

//* -------------------------------------------------------------------------- */
//*                                  DATABASE                                  */
//* -------------------------------------------------------------------------- */
//* docker run --name postgresdb -p 1111:5432 -e POSTGRES_PASSWORD=dbpass -d postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func buildConnString() string {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	var connString = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	fmt.Println(connString)
	return connString
}

func GetDatabase() *gorm.DB {
	databaseurl := buildConnString()
	connection, err := gorm.Open(os.Getenv("PG_DBNAME"), databaseurl)
	if err != nil {
		log.Fatalln("wrong database url", err)
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}

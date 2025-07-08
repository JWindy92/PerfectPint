package database

import "testing"

func TestPostgresConnect(t *testing.T) {
	db := PostgresImpl{}
	db.ConnectDB()
}

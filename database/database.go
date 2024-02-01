package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		"localhost",
		5432,
		"movierental-db",
		"postgres",
		"postgres",
	)

	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}
	return dbConn
}

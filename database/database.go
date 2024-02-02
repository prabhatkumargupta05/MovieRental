package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"movierental/configs"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func CreateConnection(dbConfig configs.DatabaseConfig) *sql.DB {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Databasename,
		dbConfig.User,
		dbConfig.Password,
	)

	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}

	// Run migration scripts
	if err := runMigrations(dbConn); err != nil {
		log.Fatal("unable to run migrations ", err.Error())
	}

	return dbConn
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration", // Replace with the actual path to your migration scripts
		"movierental-db", driver)
	if err != nil {
		return err
	}
	fmt.Println("YYYYYYYY")

	// You can use m.Up() to apply all migrations, or m.Steps(n) to apply 'n' steps.
	// Example: m.Steps(2) applies the first 2 migrations.
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

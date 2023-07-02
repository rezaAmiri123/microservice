package postgres

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func DBMigrate(db *sql.DB, migratePath string, dbName string) error {
	//path = "file://../migrations"
	//dbName = "postgres"
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(migratePath, dbName, driver)
	if err != nil {
		return err
	}
	migrate.
		err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	fmt.Println(err)
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

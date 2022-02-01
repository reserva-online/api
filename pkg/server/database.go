package server

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/schedule-api/pkg/config"
)

func NewDatabase() (*sqlx.DB, error) {
	config := config.NewConfig()

	db, err := sqlx.Connect(config.DatabaseDriver, config.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = migrateDb(config)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrateDb(config config.Config) error {
	m, err := migrate.New(
		config.MigrationsPath,
		config.DatabaseConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	return err
}

func NewTestDatabase(port string) (*sqlx.DB, error) {
	os.Setenv("DATABASE_DRIVER", "postgres")
	os.Setenv("DATABASE_CONNECTION_STRING", fmt.Sprintf("postgres://postgres:postgres@localhost:%v/postgres?sslmode=disable", port))
	os.Setenv("MIGRATION_PATH", "file://../../migrations")
	return NewDatabase()
}

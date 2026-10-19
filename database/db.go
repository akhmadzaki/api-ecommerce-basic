package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/akhmadzaki/api-ecommerce-basic/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func RunMigrations(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	var err error
	DBConn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open db for migration: %v", err)
	}

	driver, err := postgres.WithInstance(DBConn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/database/migrations", // Path to your migration folder
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations applied successfully!")
}

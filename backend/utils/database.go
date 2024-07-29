package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// InitDB initializes the database connection and sets up any necessary tables.
func InitDB() {
	// Load database connection parameters from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	// Open the database connection
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	// Check if the database is reachable
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
	}

	log.Println("Database connected successfully")
}

// MigrateDb applies database migrations
func MigrateDb() {
	migrationsPath := "file://migrations" // Path to your migration files
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	m, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	} else {
		log.Println("Database migrations applied successfully")
	}
}

// GetDB returns the database connection instance.
func GetDB() *sql.DB {
	return db
}

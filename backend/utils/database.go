package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

// GetDB returns the database connection instance.
func GetDB() *sql.DB {
	return db
}

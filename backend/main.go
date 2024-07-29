package main

import (
	"log"
	"net/http"

	"go-book-review/routes"
	"go-book-review/utils"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Initialize the database connection
	utils.InitDB()

	// Apply migrations
	utils.MigrateDb()

	// Register routes
	r := routes.RegisterRoutes()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}

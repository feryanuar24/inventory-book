package main

import (
	"fmt"
	"inventory-book/config"
	"inventory-book/routes"
	"net/http"
)

func main() {
	// Connect to database
	config.ConnectDB()

	// Database migration
	config.MigrateDB()

	// Seed books
	config.SeedBooks()

	// Setup routes
	r := routes.SetupRoutes()

	// Run the server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

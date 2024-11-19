// Package main is the entry point for the gene editor backend service.
// It initializes and runs an HTTP server that provides REST API endpoints
// for ~~gene editing operations and analysis~~ (TODO: add this).
//
// The server is built using the Gin web framework and includes:
//   - Configuration management via environment variables
//   - PostgreSQL database connectivity
//   - RESTful API endpoints for variant analysis
//   - LLM integration for gene editing assistance
//
// To run the server:
//   1. Ensure environment variables are set (see .env.example)
//   2. Run `go run main.go`
//   3. Server will start on configured port (default: 8080)
package main

import (
	"log"

	"github.com/jake-yukich/gnome/backend/api"
	"github.com/jake-yukich/gnome/backend/config"
	"github.com/jake-yukich/gnome/backend/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize router
	router := gin.Default()
	
	// Register routes
	api.RegisterRoutes(router)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
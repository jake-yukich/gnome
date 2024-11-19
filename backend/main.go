package main

import (
	"log"

	"github.com/yourusername/gene-editor/backend/api"
	"github.com/yourusername/gene-editor/backend/config"
	"github.com/yourusername/gene-editor/backend/db"
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
package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jake-yukich/gnome/backend/db"
)

func RegisterRoutes(r *gin.Engine) {
	// test endpoints
	r.GET("/api/test-db", testDBConnection)
	r.POST("/api/test-db", createTestEntry)
	r.GET("/api/test-db/entries", getTestEntries)
}

func testDBConnection(c *gin.Context) {
	err := db.DB.Ping()
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Database connected successfully!"})
}

func createTestEntry(c *gin.Context) {
	var request struct {
		Message string `json:"message"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var id int
	err := db.DB.QueryRow(
		"INSERT INTO test_entries (message) VALUES ($1) RETURNING id",
		request.Message,
	).Scan(&id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create entry: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id, "message": "Entry created successfully"})
}

func getTestEntries(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, message, created_at FROM test_entries ORDER BY created_at DESC LIMIT 10")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch entries: " + err.Error()})
		return
	}
	defer rows.Close()

	var entries []gin.H
	for rows.Next() {
		var id int
		var message string
		var createdAt time.Time

		if err := rows.Scan(&id, &message, &createdAt); err != nil {
			c.JSON(500, gin.H{"error": "Failed to scan entry: " + err.Error()})
			return
		}

		entries = append(entries, gin.H{
			"id":         id,
			"message":    message,
			"created_at": createdAt,
		})
	}

	c.JSON(200, gin.H{"entries": entries})
}

package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", healthCheck)

	// API routes
	api := r.Group("/api")
	{
		api.GET("/variants", getVariants)
		api.GET("/variants/:id", getVariantByID)
		api.POST("/offtarget-risk", calculateOffTargetRisk)
		api.POST("/query-llm", queryLLM)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

// Handler stubs - implement these as needed
func getVariants(c *gin.Context) {
	c.JSON(200, gin.H{"message": "get variants endpoint"})
}

func getVariantByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "get variant by id", "id": id})
}

func calculateOffTargetRisk(c *gin.
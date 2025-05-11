package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ozturkeniss/gomicro-app/product-service/database"
	"github.com/ozturkeniss/gomicro-app/product-service/handlers"
)

// HealthCheck handles the health check endpoint
func HealthCheck(c *gin.Context) {
	// Check database connection
	if err := database.TestDatabaseConnection(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	database.InitDB()

	// Create Gin router
	router := gin.Default()

	// Define API routes
	api := router.Group("/api")
	{
		// Health check endpoint
		api.GET("/health", HealthCheck)

		products := api.Group("/products")
		{
			products.POST("/", handlers.CreateProduct)
			products.GET("/:id", handlers.GetProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
			products.GET("/", handlers.ListProducts)
			products.GET("/search", handlers.SearchProducts)
			products.PUT("/:id/stock", handlers.UpdateStock)
		}
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
} 
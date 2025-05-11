package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ozturkeniss/gomicro-app/user-service/database"
	"github.com/ozturkeniss/gomicro-app/user-service/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ozturkeniss/gomicro-app/user-service/docs" // Swagger docs
)

// @title User Service API
// @version 1.0
// @description This is a user service API.
// @host localhost:8080
// @BasePath /api

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

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define API routes
	api := router.Group("/api")
	{
		// Health check endpoint
		api.GET("/health", HealthCheck)

		users := api.Group("/users")
		{
			users.POST("/register", handlers.RegisterUser)
			users.POST("/login", handlers.LoginUser)
			users.GET("/:id", handlers.GetUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
			users.GET("/", handlers.ListUsers)
		}
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
} 
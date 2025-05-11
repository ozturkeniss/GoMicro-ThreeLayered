package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type OrderService struct{}

func (s *OrderService) CreateOrder(ctx context.Context, req *Request, rsp *Response) error {
	rsp.OrderId = fmt.Sprintf("Order created for user: %s", req.UserId)
	return nil
}

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
	// Create a new service
	service := micro.NewService(
		micro.Name("order.service"),
		micro.Version("latest"),
	)

	// Initialize the service
	service.Init()

	// Register handler
	if err := service.Server().Handle(
		service.Server().NewHandler(&OrderService{}),
	); err != nil {
		log.Fatal(err)
	}

	// Define routes
	api := r.Group("/api")
	{
		// Health check endpoint
		api.GET("/health", HealthCheck)

		// ... existing routes ...
	}

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
} 
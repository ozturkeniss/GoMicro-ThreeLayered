package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ozturkeniss/gomicro-app/order-service/database"
	"github.com/ozturkeniss/gomicro-app/order-service/models"
)

// CreateOrder handles the creation of a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate UserID and ProductID
	if order.UserID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}
	if order.ProductID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID is required"})
		return
	}

	// Fetch product details to calculate total price and check stock
	var product models.Product
	result := database.DB.First(&product, "id = ?", order.ProductID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Check if stock is sufficient
	if product.Stock < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Calculate total price
	order.TotalPrice = product.Price * float64(order.Quantity)

	order.ID = uuid.New()
	result = database.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrder retrieves an order by ID with user and product details
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	result := database.DB.Preload("User").Preload("Product").First(&order, "id = ?", orderID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder updates an order by ID
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = orderID
	result := database.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order by ID
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	result := database.DB.Delete(&models.Order{}, "id = ?", orderID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// ListOrders retrieves a list of orders with user and product details
func ListOrders(c *gin.Context) {
	var orders []models.Order
	result := database.DB.Preload("User").Preload("Product").Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrderStatus updates the status of an order
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	orderID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order
	result := database.DB.First(&order, "id = ?", orderID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = statusUpdate.Status
	result = database.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Log the status update
	logMessage := fmt.Sprintf("Order %s status updated to %s", orderID, statusUpdate.Status)
	if err := logToFile(logMessage); err != nil {
		log.Printf("Failed to log status update: %v", err)
	}

	c.JSON(http.StatusOK, order)
}

// logToFile logs a message to the order status log file
func logToFile(message string) error {
	file, err := os.OpenFile("logs/order_status.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(message)
	return nil
} 
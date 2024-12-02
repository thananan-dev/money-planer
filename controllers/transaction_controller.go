package controllers

import (
	"net/http"
	"money-planer/config"
	"money-planer/models"

	"github.com/gin-gonic/gin"
)

// CreateTransaction handles the creation of a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&transaction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// GetTransactions returns all transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	result := config.DB.Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransaction returns a specific transaction
func GetTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	
	result := config.DB.First(&transaction, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction updates a transaction
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	
	if err := config.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	var updateData models.Transaction
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&transaction).Updates(updateData)
	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction deletes a transaction
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Transaction{}, id)
	
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

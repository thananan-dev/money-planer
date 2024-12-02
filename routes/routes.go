package routes

import (
	"money-planer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		transactions := api.Group("/transactions")
		{
			transactions.POST("/", controllers.CreateTransaction)
			transactions.GET("/", controllers.GetTransactions)
			transactions.GET("/:id", controllers.GetTransaction)
			transactions.PUT("/:id", controllers.UpdateTransaction)
			transactions.DELETE("/:id", controllers.DeleteTransaction)
		}
	}
}

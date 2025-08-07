package routes

import (
	"github.com/gin-gonic/gin"
	"portfolio/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/portfolio")
	{
		api.GET("/holdings", handlers.GetHoldings)
		api.GET("/allocation", handlers.GetAllocation)
		api.GET("/performance", handlers.GetPerformance)
		api.GET("/summary", handlers.GetSummary)
	}
}

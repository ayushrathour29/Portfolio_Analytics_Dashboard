package main

import (
	"net/http"
	"portfolio/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://portfolio-analytics-dashboard-rho.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Register routes
	routes.RegisterRoutes(router)

	// Start server
	router.Run(":10000")
}

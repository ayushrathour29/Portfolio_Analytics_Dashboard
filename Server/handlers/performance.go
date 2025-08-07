package handlers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"portfolio/models"
	"strconv"
)

// Historical performance vs benchmarks
func GetPerformance(c *gin.Context) {
	filePath := filepath.Join("data", "Historical_Performance.csv")
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load performance data"})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var points []models.PerformancePoint

	for i, row := range records {
		if i == 0 {
			continue
		}

		portfolio, _ := strconv.ParseFloat(row[1], 64)
		nifty, _ := strconv.ParseFloat(row[2], 64)
		gold, _ := strconv.ParseFloat(row[3], 64)

		point := models.PerformancePoint{
			Date:      row[0],
			Portfolio: portfolio,
			Nifty50:   nifty,
			Gold:      gold,
		}
		points = append(points, point)
	}

	// Use last 3 data points for return % calculation
	n := len(points)
	if n < 3 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insufficient data for returns"})
		return
	}

	latest := points[n-1]
	mid := points[n-2]
	earliest := points[0]

	returns := map[string]models.Returns{
		"portfolio": {
			OneMonth:   ((latest.Portfolio - mid.Portfolio) / mid.Portfolio) * 100,
			ThreeMonth: ((latest.Portfolio - earliest.Portfolio) / earliest.Portfolio) * 100,
			OneYear:    15.7, // placeholder
		},
		"nifty50": {
			OneMonth:   ((latest.Nifty50 - mid.Nifty50) / mid.Nifty50) * 100,
			ThreeMonth: ((latest.Nifty50 - earliest.Nifty50) / earliest.Nifty50) * 100,
			OneYear:    12.4, // placeholder
		},
		"gold": {
			OneMonth:   ((latest.Gold - mid.Gold) / mid.Gold) * 100,
			ThreeMonth: ((latest.Gold - earliest.Gold) / earliest.Gold) * 100,
			OneYear:    8.9, // placeholder
		},
	}

	response := models.PerformanceData{
		Timeline: points,
		Returns:  returns,
	}

	c.JSON(http.StatusOK, response)
}

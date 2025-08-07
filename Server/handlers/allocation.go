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

func GetAllocation(c *gin.Context) {
	filePath := filepath.Join("data", "SamplePortfolio.csv")
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load holdings data"})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	sectorMap := make(map[string]float64)
	marketMap := make(map[string]float64)
	var total float64

	for i, row := range records {
		if i == 0 {
			continue
		}

		qty, _ := strconv.Atoi(row[2])
		currentPrice, _ := strconv.ParseFloat(row[4], 64)
		value := float64(qty) * currentPrice

		sectorMap[row[5]] += value
		marketMap[row[6]] += value
		total += value
	}

	bySector := make(map[string]models.CategoryAllocation)
	byMarket := make(map[string]models.CategoryAllocation)

	for k, v := range sectorMap {
		bySector[k] = models.CategoryAllocation{Value: v, Percentage: (v / total) * 100}
	}
	for k, v := range marketMap {
		byMarket[k] = models.CategoryAllocation{Value: v, Percentage: (v / total) * 100}
	}

	c.JSON(http.StatusOK, models.Allocation{
		BySector:    bySector,
		ByMarketCap: byMarket,
	})
}

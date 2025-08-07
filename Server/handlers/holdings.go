package handlers

import (
	"encoding/csv"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"portfolio/models"

	"github.com/gin-gonic/gin"
)

// Load portfolio holdings from CSV and calculate values
func GetHoldings(c *gin.Context) {
	filePath := filepath.Join("data", "SamplePortfolio.csv")
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load holdings data"})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	var holdings []models.Holding

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		qty, _ := strconv.Atoi(row[2])
		avgPrice, _ := strconv.ParseFloat(row[3], 64)
		currentPrice, _ := strconv.ParseFloat(row[4], 64)

		value := float64(qty) * currentPrice
		gain := (currentPrice - avgPrice) * float64(qty)
		gainPercent := (gain / (avgPrice * float64(qty))) * 100

		holding := models.Holding{
			Symbol:          row[0],
			Name:            row[1],
			Quantity:        qty,
			AvgPrice:        avgPrice,
			CurrentPrice:    currentPrice,
			Sector:          row[5],
			MarketCap:       row[6],
			Value:           value,
			GainLoss:        gain,
			GainLossPercent: gainPercent,
		}

		holdings = append(holdings, holding)
	}

	c.JSON(http.StatusOK, holdings)
}

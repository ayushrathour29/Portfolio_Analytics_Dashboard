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

func GetSummary(c *gin.Context) {
	// Read holdings CSV file
	filePath := filepath.Join("data", "SamplePortfolio.csv")
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load holdings data"})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var totalValue, totalInvested float64
	var top models.Holding
	var worst models.Holding

	// Initialize top and worst gain % with extreme values
	top.GainLossPercent = -99999
	worst.GainLossPercent = 99999

	sectorMap := map[string]bool{}

	// Iterate through each row in the holdings CSV
	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}

		// Parse CSV fields
		qty, _ := strconv.Atoi(row[2])                    // Quantity
		avgPrice, _ := strconv.ParseFloat(row[3], 64)     // Average Buy Price
		currentPrice, _ := strconv.ParseFloat(row[4], 64) // Current Market Price
		sector := row[5]
		name := row[1]

		// Calculate individual investment and current value
		invested := float64(qty) * avgPrice
		value := float64(qty) * currentPrice

		// Calculate gain/loss and its percentage
		gain := value - invested
		var gainPercent float64
		if invested != 0 {
			gainPercent = (gain / invested) * 100
		} else {
			gainPercent = 0
		}

		// Track unique sectors for diversification score
		sectorMap[sector] = true

		// Update portfolio totals
		totalValue += value
		totalInvested += invested

		// Create Holding struct for this row
		h := models.Holding{
			Symbol:          row[0],
			Name:            name,
			Quantity:        qty,
			AvgPrice:        avgPrice,
			CurrentPrice:    currentPrice,
			GainLossPercent: gainPercent,
		}

		// Determine top performer
		if gainPercent > top.GainLossPercent {
			top = h
		}
		// Determine worst performer
		if gainPercent < worst.GainLossPercent {
			worst = h
		}
	}

	// Calculate total portfolio gain/loss and percentage
	totalGain := totalValue - totalInvested
	var totalGainPercent float64
	if totalInvested != 0 {
		totalGainPercent = (totalGain / totalInvested) * 100
	} else {
		totalGainPercent = 0
	}

	// Prepare final summary response
	summary := models.Summary{
		TotalValue:           totalValue,
		TotalInvested:        totalInvested,
		TotalGainLoss:        totalGain,
		TotalGainLossPercent: totalGainPercent,
		TopPerformer:         top,                           // Includes gain %
		WorstPerformer:       worst,                         // Includes loss %
		DiversificationScore: float64(len(sectorMap)) * 2.0, // Simple metric
		RiskLevel:            "Moderate",
	}

	// Return JSON response
	c.JSON(http.StatusOK, summary)
}

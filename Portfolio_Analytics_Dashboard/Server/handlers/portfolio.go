package handlers

import (
	"encoding/csv"
	"fmt"
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

// Sector and MarketCap distribution
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

	sectorMap := map[string]float64{}
	marketMap := map[string]float64{}
	var total float64 = 0

	for i, row := range records {
		if i == 0 {
			continue
		}

		qty, _ := strconv.Atoi(row[2])
		currentPrice, _ := strconv.ParseFloat(row[4], 64)
		sector := row[5]
		marketCap := row[6]

		value := float64(qty) * currentPrice
		sectorMap[sector] += value
		marketMap[marketCap] += value
		total += value
	}

	bySector := map[string]models.CategoryAllocation{}
	byMarket := map[string]models.CategoryAllocation{}

	for k, v := range sectorMap {
		bySector[k] = models.CategoryAllocation{Value: v, Percentage: (v / total) * 100}
	}

	for k, v := range marketMap {
		byMarket[k] = models.CategoryAllocation{Value: v, Percentage: (v / total) * 100}
	}

	response := models.Allocation{
		BySector:    bySector,
		ByMarketCap: byMarket,
	}

	c.JSON(http.StatusOK, response)
}

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
			OneMonth:  ((latest.Portfolio - mid.Portfolio) / mid.Portfolio) * 100,
			ThreeMonth: ((latest.Portfolio - earliest.Portfolio) / earliest.Portfolio) * 100,
			OneYear:   15.7, // placeholder
		},
		"nifty50": {
			OneMonth:  ((latest.Nifty50 - mid.Nifty50) / mid.Nifty50) * 100,
			ThreeMonth: ((latest.Nifty50 - earliest.Nifty50) / earliest.Nifty50) * 100,
			OneYear:   12.4, // placeholder
		},
		"gold": {
			OneMonth:  ((latest.Gold - mid.Gold) / mid.Gold) * 100,
			ThreeMonth: ((latest.Gold - earliest.Gold) / earliest.Gold) * 100,
			OneYear:   8.9, // placeholder
		},
	}

	response := models.PerformanceData{
		Timeline: points,
		Returns:  returns,
	}

	c.JSON(http.StatusOK, response)
}

// Portfolio summary
func GetSummary(c *gin.Context) {
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
	top.GainLossPercent = -99999
	worst.GainLossPercent = 99999

	sectorMap := map[string]bool{}

	for i, row := range records {
		if i == 0 {
			continue
		}

		qty, _ := strconv.Atoi(row[2])
		avgPrice, _ := strconv.ParseFloat(row[3], 64)
		currentPrice, _ := strconv.ParseFloat(row[4], 64)
		sector := row[5]
		name := row[1]

		invested := float64(qty) * avgPrice
		value := float64(qty) * currentPrice
		gain := value - invested
		gainPercent := (gain / invested) * 100

		sectorMap[sector] = true
		totalValue += value
		totalInvested += invested

		h := models.Holding{
			Symbol:          row[0],
			Name:            name,
			Quantity:        qty,
			AvgPrice:        avgPrice,
			CurrentPrice:    currentPrice,
			GainLossPercent: gainPercent,
		}

		if gainPercent > top.GainLossPercent {
			top = h
		}
		fmt.Println(top)
		if gainPercent < worst.GainLossPercent {
			worst = h
		}
		fmt.Println(worst)
	}

	totalGain := totalValue - totalInvested
	totalGainPercent := (totalGain / totalInvested) * 100

	summary := models.Summary{
		TotalValue:         totalValue,
		TotalInvested:      totalInvested,
		TotalGainLoss:      totalGain,
		TotalGainLossPercent: totalGainPercent,
		TopPerformer:       top,
		WorstPerformer:     worst,
		DiversificationScore: float64(len(sectorMap)) * 2.0, // basic formula
		RiskLevel:          "Moderate",
	}

	c.JSON(http.StatusOK, summary)
}

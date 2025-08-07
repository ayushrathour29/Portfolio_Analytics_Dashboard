package models

type Holding struct {
	Symbol          string  `json:"symbol"`
	Name            string  `json:"name"`
	Quantity        int     `json:"quantity"`
	AvgPrice        float64 `json:"avgPrice"`
	CurrentPrice    float64 `json:"currentPrice"`
	Sector          string  `json:"sector"`
	MarketCap       string  `json:"marketCap"`
	Value           float64 `json:"value"`
	GainLoss        float64 `json:"gainLoss"`
	GainLossPercent float64 `json:"gainLossPercent"`
}

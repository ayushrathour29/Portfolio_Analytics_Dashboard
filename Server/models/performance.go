package models

type PerformancePoint struct {
	Date      string  `json:"date"`
	Portfolio float64 `json:"portfolio"`
	Nifty50   float64 `json:"nifty50"`
	Gold      float64 `json:"gold"`
}

type Returns struct {
	OneMonth  float64 `json:"1month"`
	ThreeMonth float64 `json:"3months"`
	OneYear   float64 `json:"1year"`
}

type PerformanceData struct {
	Timeline []PerformancePoint         `json:"timeline"`
	Returns  map[string]Returns         `json:"returns"`
}

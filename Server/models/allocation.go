package models

type Allocation struct {
	BySector    map[string]CategoryAllocation `json:"bySector"`
	ByMarketCap map[string]CategoryAllocation `json:"byMarketCap"`
}

type CategoryAllocation struct {
	Value      float64 `json:"value"`
	Percentage float64 `json:"percentage"`
}

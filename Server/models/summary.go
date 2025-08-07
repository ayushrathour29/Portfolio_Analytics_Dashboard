package models

type Summary struct {
	TotalValue         float64   `json:"totalValue"`
	TotalInvested      float64   `json:"totalInvested"`
	TotalGainLoss      float64   `json:"totalGainLoss"`
	TotalGainLossPercent float64 `json:"totalGainLossPercent"`
	TopPerformer       Holding   `json:"topPerformer"`
	WorstPerformer     Holding   `json:"worstPerformer"`
	DiversificationScore float64 `json:"diversificationScore"`
	RiskLevel          string    `json:"riskLevel"`
}

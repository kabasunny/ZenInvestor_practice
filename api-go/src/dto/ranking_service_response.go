// api-go\src\dto\ranking_service_response.go
package dto

// RankingServiceResponse はランキング表示用のレスポンスのDTO
type RankingServiceResponse struct {
	Ranking     int     `json:"ranking"`
	Ticker      string  `json:"ticker"`
	Date        string  `json:"date"`
	AvgTurnover float64 `json:"avg_volue"`
	Name        string  `json:"name"`
	LatestClose float64 `json:"latest_close"`
}

// api-go\src\dto\ranking_service_respons.go
package dto

// RankingServiceRespons はランキング表示用のレスポンスのDTO
type RankingServiceRespons struct {
	Ranking     int     `json:"ranking"`
	Ticker      string  `json:"ticker"`
	Date        string  `json:"date"`
	AvgVolue    float64 `json:"avg_volue"`
	Name        string  `json:"name"`
	LatestClose float64 `json:"latest_close"`
}

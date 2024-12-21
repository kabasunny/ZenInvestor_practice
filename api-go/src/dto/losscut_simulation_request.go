// api-go\src\dto\get_stock_service_request.go

package dto

type LosscutSimulationRequest struct {
	Ticker              string  `json:"ticker"`
	SimulationDate      string  `json:"simulationDate"`
	StopLossPercentage  float64 `json:"stopLossPercentage"`
	TrailingStopTrigger float64 `json:"trailingStopTrigger"`
	TrailingStopUpdate  float64 `json:"trailingStopUpdate"`
}

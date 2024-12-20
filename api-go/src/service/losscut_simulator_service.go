// api-go\src\service\losscut_simulator_service.go
package service

import (
	generate_chart_lc_sim "api-go/src/service/ms_gateway/generate_chart_lc_sim"
	"context"
	"time"
)

// LosscutSimulatorService はシミュレーション用の株価データ等を取得するためのインターフェース
type LosscutSimulatorService interface {
	GetStockChartForLCSim(ctx context.Context, ticker string, simulationDate time.Time, stopLossPercentage, trailingStopTrigger, trailingStopUpdate float64) (*generate_chart_lc_sim.GenerateChartLCResponse, float64, error)
}

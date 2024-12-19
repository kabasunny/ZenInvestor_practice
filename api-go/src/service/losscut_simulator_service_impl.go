// api-go\src\service\losscut_simulator_service.go

package service

import (
	client "api-go/src/service/ms_gateway/client"
	generate_chart_lc_sim "api-go/src/service/ms_gateway/generate_chart_lc_sim"
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LosscutSimulatorServiceImpl は LosscutSimulatorService インターフェースの実装
type LosscutSimulatorServiceImpl struct {
	clients map[string]interface{}
}

// NewLosscutSimulatorServiceImpl は LosscutSimulatorServiceImpl の新しいインスタンスを作成
func NewLosscutSimulatorServiceImpl(clients map[string]interface{}) LosscutSimulatorService {
	return &LosscutSimulatorServiceImpl{
		clients: clients,
	}
}

// GetStockChartForLCSim はシミュレーションのために株価データを取得し、ロスカットとトレーリングストップを考慮してシミュレーションを行う
func (s *LosscutSimulatorServiceImpl) GetStockChartForLCSim(ctx context.Context, ticker string, simulationDate time.Time, stopLossPercentage, trailingStopTrigger, trailingStopUpdate float64) (*generate_chart_lc_sim.GenerateChartLCResponse, float64, error) {
	stockClient := s.clients["get_stock_data"].(client.GetStockDataClient)
	generateChartClient := s.clients["generate_chart_lc_sim"].(client.GenerateChartLCClient)

	// 期間にはシミュレーション日の1年前から2年後の日付を指定
	startDate := simulationDate.AddDate(-1, 0, 0)
	endDate := simulationDate.AddDate(2, 0, 0)
	period := fmt.Sprintf("%s_%s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	req := &getstockdata.GetStockDataRequest{
		Ticker: ticker,
		Period: period,
	}

	res, err := stockClient.GetStockData(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, 0, fmt.Errorf("stock data not found: %w", err)
		}
		return nil, 0, fmt.Errorf("failed to get stock data: %w", err)
	}

	purchaseDate, purchasePrice, finalDate, finalPrice, profitLoss, err := GetLossCutSimulatorResults(res.StockData, simulationDate, stopLossPercentage, trailingStopTrigger, trailingStopUpdate)
	if err != nil {
		return nil, 0, fmt.Errorf("simulation failed: %w", err)
	}

	generateChartReq := &generate_chart_lc_sim.GenerateChartLCRequest{
		Dates:         []string{simulationDate.Format("2006-01-02"), finalDate.Format("2006-01-02")},
		ClosePrices:   []float64{purchasePrice, finalPrice},
		PurchaseDate:  purchaseDate.Format("2006-01-02"),
		PurchasePrice: purchasePrice,
		EndDate:       finalDate.Format("2006-01-02"),
		EndPrice:      finalPrice,
	}

	generateChartRes, err := generateChartClient.GenerateChart(ctx, generateChartReq)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to generate chart: %w", err)
	}

	return generateChartRes, profitLoss, nil
}

// api-go\src\service\losscut_simulator_service.go

package service

import (
	client "api-go/src/service/ms_gateway/client"
	generate_chart_lc_sim "api-go/src/service/ms_gateway/generate_chart_lc_sim"
	getstockdatawithdates "api-go/src/service/ms_gateway/get_stock_data_with_dates" // 修正されたインポート
	"context"
	"fmt"
	"sort"
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
func (s *LosscutSimulatorServiceImpl) GetStockChartForLCSim(ctx context.Context, ticker string, simulationDate string, stopLossPercentage, trailingStopTrigger, trailingStopUpdate float64) (*generate_chart_lc_sim.GenerateChartLCResponse, float64, error) {
	stockClient := s.clients["get_stock_data_with_dates"].(client.GetStockDataWithDatesClient) // 修正されたクライアント
	generateChartClient := s.clients["generate_chart_lc_sim"].(client.GenerateChartLCClient)

	// 期間にはシミュレーション日の1年前から2年後の日付を指定
	parsedSimulationDate, err := time.Parse("2006-01-02", simulationDate)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse simulation date: %w", err)
	}

	startDate := parsedSimulationDate.AddDate(0, -3, 0)
	endDate := parsedSimulationDate.AddDate(0, 6, 0)

	req := &getstockdatawithdates.GetStockDataWithDatesRequest{ // 修正されたリクエスト
		Ticker:    ticker,
		StartDate: startDate.Format("2006-01-02"),
		EndDate:   endDate.Format("2006-01-02"),
	}

	res, err := stockClient.GetStockData(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, 0, fmt.Errorf("stock data not found: %w", err)
		}
		return nil, 0, fmt.Errorf("failed to get stock data: %w", err)
	}

	fmt.Println("res OK")

	purchaseDate, purchasePrice, finalDate, finalPrice, profitLoss, err := GetLossCutSimulatorResults(res.StockData, simulationDate, stopLossPercentage, trailingStopTrigger, trailingStopUpdate)
	if err != nil {
		return nil, 0, fmt.Errorf("simulation failed: %w", err)
	}

	// resから日付と終値を取得し、日付順にソートする
	var dates []string
	var closePrices []float64
	for date, stockData := range res.StockData {
		dates = append(dates, date)
		closePrices = append(closePrices, stockData.Close)
	}

	// 日付順にソート
	sort.SliceStable(dates, func(i, j int) bool {
		return dates[i] < dates[j]
	})

	// ソートされた日付に基づいて終値を並び替える
	var sortedClosePrices []float64
	for _, date := range dates {
		sortedClosePrices = append(sortedClosePrices, res.StockData[date].Close)
	}

	generateChartReq := &generate_chart_lc_sim.GenerateChartLCRequest{
		Dates:         dates,
		ClosePrices:   sortedClosePrices,
		PurchaseDate:  purchaseDate,
		PurchasePrice: purchasePrice,
		EndDate:       finalDate,
		EndPrice:      finalPrice,
	}

	generateChartRes, err := generateChartClient.GenerateChart(ctx, generateChartReq)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to generate chart: %w", err)
	}

	return generateChartRes, profitLoss, nil
}

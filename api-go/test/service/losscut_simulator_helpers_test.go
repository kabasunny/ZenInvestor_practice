// api-go\test\service\losscut_simulator_helpers_test.go
package service

import (
	"testing"

	"api-go/src/service"                                      // 正しいインポート
	"api-go/src/service/ms_gateway/get_stock_data_with_dates" // 正しいインポート

	"github.com/stretchr/testify/assert"
)

func TestGetLossCutSimulatorResults(t *testing.T) {
	stockData := map[string]*get_stock_data_with_dates.StockDataWithDates{
		"2023-01-01": {Open: 100.0, Close: 105.0, High: 110.0, Low: 96.0, Volume: 1000},
		"2023-01-02": {Open: 105.0, Close: 115.0, High: 120.0, Low: 100.0, Volume: 1500},
		"2023-01-03": {Open: 115.0, Close: 120.0, High: 125.0, Low: 113.0, Volume: 2000},
		"2023-01-04": {Open: 120.0, Close: 125.0, High: 130.0, Low: 118.0, Volume: 2500},
		"2023-01-05": {Open: 125.0, Close: 128.0, High: 133.0, Low: 120.0, Volume: 3000},
		"2023-01-06": {Open: 130.0, Close: 130.0, High: 135.0, Low: 128.0, Volume: 3500},
	}

	startDate := "2023-01-01"
	stopLossPercentage := 5.0
	trailingStopTrigger := 10.0
	trailingStopUpdate := 2.0

	purchaseDate, purchasePrice, endDate, endPrice, profitLoss, err := service.GetLossCutSimulatorResults(stockData, startDate, stopLossPercentage, trailingStopTrigger, trailingStopUpdate)
	assert.NoError(t, err)
	assert.Equal(t, "2023-01-01", purchaseDate)
	assert.Equal(t, 100.0, purchasePrice)
	assert.Equal(t, "2023-01-06", endDate)
	assert.Equal(t, 130.0, endPrice)
	assert.Equal(t, 30.0, profitLoss)
}

func TestGetLossCutSimulatorResults_StartDateOutOfRange(t *testing.T) {
	// テストデータを作成
	stockData := map[string]*get_stock_data_with_dates.StockDataWithDates{
		"2023-01-01": {Open: 100.0, Close: 105.0, High: 110.0, Low: 95.0, Volume: 1000},
		"2023-01-02": {Open: 105.0, Close: 110.0, High: 115.0, Low: 100.0, Volume: 1500},
		"2023-01-03": {Open: 110.0, Close: 115.0, High: 120.0, Low: 105.0, Volume: 2000},
	}

	startDate := "2025-01-01"
	stopLossPercentage := 5.0
	trailingStopTrigger := 10.0
	trailingStopUpdate := 2.0

	// 関数を呼び出し
	_, _, _, _, _, err := service.GetLossCutSimulatorResults(stockData, startDate, stopLossPercentage, trailingStopTrigger, trailingStopUpdate)
	assert.EqualError(t, err, "開始日がデータの範囲外です。無限ループを防ぐため、処理を中断")
}

// go test -v ./test/service/losscut_simulator_helpers_test.go -run TestGetLossCutSimulatorResults

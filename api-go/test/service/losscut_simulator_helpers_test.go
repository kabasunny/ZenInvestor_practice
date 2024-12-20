// api-go\test\service\losscut_simulator_helpers_test.go

package service_test

import (
	"api-go/src/service"
	getstockdatawithdates "api-go/src/service/ms_gateway/get_stock_data_with_dates" // 修正されたインポート
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLossCutSimulatorResults(t *testing.T) {
	// テストデータの作成
	stockData := map[string]*getstockdatawithdates.StockDataWithDates{
		"2023-01-01": {Open: 100, Close: 105, High: 106, Low: 99, Volume: 1000},
		"2023-01-02": {Open: 105, Close: 104, High: 108, Low: 102, Volume: 1500},
		"2023-01-03": {Open: 104, Close: 107, High: 110, Low: 103, Volume: 1200},
		"2023-01-04": {Open: 107, Close: 103, High: 108, Low: 101, Volume: 1100},
		"2023-01-05": {Open: 103, Close: 102, High: 105, Low: 100, Volume: 900},
		"2023-01-06": {Open: 102, Close: 101, High: 104, Low: 99, Volume: 800},
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	stopLossPercentage := 10.0
	trailingStopTrigger := 5.0
	trailingStopUpdate := 3.0

	// 関数の呼び出し
	purchaseDate, purchasePrice, endDate, endPrice, profitLoss, err := service.GetLossCutSimulatorResults(stockData, startDate, stopLossPercentage, trailingStopTrigger, trailingStopUpdate)

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, startDate, purchaseDate)
	assert.Equal(t, 100.0, purchasePrice)
	assert.NotZero(t, endDate)
	assert.Greater(t, endPrice, 0.0)
	assert.NotZero(t, profitLoss)
}

func TestGetLossCutSimulatorResults_StartDateOutOfRange(t *testing.T) {
	// テストデータの作成
	stockData := map[string]*getstockdatawithdates.StockDataWithDates{
		"2023-01-01": {Open: 100, Close: 105, High: 106, Low: 99, Volume: 1000},
		"2023-01-02": {Open: 105, Close: 104, High: 108, Low: 102, Volume: 1500},
		"2023-01-03": {Open: 104, Close: 107, High: 110, Low: 103, Volume: 1200},
	}

	startDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC) // データの範囲外の日付

	// 関数の呼び出し
	_, _, _, _, _, err := service.GetLossCutSimulatorResults(stockData, startDate, 10.0, 5.0, 3.0)

	// アサーション
	assert.Error(t, err)
	assert.Equal(t, "開始日がデータの範囲外です。無限ループを防ぐため、処理を中断", err.Error())
}

func TestRound(t *testing.T) {
	// テストケース
	tests := []struct {
		val       float64
		precision int
		expected  float64
	}{
		{123.456789, 2, 123.46},
		{123.454, 2, 123.45},
		{123.4, 2, 123.40},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("round(%f, %d)", tt.val, tt.precision), func(t *testing.T) {
			result := service.Round(tt.val, tt.precision)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// go test -v ./test/service/losscut_simulator_helpers_test.go -run TestGetLossCutSimulatorResults

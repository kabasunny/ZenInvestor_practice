package service

import (
	sma "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average"
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
)

// convertStockDataForSMA は get_stock_data.StockData を simple_moving_average.StockDataForSMA に変換するヘルパー関数
func convertStockDataForSMA(stockData map[string]*getstockdata.StockData) map[string]*sma.StockDataForSMA {
	converted := make(map[string]*sma.StockDataForSMA)
	for date, data := range stockData {
		converted[date] = &sma.StockDataForSMA{
			Open:   data.Open,
			Close:  data.Close,
			High:   data.High,
			Low:    data.Low,
			Volume: data.Volume,
		}
	}
	return converted
}

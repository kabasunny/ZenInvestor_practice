package controller

import (
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
)

// StockController は株価データを管理するためのインターフェース
type StockController interface {
	GetStockData(ctx context.Context, ticker string, period string) (*ms_gateway.GetStockDataResponse, error)
}

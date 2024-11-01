package service

import (
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
)

// StockService は株価データを取得するためのインターフェース
type StockService interface {
	GetStockData(ctx context.Context, ticker string, period string) (*ms_gateway.GetStockDataResponse, error)
}

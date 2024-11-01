package service

import (
	"context"

	"api-go/src/service/gateway"
)

// StockService は株価データを取得するためのインターフェース
type StockService interface {
	GetStockData(ctx context.Context, ticker string, period string) (*gateway.GetStockResponse, error)
}

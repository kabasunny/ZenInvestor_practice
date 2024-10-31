package service

import (
	"context"

	"api-go/src/service/gateway"
)

// StockService は株価データを取得するためのインターフェースです。
type StockService interface {
	GetStockData(ctx context.Context, ticker string, period string) (*gateway.GetStockResponse, error)
}

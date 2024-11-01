package controller

import (
	"api-go/src/service/gateway"
	"context"
)

// StockController は株価データを管理するためのインターフェース
type StockController interface {
	GetStockData(ctx context.Context, ticker string, period string) (*gateway.GetStockResponse, error)
}

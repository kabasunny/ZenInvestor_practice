package controller

import (
	"api-go/src/service"
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
)

// stockControllerImpl は StockController インターフェースの実装
type stockControllerImpl struct {
	stockService service.StockService
}

// NewStockController は StockController の新しいインスタンスを作成
func NewStockControllerImpl(stockService service.StockService) StockController {
	return &stockControllerImpl{stockService: stockService}
}

// GetStockData は指定された銘柄と期間の株価データを取得
func (c *stockControllerImpl) GetStockData(ctx context.Context, ticker string, period string) (*ms_gateway.GetStockDataResponse, error) {
	return c.stockService.GetStockData(ctx, ticker, period)
}

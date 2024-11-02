package controller

import (
	"api-go/src/service"
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// stockControllerImpl は StockController インターフェースの実装
type stockControllerImpl struct {
	stockService service.StockService
}

// NewStockControllerImpl は StockController の新しいインスタンスを作成
func NewStockControllerImpl(stockService service.StockService) StockController {
	return &stockControllerImpl{stockService: stockService}
}

// GetStockData は指定された銘柄と期間の株価データを取得
func (c *stockControllerImpl) GetStockData(ctx context.Context, ticker string, period string) (*ms_gateway.GetStockDataResponse, error) {
	return c.stockService.GetStockData(ctx, ticker, period)
}

// GetStockDataHandler はHTTPリクエストを処理し、株価データを取得する
func GetStockDataHandler(c *gin.Context, stockController StockController) {
	ctx := context.Background()
	ticker := c.Query("ticker")
	period := c.Query("period")

	response, err := stockController.GetStockData(ctx, ticker, period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stockData": response})
}

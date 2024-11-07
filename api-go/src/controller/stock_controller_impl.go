package controller

import (
	"api-go/src/service"
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

// GetStockDataHandler はHTTPリクエストを処理し、株価データを取得する
func (c *stockControllerImpl) GetStockData(ctx *gin.Context) {
	reqCtx := context.Background() // リクエストコンテキスト
	ticker := ctx.Query("ticker")
	period := ctx.Query("period")

	response, err := c.stockService.GetStockData(reqCtx, ticker, period) // サービスを直接呼び出す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"stockData": response})
}

package controller

import (
	"api-go/src/dto" // DTOのパッケージ
	"api-go/src/service"
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
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

// 株価データを取得する
func (c *stockControllerImpl) GetStockData(ctx *gin.Context) {
	reqCtx := context.Background() // リクエストコンテキスト
	ticker := ctx.Query("ticker")
	period := ctx.Query("period")

	response, err := c.stockService.GetStockData(reqCtx, ticker, period) // サービスを直接呼び出す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスデータの変換
	transformedResponse := transformStockDataResponse(response)

	ctx.JSON(http.StatusOK, gin.H{"stock_data": transformedResponse}) // ラベル名はフロントとの兼ね合いに注意
}

// transformStockDataResponse は stock_data を直接 stockData に格納するための変換関数
func transformStockDataResponse(response *getstockdata.GetStockDataResponse) map[string]interface{} {
	transformed := make(map[string]interface{})

	for date, data := range response.StockData {
		transformed[date] = data
	}

	return transformed
}

// 株価チャート可視化データを取得する
func (c *stockControllerImpl) GetStockChart(ctx *gin.Context) {
	var req dto.GetStockServiceRequest // DTOを使用

	// GET
	// if err := ctx.ShouldBindQuery(&req); err != nil { // URLクエリパラメータをバインド
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// package routerのSetupRouterメソッド内でGETメソッドに変更して、以下でテスト可能
	// http://localhost:8086/getStockData?ticker=AAPL&period=1y&indicators[0][type]=SMA&indicators[0][params][window_size]=20

	// POST
	if err := ctx.ShouldBindJSON(&req); err != nil { // リクエストボディからJSONをバインド
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// DTOのバリデーション (必要に応じて)
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.stockService.GetStockChart(ctx, req.Ticker, req.Period, req.Indicators) // DTOを渡す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

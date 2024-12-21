// api-go\src\controller\stock_controller_impl.go

package controller

import (
	"api-go/src/dto" // DTOのパッケージ
	"api-go/src/service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// stockControllerImpl は StockController インターフェースの実装
type stockControllerImpl struct {
	stockService      service.StockService
	rankingService    service.RankingService
	losscutSIMService service.LosscutSimulatorService
}

// NewStockControllerImpl は StockController の新しいインスタンスを作成
func NewStockControllerImpl(stockService service.StockService, rankingService service.RankingService, losscutSIMService service.LosscutSimulatorService) StockController {
	return &stockControllerImpl{
		stockService:      stockService,
		rankingService:    rankingService,
		losscutSIMService: losscutSIMService,
	}
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

	// フロントエンド向けにStockNameも返す
	ctx.JSON(http.StatusOK, gin.H{
		"stock_data": response.StockData, // 変換不要
		"stock_name": response.StockName, // 銘柄名を追加
	})
}

// 株価チャート可視化データを取得する
func (c *stockControllerImpl) GetStockChart(ctx *gin.Context) {
	var req dto.GetStockServiceRequest // DTOを使用

	// GET
	// if err := ctx.ShouldBindQuery(&req); err != nil { // URLクエリパラメータをバインド
	//  ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//  return
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

	res, err := c.stockService.GetStockChart(ctx, req.Ticker, req.Period, req.Indicators, req.IncludeVolume) // DTOを渡す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetTop100Ranking はトップ100のランキングデータを取得する
func (c *stockControllerImpl) GetInitialRanking(ctx *gin.Context) {
	// context.Context を取得
	// reqCtx := ctx.Request.Context()
	rankingData, err := c.rankingService.GetTop100RankingData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rankingData)
}

// ランキングデータ取得
func (c *stockControllerImpl) GetRankingDataByRange(ctx *gin.Context) {
	// クエリパラメータの取得とバインド
	startRankStr := ctx.Query("startRank")
	endRankStr := ctx.Query("endRank")

	startRank, err := strconv.Atoi(startRankStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startRank parameter"})
		return
	}

	endRank, err := strconv.Atoi(endRankStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endRank parameter"})
		return
	}

	// context.Context を取得
	// reqCtx := ctx.Request.Context()

	// サービスの呼び出し
	res, err := c.rankingService.GetRankingDataByRange(startRank, endRank)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ロスカットシミュレーション結果を取得
func (c *stockControllerImpl) GetLosscutSimulation(ctx *gin.Context) {
	reqCtx := context.Background() // リクエストコンテキスト

	var request dto.LosscutSimulationRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	generateChartRes, profitLoss, err := c.losscutSIMService.
		GetStockChartForLCSim(reqCtx, request.Ticker, request.SimulationDate, request.StopLossPercentage, request.TrailingStopTrigger, request.TrailingStopUpdate) // サービスを直接呼び出す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"chart_data": generateChartRes,
		"profitLoss": profitLoss,
	})
}

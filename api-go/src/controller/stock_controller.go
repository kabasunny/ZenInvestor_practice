package controller

import "github.com/gin-gonic/gin"

// StockController は株価データを管理するためのインターフェース
type StockController interface {
	// 株価データ取得
	GetStockData(ctx *gin.Context)

	// 株価チャート取得
	GetStockChart(ctx *gin.Context)
}

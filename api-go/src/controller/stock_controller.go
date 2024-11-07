package controller

import "github.com/gin-gonic/gin"

// StockController は株価データを管理するためのインターフェース
type StockController interface {
	GetStockData(ctx *gin.Context)
}

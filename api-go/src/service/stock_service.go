package service

import (
	indicator "api-go/src/service/ms_gateway/calculate_indicator" // IndicatorParamsのimport元
	gc "api-go/src/service/ms_gateway/generate_chart"
	"context"
)

// StockService は株価データを取得するためのインターフェース
type StockService interface {
	GetStockData(ctx context.Context, ticker string, period string, indicators []*indicator.IndicatorParams) (*gc.GenerateChartResponse, error)
	// SimpleMovingAverageの実装時に、指標の指定を行う引数indicators []stringを追加した　指標の指定は複数あるので、配列で受け取る
	// 指標の指定を行う引数indicators []string　でいくつ受け取るかは、後にGeneratChartの実装時に検討する
}

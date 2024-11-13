package dto

import (
	indicator "api-go/src/service/ms_gateway/calculate_indicator" // IndicatorParamsのimport元
)

// GetStockServiceRequest は株価データ取得リクエストのDTO
type GetStockServiceRequest struct {
	Ticker        string                       `form:"ticker" binding:"required"`
	Period        string                       `form:"period" binding:"required"`
	Indicators    []*indicator.IndicatorParams `form:"indicators"`
	IncludeVolume bool                         `form:"includeVolune"`
}

// Validate は GetStockDataRequest のバリデーションを行う
func (r *GetStockServiceRequest) Validate() error {
	// バリデーションルールを実装 (例: tickerの形式チェックなど)
	return nil // エラーがない場合はnilを返す
}

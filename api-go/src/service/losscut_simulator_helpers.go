// api-go\src\service\losscut_simulator_helpers.go
package service

import (
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
	"errors"
	"fmt"
	"math"
	"time"
)

// GetLossCutSimulatorResults はシミュレーションの結果を計算
func GetLossCutSimulatorResults(stockData map[string]*getstockdata.StockData, startDate time.Time, stopLossPercentage, trailingStopTrigger, trailingStopUpdate float64) (time.Time, float64, time.Time, float64, float64, error) {
	// データの最終日付を取得
	var maxDate time.Time
	for date := range stockData {
		d, err := time.Parse("2006-01-02", date)
		if err != nil {
			return time.Time{}, 0, time.Time{}, 0, 0, fmt.Errorf("failed to parse date: %w", err)
		}
		if d.After(maxDate) {
			maxDate = d
		}
	}

	// データが存在する最初の日付を取得
	for {
		if startDate.After(maxDate) {
			return time.Time{}, 0, time.Time{}, 0, 0, errors.New("開始日がデータの範囲外です。無限ループを防ぐため、処理を中断")
		}
		if _, exists := stockData[startDate.Format("2006-01-02")]; exists {
			break
		}
		startDate = startDate.AddDate(0, 0, 1) // データに存在する日付になるまで日付を進める
	}

	// 購入初日の設定
	purchaseDate := startDate
	purchasePrice := stockData[purchaseDate.Format("2006-01-02")].Open              // 取引開始日の始値
	stopLossThreshold := round(purchasePrice*(1-stopLossPercentage/100), 1)         // 初期ロスカット値
	trailingStopTriggerPrice := round(purchasePrice*(1+trailingStopTrigger/100), 1) // 初期トレーリングストップ値

	// 初期化
	var endDate time.Time
	var endPrice float64

	// 取引日ごとの確認
	for dateStr, data := range stockData {
		currentDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return time.Time{}, 0, time.Time{}, 0, 0, fmt.Errorf("failed to parse date: %w", err)
		}

		if currentDate.Before(startDate) {
			continue
		}
		openPrice := data.Open
		lowPrice := data.Low
		closePrice := data.Close

		// ロスカット条件: 当日の始値がロスカット値以下
		if openPrice <= stopLossThreshold {
			endPrice = openPrice
			endDate = currentDate
			break
		}

		// トレーリングストップ発動条件: 当日の安値がトレーリングストップ値以下
		if lowPrice <= stopLossThreshold {
			endPrice = lowPrice
			endDate = currentDate
			break
		}

		// トレーリングストップ更新条件: 終値がトリガーを超えた場合
		if closePrice >= trailingStopTriggerPrice {
			stopLossThreshold = round(closePrice*(1-trailingStopUpdate/100), 1)
			trailingStopTriggerPrice = round(closePrice*(1+trailingStopTrigger/100), 1)
		}
	}

	// 取引終了条件が満たされなかった場合
	if endDate.IsZero() {
		endPrice = stockData[maxDate.Format("2006-01-02")].Close // 最終日の終値
		endDate = maxDate                                        // 最終日
	}

	// 損益の計算
	profitLoss := round((endPrice-purchasePrice)/purchasePrice*100, 1)

	// 結果の返却
	return purchaseDate, purchasePrice, endDate, endPrice, profitLoss, nil
}

// round は小数点以下の桁数を指定して四捨五入する関数
func round(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

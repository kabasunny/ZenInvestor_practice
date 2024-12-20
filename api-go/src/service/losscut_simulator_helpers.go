// api-go\src\service\losscut_simulator_helpers.go
package service

import (
	getstockdatawithdates "api-go/src/service/ms_gateway/get_stock_data_with_dates"
	"fmt"
	"math"
	"sort"
	"time"
)

// GetLossCutSimulatorResults はシミュレーションの結果を計算
func GetLossCutSimulatorResults(stockData map[string]*getstockdatawithdates.StockDataWithDates, startDateString string, stopLossPercentage, trailingStopTrigger, trailingStopUpdate float64) (string, float64, string, float64, float64, error) {
	fmt.Println("startDate : ", startDateString)
	fmt.Println("len(stockData) : ", len(stockData))

	// データの最終日付と最初の日付を取得
	var maxDate, minDate string
	for date := range stockData {
		if date > maxDate {
			maxDate = date
		}
		if minDate == "" || date < minDate {
			minDate = date
		}
	}

	fmt.Println("maxDate : ", maxDate)
	fmt.Println("minDate : ", minDate)
	fmt.Println("startDate : ", startDateString)

	// 開始日がデータの範囲内にあるか確認
	if startDateString < minDate {
		startDateString = minDate
	}
	if startDateString > maxDate {
		return "", 0, "", 0, 0, fmt.Errorf("開始日がデータの範囲外です。無限ループを防ぐため、処理を中断")
	}

	// データが存在するまで次の日に進める
	for startDateString <= maxDate {
		if data, exists := stockData[startDateString]; exists {
			if data != nil {
				break
			}
		}
		parsedDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			fmt.Println("Invalid date format for startDateString:", startDateString)
			return "", 0, "", 0, 0, fmt.Errorf("invalid date format: %w", err)
		}
		startDateString = parsedDate.AddDate(0, 0, 1).Format("2006-01-02")
	}

	fmt.Println("Updated startDate: ", startDateString)

	// データを日付順にソート
	sortedDates := make([]string, 0, len(stockData))
	for date := range stockData {
		sortedDates = append(sortedDates, date)
	}
	sort.Strings(sortedDates)

	// 購入初日の設定
	if stockData[startDateString] == nil {
		return "", 0, "", 0, 0, fmt.Errorf("データが見つかりません: %s", startDateString)
	}
	purchaseDate := startDateString
	purchasePrice := stockData[startDateString].Open                                // 取引開始日の始値
	stopLossThreshold := round(purchasePrice*(1-stopLossPercentage/100), 1)         // 初期ロスカット値
	trailingStopTriggerPrice := round(purchasePrice*(1+trailingStopTrigger/100), 1) // 初期トレーリングストップ値

	// 初期化
	var endDate string
	var endPrice float64

	// 取引日ごとの確認
	for _, dateStr := range sortedDates {
		if dateStr < startDateString {
			continue
		}
		data := stockData[dateStr]
		openPrice := data.Open
		lowPrice := data.Low
		closePrice := data.Close

		// ロスカット条件: 当日の始値がロスカット値以下
		if openPrice <= stopLossThreshold {
			endPrice = openPrice
			endDate = dateStr
			break
		}

		// トレーリングストップ発動条件: 当日の安値がトレーリングストップ値以下
		if lowPrice <= stopLossThreshold {
			endPrice = lowPrice
			endDate = dateStr
			break
		}

		// トレーリングストップ更新条件: 終値がトリガーを超えた場合
		if closePrice >= trailingStopTriggerPrice {
			stopLossThreshold = round(closePrice*(1-trailingStopUpdate/100), 1)
			trailingStopTriggerPrice = round(closePrice*(1+trailingStopTrigger/100), 1)
		}
	}

	// 取引終了条件が満たされなかった場合
	if endDate == "" {
		endPrice = stockData[maxDate].Close // 最終日の終値
		endDate = maxDate                   // 最終日
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

// Round は外部から呼び出すための公開関数として round をラップします
func Round(val float64, precision int) float64 {
	return round(val, precision)
}

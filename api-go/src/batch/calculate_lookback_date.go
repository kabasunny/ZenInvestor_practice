// api-go\src\batch\calculate_lookback_date.go
package batch

import (
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	get_trading_calendar_jq "api-go/src/service/ms_gateway/get_trading_calendar_jq"
	"context"
	"fmt"
	"sort"
	"time"
)

func CalculateLookbackDate(ctx context.Context,
	jdpRepo repository.JpDailyPriceRepository,
	// startDate string,
	jsiRepo repository.JpStockInfoRepository,
	lookbackDays int,
	gtcjClient client.GetTradingCalendarJqClient,
) ([]string, error) {

	// 銘柄テーブルが更新されている前提で、銘柄データは取得可能な最新の日付を保持している
	startDate, err := jsiRepo.GetLatestDate()
	if err != nil {
		return nil, fmt.Errorf("failed to get startDate: %w", err)
	}

	// startDate を time.Time 型に変換
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse startDate: %w", err)
	}

	// 株価データが保持している最新の日付 latestDate の取得とエラーチェック
	latestDate, err := jdpRepo.GetLatestDate() // DB内の銘柄情報として保持する日付を取得
	if err != nil || latestDate == "" {
		fmt.Println("GetLatestDate error or latestDate is empty")
		latestDate = start.AddDate(0, 0, -7).Format("2006-01-02")
	}

	// latestDate を time.Time 型に変換
	latest, err := time.Parse("2006-01-02", latestDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse latestDate: %w", err)
	}

	// 日付の差を計算
	dateDifference := int(start.Sub(latest).Hours() / 24)

	// 必要日数lookbackDays を調整
	trueLookbackDays := lookbackDays
	if dateDifference < lookbackDays {
		trueLookbackDays = dateDifference
	}

	// 必要日数が0の場合
	if trueLookbackDays == 0 {
		fmt.Println("trueLookbackDays==0 : 最新データ取得済み")
		return nil, nil
	}

	// 必要日数lookbackDays + 7日分 の営業日を確認する
	reqDate := start.AddDate(0, 0, -(trueLookbackDays + 7)).Format("2006-01-02")

	// trueLookbackDays個の営業日の日付の文字列リストを返す
	req := &get_trading_calendar_jq.GetTradingCalendarJqRequest{
		FromDate: reqDate,   // 取得可能最新日から逆算した日から
		ToDate:   startDate, // 取得可能最新日まで
	}

	res, err := gtcjClient.GetTradingCalendarJq(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get trading calendar: %w", err)
	}

	// 日付の文字列リストを作成
	var businessDays []string
	for _, tradingCalendar := range res.TradingCalendar {
		if tradingCalendar.HolidayDivision == "1" { // 営業日を判定
			// 非営業日:0 営業日:1 東証半日立会日:2 非営業日(祝日取引あり):3
			businessDays = append(businessDays, tradingCalendar.Date)
		}
	}

	// 日付の昇順にソート
	sort.Slice(businessDays, func(i, j int) bool {
		return businessDays[i] < businessDays[j]
	})

	// 最新の日付から指定された営業日数を抽出し、再び昇順にソート
	if len(businessDays) > trueLookbackDays {
		businessDays = businessDays[len(businessDays)-trueLookbackDays:]
	}

	return businessDays, nil
}

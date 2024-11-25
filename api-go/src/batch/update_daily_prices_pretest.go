// api-go\src\batch\update_daily_prices_pretest.go

package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"api-go/src/util" // 追加
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
	// その他必要なインポート
)

func PreTestUpdateDailyPrices(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, jdpRepo repository.JpDailyPriceRepository, clients map[string]interface{}, startDate, endDate string) error {
	// 関数全体の処理開始時刻
	startTimeOverall := time.Now()

	gsdwdClient, ok := clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	if !ok {
		return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	}

	// 指定された銘柄
	symbols := []string{"7997", "6932"} // 適当な2銘柄

	// デバッグ: 指定された日付範囲を確認
	startDate = "2024-11-19"
	endDate = "2024-11-22"
	fmt.Printf("指定された日付範囲: 開始日: %s, 終了日: %s\n", startDate, endDate)

	// シンボル抽出の処理時間
	startTimeTicker := time.Now()
	endTimeTicker := time.Now()
	fmt.Printf("シンボルを抽出の処理時間: %s\n", endTimeTicker.Sub(startTimeTicker))
	fmt.Printf("選択されたシンボル: %v\n", symbols) // シンボルのデバッグログ

	// 共通のチャンク分割関数を使用してチャンク分割
	chunks := util.ChunkSymbols(symbols, 5) // 2銘柄なので1つのチャンク

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error

	for i, chunk := range chunks {
		wg.Add(1)
		go func(chunk []string, batchNumber int) {
			defer wg.Done()

			// 複数のシンボルのデータを一度に取得の処理時間
			startTimeDownload := time.Now()
			req := &gsdwd.GetStocksDatalistWithDatesRequest{
				Symbols:   chunk,
				StartDate: startDate,
				EndDate:   endDate,
			}

			// デバッグ: リクエスト内容を確認
			fmt.Printf("バッチ %d のリクエスト: シンボル: %s, 開始日: %s, 終了日: %s\n", batchNumber, strings.Join(chunk, ", "), req.StartDate, req.EndDate)

			gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
			endTimeDownload := time.Now()
			fmt.Printf("バッチ %d のデータ取得の処理時間: %s\n", batchNumber, endTimeDownload.Sub(startTimeDownload))

			if err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("failed to get stocks data list with dates: %w", err)
				mu.Unlock()
				return
			}

			if len(gsdwdResponse.StockPrices) == 0 {
				fmt.Printf("バッチ %d のレスポンスにデータが含まれていません\n", batchNumber)
			} else {
				fmt.Printf("バッチ %d のレスポンス: %+v\n", batchNumber, gsdwdResponse) // レスポンスのデバッグログ
			}

			// シンボルごとにデータを処理の処理時間
			startTimeProcessing := time.Now()
			var newDailyPrices []model.JpDailyPrice
			for _, data := range gsdwdResponse.StockPrices {
				date, err := time.Parse("2006-01-02", data.Date)
				if err != nil {
					mu.Lock()
					overallErr = fmt.Errorf("failed to parse date: %w", err)
					mu.Unlock()
					return
				}
				dp := model.JpDailyPrice{
					Ticker:   data.Symbol,
					Date:     date,
					Open:     data.Open,
					Close:    data.Close,
					High:     data.High,
					Low:      data.Low,
					Volume:   data.Volume,
					Turnover: data.Turnover,
				}
				newDailyPrices = append(newDailyPrices, dp)
			}
			endTimeProcessing := time.Now()
			fmt.Printf("バッチ %d のデータ処理の処理時間: %s\n", batchNumber, endTimeProcessing.Sub(startTimeProcessing))

			if err := jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("failed to add daily price data: %w", err)
				mu.Unlock()
				return
			}

			if len(newDailyPrices) > 0 {
				fmt.Printf("Batch %d completed successfully. Last added data: Ticker: %s, Date: %s\n",
					batchNumber, newDailyPrices[len(newDailyPrices)-1].Ticker, newDailyPrices[len(newDailyPrices)-1].Date)
			} else {
				fmt.Printf("Batch %d completed with no data.\n", batchNumber)
			}
		}(chunk, i+1)
	}

	wg.Wait()

	if overallErr != nil {
		return overallErr
	}

	if err := udsRepo.UpdateStatus("jp_daily_price"); err != nil {
		return fmt.Errorf("failed to update status for jp_daily_price: %w", err)
	}

	// 関数全体の処理終了時刻
	endTimeOverall := time.Now()
	fmt.Printf("全体の処理時間: %s\n", endTimeOverall.Sub(startTimeOverall))

	return nil
}

// api-go\src\batch\update_daily_prices_pretest.go

package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
	// その他必要なインポート
)

func PreTestUpdateDailyPrices(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, jdpRepo repository.JpDailyPriceRepository, clients map[string]interface{}, symbols []string, startDate, endDate string) error {
	// 関数全体の処理開始時刻
	startTimeOverall := time.Now()

	gsdwdClient, ok := clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	if !ok {
		return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	}

	fmt.Printf("選択されたシンボル: %v\n", symbols) // シンボルのデバッグログ

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error

	wg.Add(1)
	go func() {
		defer wg.Done()

		// 複数のシンボルのデータを一度に取得の処理時間
		startTimeDownload := time.Now()
		req := &gsdwd.GetStocksDatalistWithDatesRequest{
			Symbols:   symbols,
			StartDate: startDate,
			EndDate:   endDate,
		}

		// デバッグ: リクエスト内容を確認
		fmt.Printf("リクエスト: シンボル: %s, 開始日: %s, 終了日: %s\n", strings.Join(symbols, ", "), req.StartDate, req.EndDate)

		gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
		endTimeDownload := time.Now()
		fmt.Printf("データ取得の処理時間: %s\n", endTimeDownload.Sub(startTimeDownload))

		if err != nil {
			mu.Lock()
			overallErr = fmt.Errorf("failed to get stocks data list with dates: %w", err)
			mu.Unlock()
			return
		}

		if len(gsdwdResponse.StockPrices) == 0 {
			fmt.Printf("レスポンスにデータが含まれていません\n")
		} else {
			fmt.Printf("レスポンス: %+v\n", gsdwdResponse) // レスポンスのデバッグログ
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
				Symbol:   data.Symbol,
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
		fmt.Printf("データ処理の処理時間: %s\n", endTimeProcessing.Sub(startTimeProcessing))

		if err := jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
			mu.Lock()
			overallErr = fmt.Errorf("failed to add daily price data: %w", err)
			mu.Unlock()
			return
		}

		if len(newDailyPrices) > 0 {
			fmt.Printf("Completed successfully. Last added data: Ticker: %s, Date: %s\n",
				newDailyPrices[len(newDailyPrices)-1].Symbol, newDailyPrices[len(newDailyPrices)-1].Date)
		} else {
			fmt.Printf("Completed with no data.\n")
		}
	}()

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

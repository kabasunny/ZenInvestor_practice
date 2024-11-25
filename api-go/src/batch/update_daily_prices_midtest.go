// api-go\src\batch\update_daily_prices_midtest.go

package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"api-go/src/util"
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MidTestUpdateDailyPrices(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, jdpRepo repository.JpDailyPriceRepository, clients map[string]interface{}, startDate, endDate string) error {
	// 関数全体の処理開始時刻
	startTimeOverall := time.Now()

	gsdwdClient, ok := clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	if !ok {
		return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	}

	// 全銘柄を取得の処理時間
	startTimeGetStocks := time.Now()
	stocks, err := jsiRepo.GetAllStockInfo()
	endTimeGetStocks := time.Now()
	fmt.Printf("全銘柄を取得の処理時間: %s\n", endTimeGetStocks.Sub(startTimeGetStocks))

	if err != nil {
		return fmt.Errorf("failed to get all stock info: %w", err)
	}

	// 全銘柄からランダムに ～銘柄を選択
	var symbols []string
	for _, stock := range *stocks {
		symbols = append(symbols, stock.Ticker)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(symbols), func(i, j int) {
		symbols[i], symbols[j] = symbols[j], symbols[i]
	})
	// 50銘柄に対応する範囲で選択
	selectedSymbols := symbols[:50] // ここを50に変更
	fmt.Printf("選択された銘柄数: %d 銘柄\n", len(selectedSymbols))

	// 銘柄を50銘柄ずつに分割
	chunks := util.ChunkSymbols(selectedSymbols, 50) // 分割サイズを50に設定

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error

	for i, chunk := range chunks {
		wg.Add(1)
		go func(chunk []string, batchNumber int) {
			defer wg.Done()

			// 複数の銘柄のデータを一度に取得の処理時間
			startTimeDownload := time.Now()
			req := &gsdwd.GetStocksDatalistWithDatesRequest{
				Symbols:   chunk,
				StartDate: startDate,
				EndDate:   endDate,
			}

			fmt.Printf("バッチ %d のリクエスト: シンボル: %v, 開始日: %s, 終了日: %s\n", batchNumber, chunk, startDate, endDate)
			gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
			endTimeDownload := time.Now()
			fmt.Printf("バッチ %d のデータ取得の処理時間: %s\n", batchNumber, endTimeDownload.Sub(startTimeDownload))

			if err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("failed to get stocks data list with dates: %w", err)
				mu.Unlock()
				return
			}

			// データ処理
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

			fmt.Printf("バッチ %d が完了しました。処理されたデータ数: %d\n", batchNumber, len(newDailyPrices))
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

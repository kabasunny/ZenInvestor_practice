// api-go\src\batch\update_daily_prices.go
package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"context"
	"fmt"
	"sync"
	"time"
	// その他必要なインポート
)

func chunkSymbols(symbols []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(symbols); i += chunkSize {
		end := i + chunkSize
		if end > len(symbols) {
			end = len(symbols)
		}
		chunks = append(chunks, symbols[i:end])
	}
	return chunks
}

func UpdateDailyPrices(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, jdpRepo repository.JpDailyPriceRepository, clients map[string]interface{}, startDate, endDate string) error {
	gsdwdClient, ok := clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	if !ok {
		return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	}

	var symbols []string
	stocks, err := jsiRepo.GetAllStockInfo()
	if err != nil {
		return fmt.Errorf("failed to get all stock info: %w", err)
	}

	for _, stock := range *stocks {
		symbol := stock.Ticker + ".T"
		symbols = append(symbols, symbol)
	}

	chunks := chunkSymbols(symbols, 5) // 1バッチ5銘柄に分割

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error

	for i, chunk := range chunks {
		// 最大5バッチで処理を終了
		if i >= 5 {
			break
		}

		wg.Add(1)
		go func(chunk []string, batchNumber int) {
			defer wg.Done()

			// リクエスト間の遅延を追加
			time.Sleep(2 * time.Second)

			req := &gsdwd.GetStocksDatalistWithDatesRequest{
				Symbols:   chunk,
				StartDate: startDate,
				EndDate:   endDate,
			}
			fmt.Printf("Batch %d request: %v\n", batchNumber, req) // リクエストの詳細を表示

			gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
			if err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("failed to get stocks data list with dates: %w", err)
				mu.Unlock()
				return
			}

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

			fmt.Printf("Batch %d new data: %v\n", batchNumber, newDailyPrices) // 新規データの詳細を表示

			if err := jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("failed to add daily price data: %w", err)
				mu.Unlock()
				return
			}

			// 最後のデータをコンソールに表示
			if len(newDailyPrices) > 0 {
				lastPrice := newDailyPrices[len(newDailyPrices)-1]
				fmt.Printf("Batch %d completed successfully. Last added data: Ticker: %s, Date: %s, Open: %.2f, Close: %.2f, High: %.2f, Low: %.2f, Volume: %d, Turnover: %.2f\n",
					batchNumber, lastPrice.Ticker, lastPrice.Date.Format("2006-01-02"), lastPrice.Open, lastPrice.Close, lastPrice.High, lastPrice.Low, lastPrice.Volume, lastPrice.Turnover)
			} else {
				fmt.Printf("Batch %d completed successfully. No data added.\n", batchNumber)
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

	return nil
}

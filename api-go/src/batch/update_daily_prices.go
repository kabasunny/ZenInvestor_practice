// api-go\src\batch\update_daily_prices.go
package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"api-go/src/util" // 追加
	"context"
	"fmt"
	"sync"
	"time"
	// その他必要なインポート
)

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
		symbol := stock.Ticker + ".T" // 現状の株価取得は、日本株でyfainanceを想定しているため、末尾に.Tが必要。J-QUANTSに変更を検討
		symbols = append(symbols, symbol)
	}

	// 共通のチャンク分割関数を使用
	chunks := util.ChunkSymbols(symbols, 100)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error

	for i, chunk := range chunks {
		wg.Add(1)
		go func(chunk []string, batchNumber int) {
			defer wg.Done()

			req := &gsdwd.GetStocksDatalistWithDatesRequest{
				Symbols:   chunk,
				StartDate: startDate,
				EndDate:   endDate,
			}

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

			if err := jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("failed to add daily price data: %w", err)
				mu.Unlock()
				return
			}

			fmt.Printf("Batch %d completed successfully. Last added data: Ticker: %s, Date: %s\n",
				batchNumber, newDailyPrices[len(newDailyPrices)-1].Ticker, newDailyPrices[len(newDailyPrices)-1].Date)
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

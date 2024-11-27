// api-go\src\batch\daily_prices_storage.go

package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"api-go/src/util"
	"fmt"
	"sync"
	"time"
)

// processChunks はデータをチャンクに分割し、それぞれのチャンクをGoルーチンで並行に処理してデータベースに格納
func processChunks(stockPrices []*gsdwd.StockPrice, batchSize int, jdpRepo repository.JpDailyPriceRepository, mu *sync.Mutex, wg *sync.WaitGroup, overallErr *error) {
	// データをチャンクに分割
	chunks := util.ChunkData(stockPrices, batchSize)
	fmt.Printf("データ格納バッチ数: %d\n", len(chunks))

	for i, chunk := range chunks {
		wg.Add(1)
		go func(chunk []*gsdwd.StockPrice, batchNumber int) {
			defer wg.Done()

			startTimeProcessing := time.Now()
			var newDailyPrices []model.JpDailyPrice
			for _, data := range chunk {
				date, err := time.Parse("2006-01-02", data.Date)
				if err != nil {
					mu.Lock()
					*overallErr = fmt.Errorf("failed to parse date: %w", err)
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
			fmt.Printf("バッチ %d のデータ処理の処理時間: %s\n", batchNumber, endTimeProcessing.Sub(startTimeProcessing))

			if err := jdpRepo.AddDailyPriceData(&newDailyPrices); err != nil {
				mu.Lock()
				*overallErr = fmt.Errorf("failed to add daily price data: %w", err)
				mu.Unlock()
				return
			}

			if len(newDailyPrices) > 0 {
				fmt.Printf("Batch %d completed successfully. Last added data: Symbol: %s, Date: %s\n",
					batchNumber, newDailyPrices[len(newDailyPrices)-1].Symbol, newDailyPrices[len(newDailyPrices)-1].Date)
			} else {
				fmt.Printf("Batch %d completed with no data.\n", batchNumber)
			}
		}(chunk, i+1)
	}
}

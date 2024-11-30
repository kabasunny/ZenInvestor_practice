// api-go\src\batch\update_daily_prices_3.go
package batch

import (
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"api-go/src/util"
	"context"
	"fmt"
	"sync"
	"time"
	// その他必要なインポート
)

func UpdateDailyPrices_3(ctx context.Context,
	udsRepo repository.UpdateStatusRepository,
	jsiRepo repository.JpStockInfoRepository,
	jdpRepo repository.JpDailyPriceRepository,
	gsdwdClient client.GetStocksDatalistWithDatesClient,
	//clients map[string]interface{},
	startDate string, // 最新データ日付
	// lookbackDays []string, // 何日前までさかのぼってデータを取得するか
	batchSize int, // DB格納時のGoルーチン毎のデータ数
	symbolChunkSize int, // 株価取得時のリクエスト毎のデータ数
) error {
	// 関数全体の処理開始時刻
	startTimeOverall := time.Now()
	// gsdwdClient, ok := clients["get_stocks_datalist_with_dates"].(client.GetStocksDatalistWithDatesClient)
	// if !ok {
	//  return fmt.Errorf("failed to get get_stocks_datalist_with_dates_client")
	// }

	// シンボル抽出の処理時間
	startTimeTicker := time.Now()
	symbols, err := jsiRepo.GetAllSymbols()
	if err != nil {
		return fmt.Errorf("failed to get all symbols: %w", err)
	}
	endTimeTicker := time.Now()
	fmt.Printf("シンボルを抽出の処理時間: %s\n", endTimeTicker.Sub(startTimeTicker))
	fmt.Printf("抽出したシンボルの数: %d\n", len(symbols))

	// // シンボルの数を500に制限 //test
	// if len(symbols) > 500 { //test
	// 	symbols = symbols[:500] //test
	// } //test
	// fmt.Printf("制限されたシンボルの数: %d\n", len(symbols)) //test

	// シンボルリストをチャンクに分割
	symbolChunks := util.ChunkSymbols(symbols, symbolChunkSize)
	fmt.Printf("シンボルのチャンク数: %d\n", len(symbolChunks))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error
	delay := 1 * time.Second // 遅延時間を1秒に設定

	// チャンクごとにデータを取得して処理する
	for i, chunk := range symbolChunks {
		// データ取得の処理時間
		startTimeDownload := time.Now()
		req := &gsdwd.GetStocksDatalistWithDatesRequest{
			Symbols:   chunk,
			StartDate: startDate,
			EndDate:   startDate,
		}
		wg.Add(1)
		go func(req *gsdwd.GetStocksDatalistWithDatesRequest, batchNumber int) {
			defer wg.Done()
			gsdwdResponse, err := gsdwdClient.GetStocksDatalist(ctx, req)
			endTimeDownload := time.Now()
			if err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("バッチ %d でシンボル %v のデータ取得に失敗: %w", batchNumber, chunk, err)
				mu.Unlock()
				return
			}

			fmt.Printf("バッチ %d のデータ取得の処理時間: %s\n", batchNumber, endTimeDownload.Sub(startTimeDownload))

			// チャンクを処理する関数を呼び出し
			storeChunks(gsdwdResponse.StockPrices, batchSize, jdpRepo, &mu, &wg, &overallErr)
		}(req, i+1)

		// リクエスト間の遅延を設ける (例: 1秒)
		time.Sleep(delay)
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
	fmt.Printf("UpdateDailyPrices_3 全体の処理時間: %s\n", endTimeOverall.Sub(startTimeOverall))

	return nil
}

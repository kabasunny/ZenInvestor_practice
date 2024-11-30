// api-go\src\batch\update_stock_info.go
package batch

import (
	"api-go/src/model"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	get_stock_info_jq "api-go/src/service/ms_gateway/get_stock_info_jq"
	"api-go/src/util" // 追加
	"context"
	"fmt"
	"sync"
	"time"
)

// UpdateStockInfo は銘柄情報を更新し、ステータスを更新
func UpdateStockInfo(ctx context.Context, udsRepo repository.UpdateStatusRepository, jsiRepo repository.JpStockInfoRepository, gsijClient client.GetStockInfoJqClient) error {
	startTime := time.Now() // 処理開始時刻の記録

	// stockInfoClient, ok := clients["get_stock_info_jq"].(client.GetStockInfoJqClient)
	// if !ok {
	// 	return fmt.Errorf("failed to get get_stock_info_jq_client")
	// }

	// 銘柄情報の取得
	stockInfoReq := &get_stock_info_jq.GetStockInfoJqRequest{}
	stockInfoRes, err := gsijClient.GetStockInfoJq(ctx, stockInfoReq)
	if err != nil {
		return fmt.Errorf("failed to get stock info: %w", err)
	}

	// 新しい銘柄情報をリストに追加
	var newStockInfos []model.JpStockInfo
	for _, data := range stockInfoRes.Stocks {
		si := model.JpStockInfo{
			Symbol:   data.Ticker,
			Name:     data.Name,
			Sector:   data.Sector,
			Industry: data.Industry,
			Date:     data.Date,
		}
		newStockInfos = append(newStockInfos, si)
	}

	// 既存のデータを削除
	if err := jsiRepo.DeleteAllStockInfo(); err != nil {
		return fmt.Errorf("failed to delete all stock info: %w", err)
	}
	fmt.Println("Deleted all existing stock info from DB")

	// 銘柄情報のティッカーのみを抽出して文字列のスライスに変換
	var symbols []string
	for _, stock := range newStockInfos {
		symbols = append(symbols, stock.Symbol)
	}

	// 銘柄情報を100銘柄ごとに分割
	chunks := util.ChunkSymbols(symbols, 100)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var overallErr error

	for i, chunk := range chunks {
		wg.Add(1)
		go func(chunk []string, batchNumber int) {
			defer wg.Done()

			// チャンク内の各ティッカーに対応する銘柄情報を抽出
			var chunkStockInfos []model.JpStockInfo
			for _, ticker := range chunk {
				for _, stock := range newStockInfos {
					if stock.Symbol == ticker {
						chunkStockInfos = append(chunkStockInfos, stock)
						break
					}
				}
			}

			// データベースに挿入
			if err := jsiRepo.InsertStockInfo(&chunkStockInfos); err != nil {
				mu.Lock()
				overallErr = fmt.Errorf("batch %d failed to insert stock info: %w", batchNumber, err)
				mu.Unlock()
				return
			}

			// バッチの完了メッセージを表示
			mu.Lock()
			fmt.Printf("Batch %d completed successfully.\n", batchNumber)
			mu.Unlock()
		}(chunk, i+1)
	}

	wg.Wait()

	if overallErr != nil {
		return overallErr
	}

	if err := udsRepo.UpdateStatus("jp_stocks_info"); err != nil {
		return fmt.Errorf("failed to update status for jp_stocks_info: %w", err)
	}

	endTime := time.Now()                 // 処理終了時刻の記録
	elapsedTime := endTime.Sub(startTime) // 処理時間の計算
	fmt.Printf("UpdateStockInfo completed in %s\n", elapsedTime)

	return nil
}

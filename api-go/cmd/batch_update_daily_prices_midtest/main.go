// api-go\cmd\batch_update_daily_prices_midtest\main.go
package main

import (
	"api-go/src/batch"
	"api-go/src/infra"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	"context"
	"log"
	"time"
	// その他必要なインポート
)

func main() {
	infra.Initialize() // 初期化処理
	db := infra.SetupDB()
	udsRepo := repository.NewUpdateStatusRepository(db)
	jsiRepo := repository.NewJpStockInfoRepository(db)
	jdpRepo := repository.NewJpDailyPriceRepository(db)

	ctx := context.Background()
	clients := make(map[string]interface{})

	// client.NewGetStocksDatalistWithDatesClientから2つの値を受け取る
	gsdwdClient, err := client.NewGetStocksDatalistWithDatesClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	// clientsの初期化
	clients["get_stocks_datalist_with_dates"] = gsdwdClient

	// 現在の日付を取得し、12週間前の日付を計算
	now := time.Now()
	startDate := now.AddDate(0, 0, -86).Format("2006-01-02") // 12週間 + 2日前の日付

	// コード内でバッチサイズとシンボルチャンクサイズを指定
	batchSize := 50        // DB格納時のGoルーチン毎のデータ数
	symbolChunkSize := 250 // 株価取得時のリクエスト毎のデータ数
	log.Printf("Batch size: %d, Symbol chunk size: %d\n", batchSize, symbolChunkSize)

	err = batch.MidTestUpdateDailyPrices(ctx, udsRepo, jsiRepo, jdpRepo, clients, startDate, batchSize, symbolChunkSize)
	if err != nil {
		log.Fatalf("Failed to update daily prices: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_daily_prices_midtest/main.go

// 全体の処理時間: 55.8217402s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / batchSize := 50 symbolChunkSize := 250

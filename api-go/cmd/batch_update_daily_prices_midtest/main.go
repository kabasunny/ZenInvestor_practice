// api-go\cmd\batch_update_daily_prices_midtest\main.go
package main

import (
	"api-go/src/batch"
	"api-go/src/infra"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	"context"
	"log"
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

	// 日付の設定
	startDate := "2024-11-19"
	endDate := "2024-11-22"

	// バッチ処理の呼び出し
	err = batch.MidTestUpdateDailyPrices(ctx, udsRepo, jsiRepo, jdpRepo, clients, startDate, endDate)
	if err != nil {
		log.Fatalf("Failed to update daily prices: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_daily_prices_midtest/main.go

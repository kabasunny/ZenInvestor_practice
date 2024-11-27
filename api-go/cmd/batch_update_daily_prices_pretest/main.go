// api-go\cmd\batch_update_daily_prices_pretest\main.go
package main

import (
	"api-go/src/batch"
	"api-go/src/infra"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	"context"
	"fmt"
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

	// 指定された銘柄
	symbols := []string{"130A", "6932"} // 適当な2銘柄   1311 131A で実験中

	// デバッグ: 指定された日付範囲を確認
	startDate := "2024-08-19" // J-Quantsは12週間前の日付
	endDate := "2024-08-26"   // J-Quantsは12週間前の日付
	fmt.Printf("指定された日付範囲: 開始日: %s, 終了日: %s\n", startDate, endDate)

	// バッチ処理の呼び出し
	err = batch.PreTestUpdateDailyPrices(ctx, udsRepo, jsiRepo, jdpRepo, clients, symbols, startDate, endDate)
	if err != nil {
		log.Fatalf("Failed to update daily prices: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_daily_prices_pretest/main.go

// api-go\cmd\batch_update_daily_prices\batch_update_daily_prices.go
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
	infra.Initialize()    // 主に環境変数の初期化処理
	db := infra.SetupDB() // DBのセットアップ
	udsRepo := repository.NewUpdateStatusRepository(db)
	jsiRepo := repository.NewJpStockInfoRepository(db)
	jdpRepo := repository.NewJpDailyPriceRepository(db)

	ctx := context.Background()
	msClients := make(map[string]interface{})

	// gRPC クライアント -------------------------------------------------------------------

	// GetStocksDatalistWithDatesClient初期化
	gsdwdClient, err := client.NewGetStocksDatalistWithDatesClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create get stocks datalist with dates client: %v", err)
	}
	msClients["get_stocks_datalist_with_dates"] = gsdwdClient // そのままインスタンスをUpdateDailyPricesの引数で渡してもよいが…

	// GetTradingCalendarJqClient初期化
	gtcjClient, err := client.NewGetTradingCalendarJqClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create get stock info from jq client: %v", err)
	}
	msClients["get_trading_calendar_jq"] = gtcjClient // mapに追加

	// ----------------------------------------------------------------------------------

	// J-QUANTS API フリープラン用の日付設定
	now := time.Now()
	startDate := now.AddDate(0, 0, -86).Format("2006-01-02") // 12週間 + 2日前の日付

	// コード内でバッチサイズとシンボルチャンクサイズを指定
	batchSize := 200 // DB格納時のGoルーチン毎のデータ数
	dataSize := 1000 // 株価取得時のリクエスト毎のデータ数

	// 日数計算ロジック
	days := 5                                                                       // 何日前までさかのぼってデータが必要か設定値
	lookbackDays, err := batch.CalculateLookbackDate(ctx, jdpRepo, startDate, days) // 実際にさかのぼってデータを取得する日数
	if err != nil {
		log.Fatalf("Failed to calculate lookback start date: %v", err)
	}

	symbolChunkSize := dataSize / lookbackDays // さかのぼる日数によって、チャンクサイズを小さくすることで、リクエスト負荷を均一に保つ

	log.Printf("Batch size: %d, Symbol chunk size: %d\n, Look back days: %d\n", batchSize, symbolChunkSize, lookbackDays)

	err = batch.UpdateDailyPrices(ctx, udsRepo, jsiRepo, jdpRepo, msClients, startDate, batchSize, symbolChunkSize, lookbackDays)
	if err != nil {
		log.Fatalf("Failed to update daily prices: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_daily_prices/batch_update_daily_prices.go

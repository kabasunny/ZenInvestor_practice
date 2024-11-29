// api-go\cmd\batch_update_daily_prices\batch_update_daily_prices.go
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
	infra.Initialize()          // 主に環境変数の初期化処理
	db := infra.SetupDB()       // DBのセットアップ
	ctx := context.Background() // コンテキスト

	// 各DBアクセス用のリポジトリのインスタンス
	udsRepo := repository.NewUpdateStatusRepository(db)
	jsiRepo := repository.NewJpStockInfoRepository(db)
	jdpRepo := repository.NewJpDailyPriceRepository(db)

	// gRPCクライアント -------------------------------------------------------------------

	// MSクライアント格納用変数
	// msClients := make(map[string]interface{})

	// 株価データを取得する GetStocksDatalistWithDatesClient初期化
	gsdwdClient, err := client.NewGetStocksDatalistWithDatesClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create get stocks datalist with dates client: %v", err) // サーバーを立ち上げ忘れるので、今は入れておく
	}
	// msClients["get_stocks_datalist_with_dates"] = gsdwdClient // そのままインスタンスをUpdateDailyPricesの引数で渡してもよいが…

	// 休業日を確認する GetTradingCalendarJqClient初期化
	gtcjClient, err := client.NewGetTradingCalendarJqClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create get stock info from jq client: %v", err)
	}
	// msClients["get_trading_calendar_jq"] = gtcjClient // mapに追加

	// ----------------------------------------------------------------------------------

	// J-QUANTS API フリープラン用の日付設定
	// now := time.Now()
	// startDate := now.AddDate(0, 0, -86).Format("2006-01-02") // 12週間 + 余裕2日前の日付

	//テスト用 //test
	startDate := "2023-11-01" //test

	// コード内でバッチサイズとシンボルチャンクサイズを指定
	batchSize := 200 // DB格納時のGoルーチン毎のデータ数
	dataSize := 1000 // 株価取得時のリクエスト毎のデータ数

	// 日数計算ロジック
	days := 5 // 何日前までさかのぼってデータが必要か設定値

	lookbackDays, err := batch.CalculateLookbackDate(ctx, jdpRepo, startDate, days, gtcjClient) // 実際にさかのぼってデータを取得する日数
	if err != nil {
		log.Fatalf("Failed to calculate lookback start date: %v", err)
	}

	if lookbackDays == nil {
		fmt.Printf("lookbackDays==nil : 最新データ取得済み")
		return
	}

	symbolChunkSize := dataSize / len(lookbackDays) // さかのぼる日数によって、チャンクサイズを小さくすることで、リクエスト負荷を均一に保つ

	log.Printf("Look back days: %s\n, Batch size: %d, Symbol chunk size: %d\n", lookbackDays, batchSize, symbolChunkSize)

	err = batch.UpdateDailyPrices_3(ctx, udsRepo, jsiRepo, jdpRepo, gsdwdClient, startDate, lookbackDays, batchSize, symbolChunkSize)
	if err != nil {
		log.Fatalf("Failed to update daily prices: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_daily_prices/batch_update_daily_prices.go

// UpdateDailyPrices 全体の処理時間: 57.6495232s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / test 500 銘柄
// UpdateDailyPrices_2 全体の処理時間: 1m4.0582447s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / test 500 銘柄
// UpdateDailyPrices_3 全体の処理時間: 58.7952742s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / test 500 銘柄

// 全体の処理時間:  : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / test 500 銘柄

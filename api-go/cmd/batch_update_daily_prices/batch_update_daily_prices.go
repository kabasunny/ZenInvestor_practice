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
	"time"
	// その他必要なインポート
)

func main() {
	startTimeOverall := time.Now()

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

	// // J-QUANTS API フリープラン用の日付設定 特に指定せずとも大丈夫のようだ
	// now := time.Now()
	// startDate := now.AddDate(0, 0, -86).Format("2006-01-02") // 12週間 + 余裕2日前の日付

	//テスト用 //test
	// startDate := "2023-11-01" //test

	// コード内でバッチサイズとシンボルチャンクサイズを指定
	batchSize := 100 // DB格納時のGoルーチン毎のデータ数
	dataSize := 500  // 株価取得時のリクエスト毎のデータ数

	// 日数計算ロジック
	days := 5 // 何日前までさかのぼってデータが必要か設定値

	lookbackDays, err := batch.CalculateLookbackDate(ctx, jdpRepo, jsiRepo, days, gtcjClient) // 実際にさかのぼってデータを取得する日数
	if err != nil {
		log.Fatalf("Failed to calculate lookback start date: %v", err)
	}

	if lookbackDays == nil {
		fmt.Printf("lookbackDays==nil : 最新データ取得済み")
		return
	}

	// symbolChunkSize := dataSize / len(lookbackDays) // さかのぼる日数によって、チャンクサイズを小さくすることで、リクエスト負荷を均一に保つ 今となっては意味ないよね

	log.Printf("Look back days: %s\n, Batch size: %d, Symbol chunk size: %d\n", lookbackDays, batchSize, dataSize)

	// lookbackDays の日付をループで取り出して処理
	for _, fetchDate := range lookbackDays {
		fmt.Printf("取得日: %s\n", fetchDate) // 取得日を表示

		err = batch.UpdateDailyPrices_3(ctx, udsRepo, jsiRepo, jdpRepo, gsdwdClient, fetchDate, batchSize, dataSize)
		if err != nil {
			log.Fatalf("Failed to update daily prices: %v", err)
		}
	}

	// lookbackDays[0] より5日前の日付を計算
	beforeDate, err := time.Parse("2006-01-02", lookbackDays[0])
	if err != nil {
		log.Fatalf("Failed to parse lookbackDays[0]: %v", err)
	}
	beforeDate = beforeDate.AddDate(0, 0, -5)

	// lookbackDays[0]より5日以上前のデータを消去する
	jdpRepo.DeleteBeforeSpecifiedDate(beforeDate.Format("2006-01-02"))

	// 関数全体の処理終了時刻
	endTimeOverall := time.Now()
	fmt.Printf("株価取得バッチの処理時間: %s\n", endTimeOverall.Sub(startTimeOverall))
}

// 実行コマンド
// go run ./cmd/batch_update_daily_prices/batch_update_daily_prices.go

// UpdateDailyPrices   : 非同期処理のリクエスト
// UpdateDailyPrices_2 : シリアル処理のリクエスト
// UpdateDailyPrices_3 : 非同期処理の遅延リクエスト　これかな

// UpdateDailyPrices 全体の処理時間: 57.6495232s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / test 500 銘柄 startDate := "2023-11-01"
// UpdateDailyPrices_2 全体の処理時間: 1m48.8636617s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / test 500 銘柄 startDate := "2023-11-01"
// UpdateDailyPrices_3 全体の処理時間: 58.7952742s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / test 500 銘柄 startDate := "2023-11-01"

// UpdateDailyPrices_3 全体の処理時間: 1m1.665574s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / test 500 銘柄 startDate := "2023-11-01"
// 株価取得バッチの処理時間: 5m14.2006247s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / test 500 銘柄 startDate := "2023-11-01"

// UpdateDailyPrices_3 全体の処理時間: 2m23.8292159s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / startDate := "2023-11-01"
// 株価取得バッチの処理時間: 12m4.9949104s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / startDate := "2023-11-01"

// UpdateDailyPrices_3 全体の処理時間: 2m26.9501114s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / 本番
// 株価取得バッチの処理時間: 12m27.3960013s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz / 本番

// UpdateDailyPrices_3 全体の処理時間: 7m20.3971786s : 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz / batchSize := 200, dataSize := 800, delay := 5 * time.Second
// 株価取得バッチの処理時間: 21m32.3409058s : 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz / batchSize := 200, dataSize := 800, delay := 5 * time.Second

// UpdateDailyPrices_3 全体の処理時間: 3m9.4279275s : 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz / batchSize := 100, dataSize := 400, delay := 10 * time.Second
// 株価取得バッチの処理時間: 15m39.2943609s : 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz / batchSize := 100, dataSize := 400, delay := 10 * time.Second

// UpdateDailyPrices_3 全体の処理時間: 3m17.9234813s : 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz / batchSize := 100, dataSize := 500, delay := 15 * time.Second
// 株価取得バッチの処理時間: 16m35.0355377s : 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz / batchSize := 100, dataSize := 400, delay := 10 * time.Second

// UpdateDailyPrices_3 全体の処理時間: 3m22.0738907s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / batchSize := 100, dataSize := 500, delay := 15 * time.Second
// 株価取得バッチの処理時間: 17m8.9543363s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz / batchSize := 100, dataSize := 500, delay := 15 * time.Second

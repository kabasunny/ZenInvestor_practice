// api-go\test\service\ranking_service_test.go

package service_test

// import (
// 	"context"
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"api-go/src/repository"
// 	"api-go/src/service"
// 	"api-go/src/service/ms_gateway/client"
// 	"api-go/test/repository/repository_test_helper"

// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetRankingDataIntegration(t *testing.T) {
// 	// 1. gRPCサーバーを起動 (別プロセスで)
// 	fmt.Println("Step 1: Starting gRPC server...")

// 	err := godotenv.Load("../../.env") //テストではパスを指定しないとうまく読み取らない
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	// 2. クライアントをセットアップ
// 	fmt.Println("Step 2: Setting up clients...")
// 	ctx, cancel := context.WithTimeout(context.Background(), 1800*time.Second) // タイムアウトを設定
// 	defer cancel()

// 	msClients := make(map[string]interface{})
// 	fmt.Println("Clients setup...")

// 	getStockInfoJqClient, err := client.NewGetStockInfoJqClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create get stock info jq client: %v", err)
// 	}
// 	msClients["get_stock_info_jq"] = getStockInfoJqClient

// 	getStocksDatalistWithDatesClient, err := client.NewGetStocksDatalistWithDatesClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create get stocks datalist with dates client: %v", err)
// 	}
// 	msClients["get_stocks_datalist_with_dates"] = getStocksDatalistWithDatesClient

// 	// 3. データベースのセットアップ
// 	fmt.Println("Step 3: Setting up the database...")
// 	db := repository_test_helper.SetupTestDB()
// 	if db == nil {
// 		log.Fatalf("Failed to set up the database")
// 	}
// 	// repository_test_helper.InitializeUpdateStatusTable(db) // 追加: 初期データの投入
// 	repository_test_helper.PrintUpdateStatusTable(db) // 追加: デバッグ用のテーブル内容表示

// 	// 4. サービスの初期化
// 	fmt.Println("Step 4: Initializing RankingService...")
// 	udsRepo := repository.NewUpdateStatusRepository(db)
// 	jsiRepo := repository.NewJpStockInfoRepository(db)
// 	jdpRepo := repository.NewJpDailyPriceRepository(db)
// 	j5mrRepo := repository.NewJp5dMvaRankingRepository(db)

// 	rankingService := service.NewRankingService(udsRepo, jsiRepo, jdpRepo, j5mrRepo, msClients)

// 	// 5. サービスの呼び出し
// 	fmt.Println("Step 5: Calling GetRankingData service...")
// 	res, err := rankingService.GetRankingData(ctx)
// 	if err != nil {
// 		fmt.Printf("Error calling GetRankingData service: %v\n", err)
// 	}
// 	fmt.Println("Service call completed.")

// 	// 6. アサーション
// 	fmt.Println("Step 6: Performing assertions...")
// 	assert.NoError(t, err)
// 	assert.NotNil(t, res)
// 	assert.NotEmpty(t, *res) // レスポンスにデータが含まれていることを検証
// 	fmt.Println("Assertions completed.")

// 	// 7. 結果をファイルに保存
// 	fmt.Println("Step 7: Saving results to CSV file...")
// 	outputDir := os.Getenv("TEST_SERVICE_OUTPUT_DIR")
// 	if outputDir == "" {
// 		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
// 	}

// 	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
// 	filename := fmt.Sprintf("get_ranking_data_service_test_%s.csv", timestamp)
// 	outputFile := filepath.Join(outputDir, filename)

// 	if err := os.MkdirAll(outputDir, 0755); err != nil {
// 		t.Errorf("failed to create output directory: %v", err)
// 		return // エラーが発生してもテストを継続
// 	}

// 	file, err := os.Create(outputFile)
// 	if err != nil {
// 		t.Errorf("failed to create output file: %v", err)
// 		return // エラーが発生してもテストを継続
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// ヘッダーを書き込む
// 	headers := []string{"Ranking", "Symbol", "Date", "AvgTurnover", "Name", "LatestClose"}
// 	if err := writer.Write(headers); err != nil {
// 		t.Errorf("failed to write headers to CSV file: %v", err)
// 		return
// 	}

// 	// データを書き込む
// 	for _, data := range *res {
// 		record := []string{
// 			fmt.Sprintf("%d", data.Ranking),
// 			data.Ticker,
// 			data.Date,
// 			fmt.Sprintf("%.2f", data.AvgTurnover),
// 			data.Name,
// 			fmt.Sprintf("%.2f", data.LatestClose),
// 		}
// 		if err := writer.Write(record); err != nil {
// 			t.Errorf("failed to write record to CSV file: %v", err)
// 		}
// 	}
// 	fmt.Println("Results written to CSV file successfully.")
// }

// // テストの実行コード
// // go test -v ./test/service/ranking_service_test.go -run TestGetRankingDataIntegration

// // 全マクロサービスを立ち上げるなら、/ZenInvestor_practiceにて
// // ./StartMicroservices.ps1
// // 本テスト用マイクロサービスを立ち上げるなら、/ZenInvestor_practiceにて
// // ./StartForRankingMicroservices.ps1

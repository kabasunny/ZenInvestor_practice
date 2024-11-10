package service_test

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"api-go/src/service"
// 	indicator "api-go/src/service/ms_gateway/calculate_indicator"
// 	"api-go/src/service/ms_gateway/client"

// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetStockDataWithSMA(t *testing.T) {
// 	// 1. gRPCサーバーを起動 (別プロセスで)
// 	fmt.Println("Step 1: Starting gRPC server...")

// 	godotenv.Load("../../.env") // テストではパスを指定しないとうまく読み取らない

// 	// 2. クライアントをセットアップ
// 	fmt.Println("Step 2: Setting up clients...")
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // タイムアウトを設定
// 	defer cancel()
// 	fmt.Println("ctx setup successfully.")

// 	// // 他のサービス実装時にテストが行えないため、あえてインフラは使わない
// 	// msClients, err := infra.SetupMsClients(ctx) // マイクロサービス接続の初期化、エラー処理を追加
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to setup microservice clients: %v", err) // エラーが発生したらログ出力して終了
// 	// }
// 	// service := service.NewStockServiceImpl(msClients.MSClients) // msClients構造体の中のmap[string]interface{}

// 	msClients := make(map[string]interface{})
// 	fmt.Println("Clients setup...")

// 	getStockDataClient, err := client.NewGetStockDataClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create get stock data client: %v", err)
// 	}
// 	msClients["get_stock_data"] = getStockDataClient
// 	fmt.Println("getStockDataClient setup successfully.")

// 	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create simple moving average client: %v", err)
// 	}
// 	msClients["simple_moving_average"] = smaClient
// 	fmt.Println("smaClient setup successfully.")

// 	// StockServiceを初期化
// 	fmt.Println("Initializing StockService...")
// 	service := service.NewStockServiceImpl(msClients)
// 	fmt.Println("StockService initialized.")

// 	// 3. リクエストデータを作成
// 	fmt.Println("Step 3: Creating request data for ticker...")
// 	ticker := "AAPL" // テスト用のティッカーシンボル
// 	period := "5d"   // テスト用の期間
// 	indicators := []*indicator.IndicatorParams{
// 		{
// 			Type: "SMA",
// 			Params: map[string]string{
// 				"window_size": "14", // テスト用のウィンドウサイズ
// 			},
// 		},
// 	}

// 	// 4. サービスの呼び出し
// 	fmt.Println("Step 4: Calling GetStockData service...")
// 	res, err := service.GetStockData(ctx, ticker, period, indicators)
// 	if err != nil {
// 		fmt.Printf("Error calling GetStockData service: %v\n", err)
// 	}
// 	fmt.Println("Service call completed.")

// 	// 5. アサーション
// 	fmt.Println("Step 5: Performing assertions...")
// 	assert.NoError(t, err)
// 	assert.NotNil(t, res)
// 	assert.NotEmpty(t, res.StockData) // レスポンスにデータが含まれていることを検証
// 	fmt.Println("Assertions completed.")

// 	// 6. 結果をファイルに保存
// 	fmt.Println("Step 6: Saving results to file in directory...")
// 	outputDir := os.Getenv("TEST_SERVICE_OUTPUT_DIR")
// 	if outputDir == "" {
// 		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
// 	}

// 	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
// 	filename := fmt.Sprintf("get_stock_data_with_sma_service_test_%s.txt", timestamp)
// 	outputFile := filepath.Join(outputDir, filename)

// 	if err := os.MkdirAll(outputDir, 0755); err != nil {
// 		t.Errorf("failed to create output directory: %v", err) // t.Errorに変更
// 		return                                                 // エラーが発生してもテストを継続
// 	}

// 	file, err := os.Create(outputFile)
// 	if err != nil {
// 		t.Errorf("failed to create output file: %v", err) // t.Errorに変更
// 		return                                            // エラーが発生してもテストを継続
// 	}
// 	defer file.Close()
// 	fmt.Println("File created successfully.")

// 	for key, value := range res.StockData {
// 		data := fmt.Sprintf("%s: open: %.2f, close: %.2f, high: %.2f, low: %.2f, volume: %.2f\n", key, value.Open, value.Close, value.High, value.Low, value.Volume)
// 		if _, err := file.WriteString(data); err != nil {
// 			t.Errorf("failed to write to output file: %v", err) // t.Errorに変更
// 			// return // エラーが発生してもテストを継続 (必要に応じて)
// 		}
// 	}
// 	fmt.Println("Results written to file successfully.")

// 	// テストの実行コード
// 	// go test -v ./test/service/simple_moving_average_service_test.go -run TestGetStockDataWithSMA

// 	// 7. gRPCサーバーを停止
// 	fmt.Println("Step 7: Stopping gRPC server...")
// }

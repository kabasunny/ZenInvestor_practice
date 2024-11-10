package service_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"api-go/src/service"
	indicator "api-go/src/service/ms_gateway/calculate_indicator"
	"api-go/src/service/ms_gateway/client"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetStockDataWithIndicators(t *testing.T) {
	// 1. gRPCサーバーを起動 (別プロセスで)
	fmt.Println("Step 1: Starting gRPC server...")

	godotenv.Load("../../.env") // テストではパスを指定しないとうまく読み取らない

	// 2. クライアントをセットアップ
	fmt.Println("Step 2: Setting up clients...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // タイムアウトを設定
	defer cancel()
	fmt.Println("ctx setup successfully.")

	msClients := make(map[string]interface{})
	fmt.Println("Clients setup...")

	getStockDataClient, err := client.NewGetStockDataClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create get stock data client: %v", err)
	}
	msClients["get_stock_data"] = getStockDataClient
	fmt.Println("getStockDataClient setup successfully.")

	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create simple moving average client: %v", err)
	}
	msClients["simple_moving_average"] = smaClient
	fmt.Println("smaClient setup successfully.")

	generateChartClient, err := client.NewGenerateChartClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create generate chart client: %v", err)
	}
	msClients["generate_chart"] = generateChartClient
	fmt.Println("generateChartClient setup successfully.")

	// StockServiceを初期化
	fmt.Println("Initializing StockService...")
	service := service.NewStockServiceImpl(msClients)
	fmt.Println("StockService initialized.")

	// 3. リクエストデータを作成
	fmt.Println("Step 3: Creating request data for ticker...")
	ticker := "AAPL" // テスト用のティッカーシンボル
	period := "1y"   // テスト用の期間
	indicators := []*indicator.IndicatorParams{
		{
			Type: "SMA",
			Params: map[string]string{
				"window_size": "30", // テスト用のウィンドウサイズ
			},
		},
	}

	// 4. サービスの呼び出し
	fmt.Println("Step 4: Calling GetStockData service...")
	res, err := service.GetStockData(ctx, ticker, period, indicators)
	if err != nil {
		fmt.Printf("Error calling GetStockData service: %v\n", err)
	}
	fmt.Println("Service call completed.")

	// 5. アサーション
	fmt.Println("Step 5: Performing assertions...")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.ChartData) // 変更点: ChartDataを確認
	fmt.Println("Assertions completed.")

	// 6. 結果をファイルに保存 (ChartDataをPNGファイルとして保存)
	fmt.Println("Step 6: Saving chart image to file...")
	outputDir := os.Getenv("TEST_SERVICE_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs"
	}

	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("get_stock_data_with_indicators_service_test_%s.png", timestamp) // PNG拡張子に変更
	outputFile := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Errorf("failed to create output directory: %v", err)
		return
	}

	// Base64デコード
	decodedChartData, err := base64.StdEncoding.DecodeString(res.ChartData)
	if err != nil {
		t.Errorf("failed to decode chart data: %v", err)
		return
	}

	// ファイルに書き込み
	if err := os.WriteFile(outputFile, decodedChartData, 0644); err != nil {
		t.Errorf("failed to write chart image to file: %v", err)
		return
	}

	fmt.Printf("Chart image saved to: %s\n", outputFile)

	// テストの実行コード
	// go test -v ./test/service/generate_chart_service_test.go -run TestGetStockDataWithIndicators

	// 7. gRPCサーバーを停止
	fmt.Println("Step 7: Stopping gRPC server...")
}

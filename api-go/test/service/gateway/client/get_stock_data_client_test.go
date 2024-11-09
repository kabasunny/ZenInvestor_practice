package client_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"api-go/src/service/ms_gateway/client"
	get_stock_data "api-go/src/service/ms_gateway/get_stock_data"
	"api-go/test/service/gateway/client_test_helper" // ヘルパー関数のパッケージをインポート

	"github.com/stretchr/testify/assert"
)

func TestNewStockClient(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background()
	stockClient, err := client.NewGetStockDataClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stockClient)
	defer stockClient.Close() // 接続を閉じる
}

func TestGetStockData(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す
	ctx := context.Background()
	stockClient, err := client.NewGetStockDataClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stockClient)
	defer stockClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // 初回タイムアウトで失敗するため15秒
	defer cancel()

	req := &get_stock_data.GetStockDataRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト
	res, err := stockClient.GetStockData(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.StockData)

	// ファイル出力 (オプション)
	outputDir := os.Getenv("TEST_CLIENT_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
	}

	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
	filename := fmt.Sprintf("get_stock_data_client_test%s.txt", timestamp)
	outputFile := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Errorf("failed to create output directory: %v", err) // t.Errorに変更
		return                                                 // エラーが発生してもテストを継続
	}

	file, err := os.Create(outputFile)
	if err != nil {
		t.Errorf("failed to create output file: %v", err) // t.Errorに変更
		return                                            // エラーが発生してもテストを継続
	}
	defer file.Close()
	fmt.Println("File created successfully.")

	for key, value := range res.StockData {
		data := fmt.Sprintf("%s: open: %.2f, close: %.2f, high: %.2f, low: %.2f, volume: %.2f\n", key, value.Open, value.Close, value.High, value.Low, value.Volume)
		_, err := file.WriteString(data)
		if err != nil {
			t.Fatalf("Failed to write to file: %v", err)
		}
	}
}

// go test -v ./test/service/gateway/client/get_stock_data_client_test.go

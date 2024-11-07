package client_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"api-go/src/service/ms_gateway/client"
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
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

	req := &ms_gateway.GetStockDataRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト
	res, err := stockClient.GetStockData(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.StockData)

	// ファイル出力 (オプション)
	outputDir := os.Getenv("TEST_OUTPUT_DIR") // 出力ディレクトリを環境変数から取得
	if outputDir != "" {
		filename := "get_stock_data_client_test.txt"
		outputFile := filepath.Join(outputDir, filename) // パスを結合

		file, err := os.Create(outputFile)
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		for key, value := range res.StockData {
			data := fmt.Sprintf("%s: open: %.2f, close: %.2f, high: %.2f, low: %.2f, volume: %.2f\n", key, value.Open, value.Close, value.High, value.Low, value.Volume)
			_, err := file.WriteString(data)
			if err != nil {
				t.Fatalf("Failed to write to file: %v", err)
			}
		}
	}
}

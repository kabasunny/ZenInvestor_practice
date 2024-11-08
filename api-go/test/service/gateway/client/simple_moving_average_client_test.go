package client_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	ms_gateway "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average"
	"api-go/src/service/ms_gateway/client"
	"api-go/test/service/gateway/client_test_helper" // ヘルパー関数のパッケージをインポート

	"github.com/stretchr/testify/assert"
)

func TestNewSimpleMovingAverageClient(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background() // context.Background() を使用
	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, smaClient)
	defer smaClient.Close()
}

func TestCalculateSimpleMovingAverage(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background()
	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, smaClient)
	defer smaClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ms_gateway.SimpleMovingAverageRequest{
		StockData:  []float32{10, 12, 14, 16, 18, 20, 10, 12, 14, 16, 18, 20},
		WindowSize: 3,
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト

	res, err := smaClient.CalculateSimpleMovingAverage(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	expectedSMA := []float64{12, 14, 16, 18, 16, 14, 12, 14, 16, 18} // float64に変更
	assert.Equal(t, expectedSMA, res.GetValues())                    // GetValues() を使用

	// ファイル出力

	outputDir := os.Getenv("TEST_CLIENT_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
	}

	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
	filename := fmt.Sprintf("simple_moving_average_client_test%s.txt", timestamp)
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
}

// go test -v ./test/service/gateway/client/simple_moving_average_client_test.go

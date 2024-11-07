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
	outputDir := os.Getenv("TEST_OUTPUT_DIR")
	if outputDir != "" {
		filename := "simple_moving_average_client_test.txt"
		outputFile := filepath.Join(outputDir, filename)

		file, err := os.Create(outputFile)
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		for i, value := range res.GetValues() {
			data := fmt.Sprintf("SMA[%d]: %.2f\n", i, value)
			_, err := file.WriteString(data)
			if err != nil {
				t.Fatalf("Failed to write to file: %v", err)
			}
		}
	}
}

// go test -v ./test/service/gateway/client/simple_moving_average_client_test.go

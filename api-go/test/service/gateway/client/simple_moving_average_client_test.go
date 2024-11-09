package client_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	sma "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average"
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

	req := &sma.SimpleMovingAverageRequest{
		StockData: map[string]*sma.StockDataForSMA{
			"2023-11-04": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-05": {Open: 12, Close: 14, High: 16, Low: 18, Volume: 110},
			"2023-11-06": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-07": {Open: 12, Close: 14, High: 16, Low: 18, Volume: 110},
			"2023-11-08": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-09": {Open: 12, Close: 14, High: 16, Low: 18, Volume: 110},
			// その他のデータも追加
		},
		WindowSize: 3,
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト

	res, err := smaClient.CalculateSimpleMovingAverage(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	expectedSMA := map[string]float64{
		"2023-11-06": 12.666666666666666,
		"2023-11-07": 13.333333333333332,
		"2023-11-08": 12.666666666666666,
		"2023-11-09": 13.333333333333332,
	} // float64に変更
	assert.Equal(t, expectedSMA, res.MovingAverage) // フィールド名を修正

	// ファイル出力

	outputDir := os.Getenv("TEST_CLIENT_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
	}

	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
	filename := fmt.Sprintf("simple_moving_average_client_test_%s.txt", timestamp)
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

	for key, value := range res.MovingAverage {
		line := fmt.Sprintf("%s: %.2f\n", key, value)
		if _, err := file.WriteString(line); err != nil {
			t.Errorf("failed to write to output file: %v", err) // t.Errorに変更
		}
	}
	fmt.Println("File created successfully.")
}

// go test -v ./test/service/gateway/client/simple_moving_average_client_test.go

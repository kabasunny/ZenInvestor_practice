// api-go\test\service\gateway\client\generate_chart_lc_sim_client_test.go

package client_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"api-go/src/service/ms_gateway/client"
	generate_chart_lc_sim "api-go/src/service/ms_gateway/generate_chart_lc_sim"
	"api-go/test/service/gateway/client_test_helper" // ヘルパー関数のパッケージをインポート

	"github.com/stretchr/testify/assert"
)

func TestNewGenerateChartLCClient(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background()
	chartClient, err := client.NewGenerateChartLCClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, chartClient)
	defer chartClient.Close() // 接続を閉じる
}

func TestGenerateChartSuccess(t *testing.T) { // 関数名を修正
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す
	ctx := context.Background()
	chartClient, err := client.NewGenerateChartLCClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, chartClient)
	defer chartClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // 初回タイムアウトで失敗するため15秒
	defer cancel()

	// テストデータの作成
	dates := []string{
		"2023-01-01", "2023-01-02", "2023-01-03", "2023-01-04", "2023-01-05",
		"2023-01-06", "2023-01-07", "2023-01-08", "2023-01-09", "2023-01-10",
	}
	closePrices := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109}
	purchaseDate := "2023-01-03"
	purchasePrice := 102.0
	endDate := "2023-01-08"
	endPrice := 107.0

	req := &generate_chart_lc_sim.GenerateChartLCRequest{
		Dates:         dates,
		ClosePrices:   closePrices,
		PurchaseDate:  purchaseDate,
		PurchasePrice: purchasePrice,
		EndDate:       endDate,
		EndPrice:      endPrice,
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト
	res, err := chartClient.GenerateChart(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Success)
	assert.NotEmpty(t, res.ChartData)

	// ファイル出力 (オプション)
	outputDir := os.Getenv("TEST_CLIENT_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
	}

	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
	filename := fmt.Sprintf("generate_chart_lc_sim_client_test%s.png", timestamp)
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

	imageData, err := base64.StdEncoding.DecodeString(res.ChartData)
	if err != nil {
		t.Errorf("failed to decode chart data: %v", err) // t.Errorに変更
		return                                           // エラーが発生してもテストを継続
	}

	_, err = file.Write(imageData)
	if err != nil {
		t.Errorf("failed to write image data to file: %v", err) // t.Errorに変更
		return                                                  // エラーが発生してもテストを継続
	}

	fmt.Printf("Chart image has been saved to %s\n", outputFile)
}

// go test -v ./test/service/gateway/client/generate_chart_lc_sim_client_test.go

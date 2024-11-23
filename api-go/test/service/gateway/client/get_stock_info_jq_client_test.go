// api-go\test\service\gateway\client\get_stock_info_jq_client_test.go
package client_test

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"api-go/src/service/ms_gateway/client"
	get_stock_info_jq "api-go/src/service/ms_gateway/get_stock_info_jq"
	"api-go/test/service/gateway/client_test_helper" // ヘルパー関数のパッケージをインポート

	"github.com/stretchr/testify/assert"
)

func TestNewStockInfoJqClient(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background()
	stockClient, err := client.NewGetStockInfoJqClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stockClient)
	defer stockClient.Close() // 接続を閉じる
}

func TestGetStockInfoJq(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す
	ctx := context.Background()
	stockClient, err := client.NewGetStockInfoJqClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stockClient)
	defer stockClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // 初回タイムアウトで失敗するため15秒
	defer cancel()

	req := &get_stock_info_jq.GetStockInfoJqRequest{
		// J-QUANTSは日本株なので無し11/23　Country: "Japan",
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト
	res, err := stockClient.GetStockInfoJq(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Stocks)

	// ファイル出力 (オプション)
	outputDir := os.Getenv("TEST_CLIENT_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
	}

	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
	filename := fmt.Sprintf("get_stock_info_jq_client_test%s.csv", timestamp)
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

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// CSVのヘッダーを書き込む
	header := []string{"Ticker", "Name", "Sector", "Industry"}
	if err := writer.Write(header); err != nil {
		t.Fatalf("Failed to write header to file: %v", err)
	}

	// レスポンスの各株式情報を書き込む
	for _, stock := range res.Stocks {
		record := []string{stock.Ticker, stock.Name, stock.Sector, stock.Industry}
		if err := writer.Write(record); err != nil {
			t.Fatalf("Failed to write record to file: %v", err)
		}
	}
}

// go test -v ./test/service/gateway/client/get_stock_info_jq_client_test.go

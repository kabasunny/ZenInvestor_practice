// api-go\test\service\generate_chart_service_test.go
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

// setupClients は各マイクロサービスクライアントを初期化して返す
func setupClients(ctx context.Context) (map[string]interface{}, error) {
	msClients := make(map[string]interface{})

	getStockDataClient, err := client.NewGetStockDataClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create get stock data client: %w", err)
	}
	msClients["get_stock_data"] = getStockDataClient

	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create simple moving average client: %w", err)
	}
	msClients["simple_moving_average"] = smaClient

	generateChartClient, err := client.NewGenerateChartClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create generate chart client: %w", err)
	}
	msClients["generate_chart"] = generateChartClient

	return msClients, nil
}

// saveChartToFile はチャートデータをBase64デコードして指定されたファイルに保存する
func saveChartToFile(outputDir, filename, chartDataBase64 string) error {
	// Base64デコード
	decodedChartData, err := base64.StdEncoding.DecodeString(chartDataBase64)
	if err != nil {
		return fmt.Errorf("failed to decode chart data: %w", err)
	}

	// ファイルパスを生成
	outputFile := filepath.Join(outputDir, filename)

	// 出力ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// ファイルに書き込み
	if err := os.WriteFile(outputFile, decodedChartData, 0644); err != nil {
		return fmt.Errorf("failed to write chart image to file: %w", err)
	}

	fmt.Printf("Chart image saved to: %s\n", outputFile)
	return nil
}

func TestGetStockDataWithIndicators(t *testing.T) {
	// .envファイルをロードして環境変数を設定
	godotenv.Load("../../.env")

	// クライアントをセットアップ
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	msClients, err := setupClients(ctx)
	if err != nil {
		log.Fatalf("Failed to set up clients: %v", err)
	}

	// StockServiceを初期化
	stockService := service.NewStockServiceImpl(msClients)

	// テスト用のティッカーシンボルと期間を設定
	ticker := "AAPL"
	period := "1y"
	indicators := []*indicator.IndicatorParams{
		{
			Type: "SMA",
			Params: map[string]string{
				"window_size": "30",
			},
		},
	}

	// テストケースの定義
	tests := []struct {
		name          string
		includeVolume bool
		filename      string
	}{
		{
			name:          "WithVolume",
			includeVolume: true,
			filename:      "generate_chart_with_indicators_with_volume.png",
		},
		{
			name:          "WithoutVolume",
			includeVolume: false,
			filename:      "generate_chart_with_indicators_without_volume.png",
		},
	}

	// 各テストケースの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// サービスの呼び出し
			res, err := stockService.GetStockChart(ctx, ticker, period, indicators, tt.includeVolume)

			// アサーション
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res.ChartData)

			// 結果をファイルに保存
			outputDir := os.Getenv("TEST_SERVICE_OUTPUT_DIR")
			if outputDir == "" {
				outputDir = "api-go/test/test_outputs"
			}

			if err := saveChartToFile(outputDir, tt.filename, res.ChartData); err != nil {
				t.Errorf("failed to save chart to file: %v", err)
			}
		})
	}
}

// テストの実行コード
// go test -v ./test/service/generate_chart_service_test.go -run TestGetStockDataWithIndicators

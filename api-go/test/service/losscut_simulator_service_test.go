// api-go\test\service\losscut_simulator_service_test.go
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
	"api-go/src/service/ms_gateway/client"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupLossCutClients は各マイクロサービスクライアントを初期化して返す
func setupLossCutClients(ctx context.Context) (map[string]interface{}, error) {
	msClients := make(map[string]interface{})

	getStockDataClient, err := client.NewGetStockDataClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create get stock data client: %w", err)
	}
	msClients["get_stock_data"] = getStockDataClient

	generateChartClient, err := client.NewGenerateChartLCClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create generate chart client: %w", err)
	}
	msClients["generate_chart_lc_sim"] = generateChartClient

	return msClients, nil
}

// saveLossCutChartToFile はチャートデータをBase64デコードして指定されたファイルに保存する
func saveLossCutChartToFile(outputDir, filename, chartDataBase64 string) error {
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

func TestGetStockChartForLCSim(t *testing.T) {
	// .envファイルをロードして環境変数を設定
	godotenv.Load("../../.env")

	// クライアントをセットアップ
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	msClients, err := setupLossCutClients(ctx)
	if err != nil {
		log.Fatalf("Failed to set up clients: %v", err)
	}

	// LosscutSimulatorServiceを初期化
	lossCutSimulatorService := service.NewLosscutSimulatorServiceImpl(msClients)

	// テスト用のティッカーシンボルとシミュレーション日を設定
	ticker := "AAPL"
	simulationDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC) // データの範囲内にある日付
	stopLossPercentage := 10.0
	trailingStopTrigger := 5.0
	trailingStopUpdate := 3.0

	// サービスの呼び出し
	res, profitLoss, err := lossCutSimulatorService.GetStockChartForLCSim(ctx, ticker, simulationDate, stopLossPercentage, trailingStopTrigger, trailingStopUpdate)

	// アサーション
	if err != nil {
		t.Fatalf("Received unexpected error: %v", err)
	}
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.ChartData)
	assert.GreaterOrEqual(t, profitLoss, 0.0)

	// 結果をファイルに保存
	outputDir := os.Getenv("TEST_SERVICE_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs"
	}

	filename := "losscut_simulator_service_test.png"
	if err := saveLossCutChartToFile(outputDir, filename, res.ChartData); err != nil {
		t.Errorf("failed to save chart to file: %v", err)
	}
}

// テストの実行コード
// go test -v ./test/service/losscut_simulator_service_test.go -run TestGetStockChartForLCSim

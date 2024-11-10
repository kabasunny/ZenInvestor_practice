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
	gc "api-go/src/service/ms_gateway/generate_chart"
	"api-go/test/service/gateway/client_test_helper"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerateChartClient(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background()
	gcClient, err := client.NewGenerateChartClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, gcClient)
	defer gcClient.Close()
}

func TestGenerateChart(t *testing.T) {
	client_test_helper.LoadTestEnv() // ヘルパー関数を呼び出す

	ctx := context.Background()
	gcClient, err := client.NewGenerateChartClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, gcClient)
	defer gcClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &gc.GenerateChartRequest{
		StockData: map[string]*gc.StockDataForChart{
			"2023-11-04": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-05": {Open: 12, Close: 14, High: 16, Low: 18, Volume: 110},
			"2023-11-06": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-07": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-08": {Open: 12, Close: 14, High: 16, Low: 18, Volume: 110},
			"2023-11-09": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-10": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			"2023-11-11": {Open: 12, Close: 14, High: 16, Low: 18, Volume: 110},
			"2023-11-12": {Open: 10, Close: 12, High: 14, Low: 16, Volume: 100},
			// その他のデータも追加
		},
		Indicators: []*gc.IndicatorData{
			{Type: "SMA", Values: map[string]float64{"2023-11-04": 12.66, "2023-11-05": 13.33, "2023-11-06": 12.66, "2023-11-07": 13.33, "2023-11-08": 12.66, "2023-11-09": 13.33, "2023-11-10": 12.66, "2023-11-11": 13.33, "2023-11-12": 13.33}},
			{Type: "MACD", Values: map[string]float64{"2023-11-04": 13.66, "2023-11-05": 14.33, "2023-11-06": 14.66, "2023-11-07": 15.33, "2023-11-08": 13.66, "2023-11-09": 14.33, "2023-11-10": 15.66, "2023-11-11": 14.33, "2023-11-12": 15.33}},
			{Type: "SCT", Values: map[string]float64{"2023-11-04": 14.66, "2023-11-05": 15.33, "2023-11-06": 13.66, "2023-11-07": 14.33, "2023-11-08": 14.66, "2023-11-09": 15.33, "2023-11-10": 14.66, "2023-11-11": 15.33, "2023-11-12": 14.33}},
			// 必要な指標データを追加
		},
	}

	res, err := gcClient.GenerateChart(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// Base64エンコードされたデータをデコードしてPNG画像として保存
	outputDir := os.Getenv("TEST_CLIENT_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs"
	}

	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("generate_chart_client_test_%s.png", timestamp)
	outputFile := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Errorf("failed to create output directory: %v", err)
		return
	}

	// Base64データをデコード
	imageData, err := base64.StdEncoding.DecodeString(res.ChartData)
	if err != nil {
		t.Errorf("failed to decode base64 chart data: %v", err)
		return
	}

	// デコードしたデータをPNGファイルとして保存
	err = os.WriteFile(outputFile, imageData, 0644)
	if err != nil {
		t.Errorf("failed to save chart image: %v", err)
		return
	}

	fmt.Printf("Chart image saved successfully: %s\n", outputFile)
}

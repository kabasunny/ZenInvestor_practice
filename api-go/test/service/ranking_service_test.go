// api-go\test\service\ranking_service_test.go
package service_test

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"api-go/src/repository"
	"api-go/src/service"
	"api-go/test/repository/repository_test_helper"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetRankingDataIntegration(t *testing.T) {
	// 環境変数の読み込み
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// コンテキストの設定
	ctx, cancel := context.WithTimeout(context.Background(), 1800*time.Second)
	defer cancel()

	// データベースのセットアップ
	db := repository_test_helper.SetupTestDB()
	if db == nil {
		log.Fatalf("Failed to set up the database")
	}

	// リポジトリの初期化
	udsRepo := repository.NewUpdateStatusRepository(db)
	jsiRepo := repository.NewJpStockInfoRepository(db)
	jdpRepo := repository.NewJpDailyPriceRepository(db)
	j5mrRepo := repository.NewJp5dMvaRankingRepository(db)

	// サービスの初期化
	rankingService := service.NewRankingService(udsRepo, jsiRepo, jdpRepo, j5mrRepo, nil)

	// サービスの呼び出し
	res, err := rankingService.GetRankingData(ctx)
	if err != nil {
		fmt.Printf("Error calling GetRankingData service: %v\n", err)
	}
	fmt.Println("Service call completed.")

	// アサーション
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, *res) // レスポンスにデータが含まれていることを検証

	// CSV出力
	outputDir := os.Getenv("TEST_SERVICE_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "api-go/test/test_outputs" // デフォルトの出力ディレクトリ
	}

	timestamp := time.Now().Format("20060102150405") // タイムスタンプ
	filename := fmt.Sprintf("get_ranking_data_service_test_%s.csv", timestamp)
	outputFile := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Errorf("failed to create output directory: %v", err)
		return // エラーが発生してもテストを継続
	}

	file, err := os.Create(outputFile)
	if err != nil {
		t.Errorf("failed to create output file: %v", err)
		return // エラーが発生してもテストを継続
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ヘッダーを書き込む
	headers := []string{"Ranking", "Symbol", "Date", "AvgTurnover", "Name", "LatestClose"}
	if err := writer.Write(headers); err != nil {
		t.Errorf("failed to write headers to CSV file: %v", err)
		return
	}

	// データを書き込む
	for _, data := range *res {
		record := []string{
			fmt.Sprintf("%d", data.Ranking),
			data.Ticker,
			data.Date,
			fmt.Sprintf("%.2f", data.AvgTurnover),
			data.Name,
			fmt.Sprintf("%.2f", data.LatestClose),
		}
		if err := writer.Write(record); err != nil {
			t.Errorf("failed to write record to CSV file: %v", err)
		}
	}
	fmt.Println("Results written to CSV file successfully.")
}

// テストの実行コード
// go test -v ./test/service/ranking_service_test.go -run TestGetRankingDataIntegration

// 全マクロサービスを立ち上げるなら、/ZenInvestor_practiceにて
// ./StartMicroservices.ps1
// 本テスト用マイクロサービスを立ち上げるなら、/ZenInvestor_practiceにて
// ./StartForRankingMicroservices.ps1

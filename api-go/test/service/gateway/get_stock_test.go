package gateway_test

import (
	"context"
	"fmt"
	"os" // ファイル操作のためのパッケージ
	"testing"
	"time"

	"api-go/src/service/gateway"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetStockDataIntegration(t *testing.T) {
	// テスト用のコンテキストを15秒のタイムアウトで作成
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // テスト終了時にコンテキストをキャンセル

	address := "localhost:50051" // gRPCサーバーのアドレスを指定

	// gRPCクライアント接続を作成
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc.DialContext failed: %v", err) // 接続失敗時にテストを失敗として終了
	}
	defer conn.Close() // テスト終了時に接続を閉じる

	client := gateway.NewGetStockServiceClient(conn) // gRPCクライアントを作成
	req := &gateway.GetStockRequest{
		Ticker: "AAPL",
		Period: "5d",
	} // リクエストを作成

	// 呼び出し用のコンテキストを1秒のタイムアウトで作成
	callCtx, callCancel := context.WithTimeout(context.Background(), time.Second)
	defer callCancel() // 呼び出し終了時にコンテキストをキャンセル

	// gRPCメソッドを呼び出し
	res, err := client.GetStockData(callCtx, req)
	if err != nil {
		t.Fatalf("GetStockData failed: %v", err) // 呼び出し失敗時にテストを失敗として終了
	}

	assert.NoError(t, err)            // エラーがないことを確認
	assert.NotNil(t, res)             // レスポンスがnilでないことを確認
	assert.NotEmpty(t, res.StockData) // 株価データが空でないことを確認

	// 取得した株価データと日付を表示
	fmt.Printf("Stock data for %s:\n", req.Ticker)
	file, err := os.Create("get_stock_test_data.txt") // テキストファイルを作成
	if err != nil {
		t.Fatalf("Failed to create file: %v", err) // ファイル作成失敗時にテストを失敗として終了
	}
	defer file.Close() // テスト終了時にファイルを閉じる

	// 株価データをファイルに書き込み
	for key, value := range res.StockData {
		data := fmt.Sprintf("%s: open: %.2f, close: %.2f, high: %.2f, low: %.2f, volume: %.2f\n", key, value.Open, value.Close, value.High, value.Low, value.Volume)
		fmt.Println(data)                // データをターミナルに表示
		_, err := file.WriteString(data) // データをファイルに書き込み
		if err != nil {
			t.Fatalf("Failed to write to file: %v", err) // 書き込み失敗時にテストを失敗として終了
		}
	}
}

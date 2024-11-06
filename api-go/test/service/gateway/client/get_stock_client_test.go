package client_test

import (
	"context"
	"testing"
	"time"

	"api-go/src/service/ms_gateway/client"
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewStockClient(t *testing.T) {
	// クライアントを初期化
	godotenv.Load("../../../../.env") //テストではパスを指定しないとうまく読み取らない
	// 上記でgrpcクライアントのポートを読み込む必要がある

	ctx := context.Background()
	stockClient, err := client.NewGetStockDataClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stockClient)
	defer stockClient.Close() // 接続を閉じる
}

func TestGetStockData(t *testing.T) {
	// クライアントを初期化
	godotenv.Load("../../../../.env") //テストではパスを指定しないとうまく読み取らない
	// 上記でgrpcクライアントのポートを読み込む必要がある

	ctx := context.Background()
	stockClient, err := client.NewGetStockDataClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stockClient)
	defer stockClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ms_gateway.GetStockDataRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト
	res, err := stockClient.GetStockData(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.StockData)
}

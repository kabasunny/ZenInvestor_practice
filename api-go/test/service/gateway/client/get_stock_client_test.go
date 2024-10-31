package client_test

import (
	"context"
	"testing"
	"time"

	"api-go/src/service/gateway"
	"api-go/src/service/gateway/client"

	"github.com/stretchr/testify/assert"
)

func TestNewStockClient(t *testing.T) {
	client, err := client.NewGetStockClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	defer client.Close() // 接続を閉じる
}

func TestGetStockData(t *testing.T) {
	client, err := client.NewGetStockClient()
	assert.NoError(t, err)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &gateway.GetStockRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// 実際の gRPC サーバーが起動していることを前提とした統合テスト
	res, err := client.GetStockData(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.StockData)
}

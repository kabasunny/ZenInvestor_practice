package client_test

import (
	"context"
	"testing"
	"time"

	ms_gateway "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average"
	"api-go/src/service/ms_gateway/client"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleMovingAverageClient(t *testing.T) {
	godotenv.Load("../../../../.env")

	ctx := context.Background() // context.Background() を使用
	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, smaClient)
	defer smaClient.Close()
}

func TestCalculateSimpleMovingAverage(t *testing.T) {
	godotenv.Load("../../../../.env")

	ctx := context.Background()
	smaClient, err := client.NewSimpleMovingAverageClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, smaClient)
	defer smaClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ms_gateway.SimpleMovingAverageRequest{
		StockData:  []float32{10, 12, 14, 16, 18, 20},
		WindowSize: 3,
	}

	res, err := smaClient.CalculateSimpleMovingAverage(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	expectedSMA := []float64{12, 14, 16, 18}      // float64に変更
	assert.Equal(t, expectedSMA, res.GetValues()) // GetValues() を使用
}

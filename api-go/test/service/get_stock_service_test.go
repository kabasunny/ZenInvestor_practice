package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"api-go/src/service"
	"api-go/src/service/gateway"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStockClient は StockClient インターフェースのモック実装です。
type MockStockClient struct {
	mock.Mock
}

func (m *MockStockClient) GetStockData(ctx context.Context, req *gateway.GetStockRequest) (*gateway.GetStockResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*gateway.GetStockResponse), args.Error(1)
}

func (m *MockStockClient) Close() error {
	return nil
}

func TestGetStockDataSuccess(t *testing.T) {
	mockClient := new(MockStockClient)
	service := service.NewStockServiceImpl(mockClient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &gateway.GetStockRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// ダミーのレスポンスデータ（データ構造はチェックしないため、StockData を nil に設定）
	expectedResponse := &gateway.GetStockResponse{
		StockData: nil,
	}

	// モックの設定
	mockClient.On("GetStockData", ctx, req).Return(expectedResponse, nil)

	// テストの実行
	res, err := service.GetStockData(ctx, "AAPL", "5d")

	// エラーが発生せず、データが取得できたか確認
	assert.NoError(t, err)
	assert.NotNil(t, res)
	mockClient.AssertExpectations(t)
}

func TestGetStockDataFailure(t *testing.T) {
	mockClient := new(MockStockClient)
	service := service.NewStockServiceImpl(mockClient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &gateway.GetStockRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// エラーを返すようにモックを設定
	mockClient.On("GetStockData", ctx, req).Return((*gateway.GetStockResponse)(nil), fmt.Errorf("mock error: stock data not found"))

	// サービスメソッドの呼び出し
	res, err := service.GetStockData(ctx, "AAPL", "5d")

	// エラーが返されていることと、結果が nil であることを確認
	assert.Error(t, err)
	assert.Nil(t, res)
	mockClient.AssertExpectations(t)
}

package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"api-go/src/service"
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStockClient は StockClient インターフェースのモック実装です。
type MockStockClient struct {
	mock.Mock
}

func (m *MockStockClient) GetStockData(ctx context.Context, req *ms_gateway.GetStockDataRequest) (*ms_gateway.GetStockDataResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*ms_gateway.GetStockDataResponse), args.Error(1)
}

func (m *MockStockClient) Close() error {
	return nil
}

func TestGetStockDataSuccess(t *testing.T) {
	godotenv.Load("../../.env")

	mockClient := new(MockStockClient)
	clients := map[string]interface{}{
		"get_stock_data": mockClient,
	}
	service := service.NewStockServiceImpl(clients)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ms_gateway.GetStockDataRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// ダミーのレスポンスデータ（データ構造はチェックしないため、StockData を nil に設定）
	expectedResponse := &ms_gateway.GetStockDataResponse{
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
	godotenv.Load("../../.env") //テストではパスを指定しないとうまく読み取らない
	// 上記でgrpcクライアントのポートを読み込む必要がある
	mockClient := new(MockStockClient)
	clients := map[string]interface{}{
		"get_stock_data": mockClient,
	}
	service := service.NewStockServiceImpl(clients)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ms_gateway.GetStockDataRequest{
		Ticker: "AAPL",
		Period: "5d",
	}

	// エラーを返すようにモックを設定
	mockClient.On("GetStockData", ctx, req).Return((*ms_gateway.GetStockDataResponse)(nil), fmt.Errorf("mock error: stock data not found"))

	// サービスメソッドの呼び出し
	res, err := service.GetStockData(ctx, "AAPL", "5d")

	// エラーが返されていることと、結果が nil であることを確認
	assert.Error(t, err)
	assert.Nil(t, res)
	mockClient.AssertExpectations(t)
}

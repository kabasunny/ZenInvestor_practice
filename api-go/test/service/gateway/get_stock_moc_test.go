package gateway_test

import (
	"api-go/src/service/gateway" // 生成された gRPC クライアントコードをインポート
	"context"                    // コンテキストの使用
	"testing"                    // テストフレームワーク

	"github.com/stretchr/testify/assert" // アサーションライブラリ
	"github.com/stretchr/testify/mock"   // モックライブラリ
	"google.golang.org/grpc"             // gRPC フレームワーク
)

// Mock クライアントを定義
type MockGetStockServiceClient struct {
	mock.Mock // testify/mock を埋め込んでモックを作成
}

// Mock メソッドを実装
func (m *MockGetStockServiceClient) GetStockData(ctx context.Context, in *gateway.GetStockRequest, opts ...grpc.CallOption) (*gateway.GetStockResponse, error) {
	args := m.Called(ctx, in) // モックされたメソッドを呼び出す
	return args.Get(0).(*gateway.GetStockResponse), args.Error(1)
}

// TestGetStockData 関数
func TestGetStockData(t *testing.T) {
	mockClient := new(MockGetStockServiceClient)                                   // Mock クライアントを作成
	req := &gateway.GetStockRequest{Ticker: "AAPL"}                                // リクエストを作成
	res := &gateway.GetStockResponse{StockData: map[string]float64{"AAPL": 150.0}} // レスポンスを作成

	mockClient.On("GetStockData", mock.Anything, req).Return(res, nil) // モックの期待値を設定

	ctx := context.Background()                    // コンテキストを作成
	data, err := mockClient.GetStockData(ctx, req) // Mock メソッドを呼び出す

	assert.NoError(t, err)           // エラーがないことを確認
	assert.Equal(t, res, data)       // 期待されるデータと一致することを確認
	mockClient.AssertExpectations(t) // モックの期待値が満たされていることを確認
}

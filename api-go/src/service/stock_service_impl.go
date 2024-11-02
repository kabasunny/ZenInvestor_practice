package service

import (
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"

	"api-go/src/service/ms_gateway/client" // stock_client.go のパッケージパス

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StockServiceImpl は StockService インターフェースの実装
type StockServiceImpl struct {
	clients map[string]interface{} // 複数のgRPCクライアントを保持するマップ
}

// NewStockServiceImpl は StockServiceImpl の新しいインスタンスを作成
func NewStockServiceImpl(clients map[string]interface{}) StockService {
	return &StockServiceImpl{
		clients: clients,
	}
}

// GetStockData は指定された銘柄と期間の株価データを取得
func (s *StockServiceImpl) GetStockData(ctx context.Context, ticker string, period string) (*ms_gateway.GetStockDataResponse, error) {
	stockClient := s.clients["get_stock_data"].(client.GetStockDataClient)
	req := &ms_gateway.GetStockDataRequest{
		Ticker: ticker,
		Period: period,
	}

	res, err := stockClient.GetStockData(ctx, req)
	if err != nil {
		// エラー処理。必要に応じてより詳細なエラーハンドリングを行う
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, fmt.Errorf("stock data not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get stock data: %w", err)
	}

	return res, nil
}

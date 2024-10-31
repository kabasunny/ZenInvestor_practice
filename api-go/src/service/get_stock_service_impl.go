package service

import (
	"context"
	"fmt"

	"api-go/src/service/gateway"
	"api-go/src/service/gateway/client" // stock_client.go のパッケージパス

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StockServiceImpl は StockService インターフェースの実装です。
type StockServiceImpl struct {
	stockClient client.GetStockClient // gRPCクライアント
}

// NewStockServiceImpl は StockServiceImpl の新しいインスタンスを作成します。
func NewStockServiceImpl(stockClient client.GetStockClient) StockService {
	return &StockServiceImpl{
		stockClient: stockClient,
	}
}

// GetStockData は指定された銘柄と期間の株価データを取得します。
func (s *StockServiceImpl) GetStockData(ctx context.Context, ticker string, period string) (*gateway.GetStockResponse, error) {
	req := &gateway.GetStockRequest{
		Ticker: ticker,
		Period: period,
	}

	res, err := s.stockClient.GetStockData(ctx, req)
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

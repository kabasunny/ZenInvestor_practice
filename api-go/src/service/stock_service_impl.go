package service

import (
	"api-go/src/service/ms_gateway/client"
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"

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

// GetStockData は指定された銘柄と期間と指標の株価データを取得
func (s *StockServiceImpl) GetStockData(ctx context.Context, ticker string, period string) (*ms_gateway.GetStockDataResponse, error) {

	// GetStockDataClientのインスタンスを取得　:銘柄コード,表示期間
	getStockDataClient := s.clients["get_stock_data"].(client.GetStockDataClient)
	req := &ms_gateway.GetStockDataRequest{
		Ticker: ticker, // 銘柄コード
		Period: period, // 表示期間
	}

	// GetStockDataClientから株価のデータを取得
	res, err := getStockDataClient.GetStockData(ctx, req)
	if err != nil {
		// エラー処理。必要に応じてより詳細なエラーハンドリングを行う
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, fmt.Errorf("stock data not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get stock data: %w", err)
	}

	// GeneratChartClientのインスタンスを取得　:株価のデータ,指標

	// GeneratChartClientからチャート可視化データを取得

	// SimpleMovingAverageClientのインスタンスを取得　:株価のデータ,平均値の計算幅

	// SimpleMovingAverageClientから移動平均線を取得

	return res, nil
}

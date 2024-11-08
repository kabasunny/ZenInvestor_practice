package service

import (
	indicator "api-go/src/service/ms_gateway/calculate_indicator" // IndicatorParamsのimport元
	// smaパッケージをインポート
	client "api-go/src/service/ms_gateway/client"
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// anypbを使う場合に必要
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
func (s *StockServiceImpl) GetStockData(ctx context.Context, ticker string, period string, indicators []*indicator.IndicatorParams) (*getstockdata.GetStockDataResponse, error) {

	// GetStockDataClientのインスタンスを取得　:銘柄コード,表示期間
	getStockDataClient := s.clients["get_stock_data"].(client.GetStockDataClient)
	req := &getstockdata.GetStockDataRequest{
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

	// indicatorsがnilまたは空の場合、指標の計算はスキップ
	if indicators != nil || len(indicators) != 0 {
		// indicatorsが有効な値の場合、指標の計算を実行
		// for _, indicator := range indicators {
		//指標は最大3個まで付加し、チャートを生成する

		// もし、SimpleMovingAverageClientのインスタンスを取得　:株価のデータ,平均値の計算幅
		// SimpleMovingAverageClientから移動平均線作画用データを取得
		// もし、他の指標計算Clientのインスタンスを取得　:株価のデータ,平均値の計算幅
		// 他の指標1計算Clientから他の指標1データを取得
		// もし、他の指標計算Clientのインスタンスを取得　:株価のデータ,平均値の計算幅
		// 他の指標2計算Clientから他の指標2データを取得
		// }
	}

	// GeneratChartClientのインスタンスを取得　:株価のデータ,指標データ0～3個

	// GeneratChartClientからチャート可視化データを取得

	return res, nil
}

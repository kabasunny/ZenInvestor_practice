package service

import (
	indicator "api-go/src/service/ms_gateway/calculate_indicator"                 // IndicatorParamsのimport元
	sma "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average" // smaパッケージをインポート
	client "api-go/src/service/ms_gateway/client"
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"
	"log"
	"strconv"

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
	if len(indicators) > 0 {
		// if indicators != nil && len(indicators) > 0 { // len()関数を使用すると、nilスライスでも長さはゼロとして返される
		for _, indicator := range indicators {
			switch indicator.Type {
			case "SMA":
				// SimpleMovingAverageClientのインスタンスを取得
				smaClient := s.clients["simple_moving_average"].(client.SimpleMovingAverageClient)

				// getstockdata.StockDataをsma.StockDataForSMAにデータを変換...めんどくせーぞ
				convertedStockData := convertStockDataForSMA(res.StockData)

				// WindowSizeを文字列からint32に変換
				windowSizeStr := indicator.Params["window_size"]           // パラメータからウィンドウサイズを取得
				windowSize, err := strconv.ParseInt(windowSizeStr, 10, 32) // 文字列をint32に変換...めんどくせーぞ
				if err != nil {
					log.Fatalf("window_sizeをint32に変換できませんでした: %v", err)
				}

				smaReq := &sma.SimpleMovingAverageRequest{
					StockData:  convertedStockData, // 株価データを取得
					WindowSize: int32(windowSize),  // 変換したウィンドウサイズを設定
				}
				smaRes, err := smaClient.CalculateSimpleMovingAverage(ctx, smaReq)
				if err != nil {
					return nil, fmt.Errorf("failed to calculate SMA: %w", err)
				}
				// 必要に応じて SMA 結果を格納・処理
				fmt.Printf("SMA result: %v\n", smaRes.MovingAverage)

			case "OtherIndicator1":
				// 他の指標1計算のクライアントを取得・実行
				// 他の指標クライアントに合わせた処理をここで実装

			case "OtherIndicator2":
				// 他の指標2計算のクライアントを取得・実行
				// 他の指標クライアントに合わせた処理をここで実装

			default:
				fmt.Printf("Unsupported indicator: %s\n", indicator.Type)
			}
		}
	}

	// GeneratChartClientのインスタンスを取得　:株価のデータ,指標データ0～3個

	// GeneratChartClientからチャート可視化データを取得

	return res, nil
}

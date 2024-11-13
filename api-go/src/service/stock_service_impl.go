// api-go\src\service\stock_service_impl.go
package service

import (
	indicator "api-go/src/service/ms_gateway/calculate_indicator"
	sma "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average"
	client "api-go/src/service/ms_gateway/client"
	getstockdata "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"
	"log"
	"strconv"

	gc "api-go/src/service/ms_gateway/generate_chart"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StockServiceImpl は StockService インターフェースの実装
type StockServiceImpl struct {
	clients map[string]interface{}
}

// NewStockServiceImpl は StockServiceImpl の新しいインスタンスを作成
func NewStockServiceImpl(clients map[string]interface{}) StockService {
	return &StockServiceImpl{
		clients: clients,
	}
}

// GetStockData は指定された銘柄と期間と指標の株価データを取得
func (s *StockServiceImpl) GetStockData(ctx context.Context, ticker string, period string) (*getstockdata.GetStockDataResponse, error) {
	stockClient := s.clients["get_stock_data"].(client.GetStockDataClient)
	req := &getstockdata.GetStockDataRequest{
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

// GetStockChart は指定された銘柄と期間と指標の株価データを取得
func (s *StockServiceImpl) GetStockChart(ctx context.Context, ticker string, period string, indicators []*indicator.IndicatorParams, includeVolume bool) (*gc.GenerateChartResponse, error) {

	getStockDataClient := s.clients["get_stock_data"].(client.GetStockDataClient)
	req := &getstockdata.GetStockDataRequest{
		Ticker: ticker,
		Period: period,
	}

	// GetStockDataClientから株価のデータを取得
	res, err := getStockDataClient.GetStockData(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, fmt.Errorf("stock data not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get stock data: %w", err)
	}

	indicatorDataList := make([]*gc.IndicatorData, 0) // nilを0に変更

	// indicatorsがnilまたは空の場合、指標の計算はスキップ
	if len(indicators) > 0 {
		for _, indicator := range indicators {
			switch indicator.Type {
			case "SMA":
				// SimpleMovingAverageClientのインスタンスを取得
				smaClient := s.clients["simple_moving_average"].(client.SimpleMovingAverageClient)

				convertedStockData := convertStockDataForSMA(res.StockData)

				windowSizeStr := indicator.Params["window_size"]
				windowSize, err := strconv.ParseInt(windowSizeStr, 10, 32)
				if err != nil {
					log.Printf("window_sizeをint32に変換できませんでした: %v", err)
					continue
				}

				smaReq := &sma.SimpleMovingAverageRequest{
					StockData:  convertedStockData,
					WindowSize: int32(windowSize),
				}
				smaRes, err := smaClient.CalculateSimpleMovingAverage(ctx, smaReq)
				if err != nil {
					return nil, fmt.Errorf("failed to calculate SMA: %w", err)
				}

				valuesMap := make(map[string]float64)
				for date, value := range smaRes.MovingAverage {
					valuesMap[date] = value
				}
				// log.Printf("SMA Indicator, valuesMap: %v\n", valuesMap) // 格納した段階でログ表示

				legendName := fmt.Sprintf("%s%s", indicator.Type, windowSizeStr)
				fmt.Println(legendName)

				indicatorDataList = append(indicatorDataList, &gc.IndicatorData{
					Type:       indicator.Type,
					Values:     valuesMap,
					LegendName: legendName, // 凡例の名称を設定
				})

			// ここに別の指標の計算ロジックを追加
			// case "MACD":
			//  macdClient := s.clients["macd"].(client.MACDClient)
			//  convertedStockData := convertStockDataForMACD(res.StockData)
			//  macdReq := &macd.MACDRequest{
			//      StockData: convertedStockData,
			//      // 他の必要なパラメータを設定
			//  }
			//  macdRes, err := macdClient.CalculateMACD(ctx, macdReq)
			//  if err != nil {
			//      return nil, fmt.Errorf("failed to calculate MACD: %w", err)
			//  }
			//  valuesMap := make(map[string]float64)
			//  for date, value := range macdRes.MovingAverage {
			//      valuesMap[date] = value
			//  }
			//  log.Printf("MACD Indicator, valuesMap: %v\n", valuesMap) // 格納した段階でログ表示
			//  indicatorDataList = append(indicatorDataList, &gc.IndicatorData{
			//      Type:   indicator.Type,
			//      Values: valuesMap,
			//  })

			default:
				fmt.Printf("Unsupported indicator: %s\n", indicator.Type)
			}
		}
	}

	generateChartClient := s.clients["generate_chart"].(client.GenerateChartClient)

	stockDataMap := make(map[string]*gc.StockDataForChart)
	for key, data := range res.StockData {
		stockDataMap[key] = &gc.StockDataForChart{
			Open:   data.Open,
			Close:  data.Close,
			High:   data.High,
			Low:    data.Low,
			Volume: data.Volume,
		}
	}

	generateChartReq := &gc.GenerateChartRequest{
		StockData:     stockDataMap,
		Indicators:    indicatorDataList,
		IncludeVolume: includeVolume, // 出来高の要否を含める
	}

	generateChartRes, err := generateChartClient.GenerateChart(ctx, generateChartReq)
	if err != nil {
		return nil, fmt.Errorf("failed to generate chart: %w", err)
	}

	// fmt.Printf("Chart data: %s\n", generateChartRes.ChartData)

	return generateChartRes, nil
}

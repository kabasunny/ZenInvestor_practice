// api-go\src\infra\ms_clients_for_app.go
package infra

import (
	"context"
	"fmt"
	"log"

	"api-go/src/service/ms_gateway/client"
)

// infra/client.go 内の SetupMsClients 関数も修正
func SetupMsClientsForApp(ctx context.Context) (*MSClients, error) { // 戻り値にerrorを追加
	fmt.Println("in SetupMsClients for APP.")
	msClients := make(map[string]interface{})
	fmt.Println("Clients setup...")

	// 株価データ取得用クライアント
	getStockDataClient, err := client.NewGetStockDataClient(ctx)
	fmt.Println("in NewGetStockDataClient.")
	if err != nil {
		log.Fatalf("Failed to create get stock data client: %v", err)
	}
	msClients["get_stock_data"] = getStockDataClient
	fmt.Println("getStockDataClient setup successfully.")

	// チャート可視化データ取得用クライアント
	generateChartClient, err := client.NewGenerateChartClient(ctx)
	fmt.Println("in NewGenerateChartClient.")
	if err != nil {
		return nil, fmt.Errorf("failed to setup generate chart client: %w", err)
	}
	msClients["generate_chart"] = generateChartClient
	fmt.Println("generateChartClient setup successfully.")

	// 単純移動平均データ取得用クライアント
	smaClient, err := client.NewSimpleMovingAverageClient(ctx) // SimpleMovingAverageClient を初期化
	fmt.Println("in NewSimpleMovingAverageClient.")
	if err != nil {
		log.Fatalf("Failed to create simple moving average client: %v", err)
	}
	msClients["simple_moving_average"] = smaClient // mapに追加
	fmt.Println("smaClient setup successfully.")

	// 期間を指定した株価データ取得用クライアント
	getStockDataWithDatesClient, err := client.NewGetStockDataWithDatesClient(ctx)
	fmt.Println("in NewGetStockDataWithDatesClient.")
	if err != nil {
		log.Fatalf("Failed to create get stock data client with dates: %v", err)
	}
	msClients["get_stock_data_with_dates"] = getStockDataWithDatesClient
	fmt.Println("getStockDataWithDatesClient setup successfully.")

	// チャート可視化データ取得用クライアント
	generateChartLCClient, err := client.NewGenerateChartLCClient(ctx)
	fmt.Println("in NewGenerateChartLCClient.")
	if err != nil {
		return nil, fmt.Errorf("failed to setup generate chart losscut simulation client: %w", err)
	}
	msClients["generate_chart_lc_sim"] = generateChartLCClient
	fmt.Println("generateChartLCClient setup successfully.")

	// 他のマイクロサービス用クライアントの初期化もここに追加

	return &MSClients{
		MSClients: msClients,
	}, nil // nilエラーを返す

}

package infra

import (
	"context"
	"fmt"
	"log"

	"api-go/src/service/ms_gateway/client"
)

// MSClients holds the gRPC clients for microservices
type MSClients struct {
	MSClients map[string]interface{}
}

// infra/client.go 内の SetupMsClients 関数も修正
func SetupMsClients(ctx context.Context) (*MSClients, error) { // 戻り値にerrorを追加
	fmt.Println("in SetupMsClients.")
	msClients := make(map[string]interface{})
	fmt.Println("Clients setup...")

	getStockDataClient, err := client.NewGetStockDataClient(ctx)
	fmt.Println("in NewGetStockDataClient.")
	if err != nil {
		log.Fatalf("Failed to create get stock data client: %v", err)
	}
	msClients["get_stock_data"] = getStockDataClient
	fmt.Println("getStockDataClient setup successfully.")

	// 単純移動平均データ取得
	smaClient, err := client.NewSimpleMovingAverageClient(ctx) // SimpleMovingAverageClient を初期化
	fmt.Println("in NewSimpleMovingAverageClient.")
	if err != nil {
		log.Fatalf("Failed to create simple moving average client: %v", err)
	}
	msClients["simple_moving_average"] = smaClient // mapに追加
	fmt.Println("smaClient setup successfully.")

	// 他のマイクロサービス用クライアントの初期化もここに追加

	return &MSClients{
		MSClients: msClients,
	}, nil // nilエラーを返す

}

package infra

import (
	"context"
	"log"

	"api-go/src/service/ms_gateway/client"
)

// MSClients holds the gRPC clients for microservices
type MSClients struct {
	MSClients map[string]interface{}
}

// infra/client.go 内の SetupMsClients 関数も修正
func SetupMsClients(ctx context.Context) (*MSClients, error) { // 戻り値にerrorを追加
	msClients := make(map[string]interface{})

	// 株価データ取得　SMAクライアントがうまくいったら、クライアントを修正
	getStockDataClient, err := client.NewGetStockDataClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create get stock data client: %v", err)
	}
	msClients["get_stock_data"] = getStockDataClient

	// 単純移動平均データ取得
	smaClient, err := client.NewSimpleMovingAverageClient(ctx) // SimpleMovingAverageClient を初期化
	if err != nil {
		log.Fatalf("Failed to create simple moving average client: %v", err)
	}
	msClients["simple_moving_average"] = smaClient // mapに追加

	// 他のマイクロサービス用クライアントの初期化もここに追加

	return &MSClients{
		MSClients: msClients,
	}, nil // nilエラーを返す

}

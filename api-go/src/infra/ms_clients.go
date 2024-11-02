package infra

import (
	"log"

	"api-go/src/service/ms_gateway/client"
)

// MSClients holds the gRPC clients for microservices
type MSClients struct {
	MSClients map[string]interface{}
}

// SetupMsClients initializes all necessary gRPC clients
func SetupMsClients() *MSClients {
	msClients := make(map[string]interface{})

	stockClient, err := client.NewGetStockDataClient()
	if err != nil {
		log.Fatalf("Failed to create stock client: %v", err)
	}
	msClients["get_stock_data"] = stockClient

	// 他のマイクロサービス用クライアントの初期化もここに追加

	return &MSClients{
		MSClients: msClients,
	}
}

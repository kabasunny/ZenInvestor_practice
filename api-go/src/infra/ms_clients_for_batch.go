// api-go\src\infra\ms_clients_for_batch.go
package infra

import (
	"context"
	"fmt"
	"log"

	"api-go/src/service/ms_gateway/client"
)

// バッチ処理用のマイクロサービス
func SetupMsClientsForBatch(ctx context.Context) (*MSClients, error) { // 戻り値にerrorを追加
	fmt.Println("in SetupMsClients for Batch.")
	msClients := make(map[string]interface{})
	fmt.Println("Clients setup...")

	// J-QUANTSからの銘柄データ取得用クライアント
	gsijClient, err := client.NewGetStockInfoJqClient(ctx)
	fmt.Println("in NewGetStockInfoJqClient.")
	if err != nil {
		log.Fatalf("Failed to create get stock info from jq client: %v", err)
	}
	msClients["get_stock_info_jq"] = gsijClient // mapに追加
	fmt.Println("gsijClient setup successfully.")

	// J-QUANTSからの休日データ取得用クライアント
	gtcjClient, err := client.NewGetTradingCalendarJqClient(ctx)
	fmt.Println("in NewGetTradingCalendarJqClient.")
	if err != nil {
		log.Fatalf("Failed to create get stock info from jq client: %v", err)
	}
	msClients["get_trading_calendar_jq"] = gtcjClient // mapに追加
	fmt.Println("gtcjClient setup successfully.")

	// ランキング用、J-QUANTSからの全株価データ取得用クライアント
	gsdwdClient, err := client.NewGetStocksDatalistWithDatesClient(ctx)
	fmt.Println("in NewGetStockInfoJqClient.")
	if err != nil {
		log.Fatalf("Failed to create get stocks datalist with dates client: %v", err)
	}
	msClients["get_stocks_datalist_with_dates"] = gsdwdClient // mapに追加
	fmt.Println("gsdwdClient setup successfully.")

	// 他のマイクロサービス用クライアントの初期化もここに追加

	return &MSClients{
		MSClients: msClients,
	}, nil // nilエラーを返す

}

// api-go\cmd\batch_update_stock_info\main.go

package main

import (
	"api-go/src/batch"
	"api-go/src/infra"
	"api-go/src/repository"
	client "api-go/src/service/ms_gateway/client"
	"context"
	"log"
	// その他必要なインポート
)

func main() {

	infra.Initialize() // 初期化処理
	db := infra.SetupDB()
	udsRepo := repository.NewUpdateStatusRepository(db)
	jsiRepo := repository.NewJpStockInfoRepository(db)

	ctx := context.Background()
	clients := make(map[string]interface{})

	// client.NewGetStockInfoJqClient から2つの値を受け取る
	gsijClient, err := client.NewGetStockInfoJqClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	// clientsの初期化
	clients["get_stock_info_jq"] = gsijClient

	err = batch.UpdateStockInfo(ctx, udsRepo, jsiRepo, clients)
	if err != nil {
		log.Fatalf("Failed to update stock info: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_stock_info/main.go
// api-go\cmd\batch_update_stock_info\batch_update_stock_info.go

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
	// clients := make(map[string]interface{})

	// client.NewGetStockInfoJqClient から2つの値を受け取る
	gsijClient, err := client.NewGetStockInfoJqClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	// clientsの初期化
	// clients["get_stock_info_jq"] = gsijClient

	err = batch.UpdateStockInfo(ctx, udsRepo, jsiRepo, gsijClient)
	if err != nil {
		log.Fatalf("Failed to update stock info: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_stock_info/batch_update_stock_info.go

// UpdateStockInfo completed in 7m32.1925772s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz Goルーチン無し、アップデートメソッド
// UpdateStockInfo completed in 18.1971501s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz Goルーチン有り、アップデートメソッド

// UpdateStockInfo completed in 5.798971s : 12th Gen Intel(R) Core(TM) i7-1255U   1.70 GHz Goルーチン無し、デリートメソッド + インサートメソッド
// UpdateStockInfo completed in 12.1142565s : Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz  3.41 GHz Goルーチン無し、デリートメソッド + インサートメソッド

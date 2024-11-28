// api-go\cmd\batch_update_ranking\batch_update_ranking.go
package main

import (
	"api-go/src/batch"
	"api-go/src/infra"
	"api-go/src/repository"
	"context"
	"log"
	// その他必要なインポート
)

func main() {

	infra.Initialize() // 初期化処理
	db := infra.SetupDB()
	udsRepo := repository.NewUpdateStatusRepository(db)
	j5mrRepo := repository.NewJp5dMvaRankingRepository(db)

	ctx := context.Background()

	err := batch.UpdateRanking(ctx, udsRepo, j5mrRepo)
	if err != nil {
		log.Fatalf("Failed to update ranking: %v", err)
	}
}

// 実行コマンド
// go run ./cmd/batch_update_ranking/batch_update_ranking.go

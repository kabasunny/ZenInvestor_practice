package main

import (
	"api-go/src/infra"
	seed "api-go/src/seed/insert_data"
	"log"
)

func main() {
	// 初期化処理、主に環境変数読み込み
	infra.Initialize()

	// データベースのセットアップ
	db := infra.SetupDB()

	// 初期データの挿入
	if err := seed.InsertInitialData(db); err != nil {
		log.Fatalf("Failed to insert initial data: %v", err)
	}
	log.Println("Initial data inserted successfully")
}

// 実行コマンド　api-go/にて
// go run ./src/seed/initialize_data.go

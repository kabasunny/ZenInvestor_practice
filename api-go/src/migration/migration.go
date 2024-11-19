package main

import (
	"api-go/src/infra"
	migration "api-go/src/migration/tasks"
	"log"
)

func main() {
	// 初期化処理、主に環境変数読み込み
	infra.Initialize()

	// データベースのセットアップ
	db := infra.SetupDB()

	// マイグレーションの実行 migration_tasks.go に実行内容を定義している
	if err := migration.MigrateUp(db); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("Migration completed successfully")
}

// 実行コマンド　api-go/にて
// go run ./src/migration/migration.go

// api-go\main.go

package main

import (
	"api-go/src/infra"
	"api-go/src/router"
	"context"
	"log" // エラーログ出力用

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background() // バックグラウンドコンテキストを作成
	infra.Initialize()          // 初期化処理

	db := infra.SetupDB() // データベース接続の初期化
	// var db *gorm.DB = nil

	msClients, err := infra.SetupMsClientsForApp(ctx) // マイクロサービス接続の初期化、エラー処理を追加　// アプリケーション用とバッチ処理用を分離
	if err != nil {
		log.Fatalf("Failed to setup microservice clients: %v", err) // エラーが発生したらログ出力して終了
	}

	ginRouter := gin.Default()                   // Ginのデフォルトルータを作成
	router.SetupRouter(ginRouter, db, msClients) // ルーティングの設定

	if err := ginRouter.Run(":8086"); err != nil { // ginRouter.Run()のエラー処理
		log.Fatalf("Failed to run server: %v", err)
	}
}

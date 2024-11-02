package main

import (
	"api-go/src/infra"
	"api-go/src/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	infra.Initialize() // 初期化処理

	// db := infra.SetupDB()               // データベース接続の初期化
	var db *gorm.DB = nil

	msClients := infra.SetupMsClients() // マイクロサービス接続の初期化

	ginRouter := gin.Default()                   // Ginのデフォルトルータを作成
	router.SetupRouter(ginRouter, db, msClients) // ルーティングの設定
	ginRouter.Run(":8086")                       // サーバーをポート8086で起動
}

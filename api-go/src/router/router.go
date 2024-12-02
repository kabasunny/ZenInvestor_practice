// api-go\src\router\router.go
package router

import (
	"api-go/src/controller"
	"api-go/src/infra"
	"api-go/src/middleware"
	"api-go/src/repository"
	"api-go/src/service"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB, msClients *infra.MSClients) {

	// 環境変数からCORS設定を読み込む
	allowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	origins := strings.Split(allowOrigins, ",")

	// CORS設定を適用
	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,                                             // 許可するオリジンを設定
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 許可するHTTPメソッドを設定
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 許可するHTTPヘッダーを設定
		ExposeHeaders:    []string{"Content-Length"},                          // レスポンスヘッダーに含めるヘッダーを設定
		AllowCredentials: true,                                                // クッキーなどの認証情報を許可するかどうかを設定
	}))

	// リポジトリの初期化
	udsRepo := repository.NewUpdateStatusRepository(db)
	jsiRepo := repository.NewJpStockInfoRepository(db)
	jdpRepo := repository.NewJpDailyPriceRepository(db)
	j5mrRepo := repository.NewJp5dMvaRankingRepository(db)

	// サービスの初期化
	rankingService := service.NewRankingService(udsRepo, jsiRepo, jdpRepo, j5mrRepo, msClients.MSClients)
	stockService := service.NewStockServiceImpl(msClients.MSClients)

	// コントローラの初期化
	stockController := controller.NewStockControllerImpl(stockService, rankingService)

	// ミドルウェアの適用
	// router.Use(middleware.AuthMiddleware())          // 認証ミドルウェア まだトークン生成を実装してない
	router.Use(middleware.RateLimitMiddleware(1, 5)) // レートリミットミドルウェア

	// 株価データ
	router.GET("/getStockData", stockController.GetStockData)

	// 株価チャート
	router.POST("/getStockChart", stockController.GetStockChart)
	// http://localhost:8086/getStockChart
	// {
	//  "ticker": "AAPL",
	//  "period": "1y",
	//  "indicators": [
	//    { "type": "SMA", "params": { "window_size": "20" } }
	//  ]
	// }

	// ランキングデータ取得用エンドポイント
	router.GET("/RankingByRange", stockController.GetRankingDataByRange) // http://localhost:8086/RankingByRange?startRank=1&endRank=100
	router.GET("/InitialRanking", stockController.GetInitialRanking)     //http://localhost:8086/InitialRanking

	// // ユーザーログイン用　後で
	// loginRepository := repository.NewLoginRepository(db)
	// loginService := service.NewLoginService(loginRepository)
	// loginController := controller.NewLoginController(loginService)
	// router.POST("/login", loginController.Login)
}

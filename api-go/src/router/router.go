package router

import (
	"api-go/src/controller"
	"api-go/src/infra"
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

	// ストックデータ取得用
	stockService := service.NewStockServiceImpl(msClients.MSClients)
	stockController := controller.NewStockControllerImpl(stockService)
	// 株価データ
	router.GET("/getStockData", stockController.GetStockData)

	// 株価チャート
	router.POST("/getStockChart", stockController.GetStockChart)
	// http://localhost:8086/getStockChart
	// {
	// 	"ticker": "AAPL",
	// 	"period": "1y",
	// 	"indicators": [
	// 	  { "type": "SMA", "params": { "window_size": "20" } }
	// 	]
	//   }

	// // ユーザーログイン用　後で
	// loginRepository := repository.NewLoginRepository(db)
	// loginService := service.NewLoginService(loginRepository)
	// loginController := controller.NewLoginController(loginService)
	// router.POST("/login", loginController.Login)

	// // スペシャルランキングデータ取得用（認証が必要）後で
	// stockRankingRepository := repository.NewStockRankingRepository(db)
	// stockRankingService := service.NewStockRankingService(stockRankingRepository)
	// stockRankingController := controller.NewStockRankingController(stockRankingService)
	// authRouter := router.Group("/auth")
	// authRouter.Use(middleware.AuthMiddleware()) // ミドルウェアを適用
	// stockRankingRouter := authRouter.Group("/stockRanking")
	// stockRankingRouter.GET("/", stockRankingController.GetStockRankingData)
}

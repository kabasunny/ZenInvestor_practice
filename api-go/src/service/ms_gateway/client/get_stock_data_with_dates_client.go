// api-go\src\service\ms_gateway\client\get_stock_data_with_dates_client.go
package client

import (
	get_stock_data_with_dates "api-go/src/service/ms_gateway/get_stock_data_with_dates"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GetStockDataWithDatesClient は株価データ取得サービスのgRPCクライアント
type GetStockDataWithDatesClient interface {
	GetStockData(ctx context.Context, req *get_stock_data_with_dates.GetStockDataWithDatesRequest) (*get_stock_data_with_dates.GetStockDataWithDatesResponse, error)
	Close() error
}

// getStockDataWithDatesClientImpl は GetStockDataWithDatesClient インターフェースの実装
type getStockDataWithDatesClientImpl struct {
	client get_stock_data_with_dates.GetStockDataWithDatesServiceClient
	conn   *grpc.ClientConn
}

// NewGetStockDataWithDatesClient は GetStockDataWithDatesClient の新しいインスタンスを作成
func NewGetStockDataWithDatesClient(ctx context.Context) (GetStockDataWithDatesClient, error) {
	port := os.Getenv("GET_STOCK_DATA_WITH_DATES_MS_PORT") // .envを確認
	address := fmt.Sprintf("localhost:%s", port)           // 環境変数からポート番号を取得

	// ClientConn を作成
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	// 接続を開始
	conn.Connect()

	// タイムアウトを設定 (15秒) 必要に応じて調整
	connectCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 接続が確立されるまで待つ
	for conn.GetState() != connectivity.Ready {
		if !conn.WaitForStateChange(connectCtx, conn.GetState()) {
			return nil, fmt.Errorf("failed to connect to gRPC server: %w", connectCtx.Err())
		}
	}

	client := get_stock_data_with_dates.NewGetStockDataWithDatesServiceClient(conn)
	return &getStockDataWithDatesClientImpl{client: client, conn: conn}, nil
}

// GetStockData は指定された銘柄と期間の株価データを取得
func (c *getStockDataWithDatesClientImpl) GetStockData(ctx context.Context, req *get_stock_data_with_dates.GetStockDataWithDatesRequest) (*get_stock_data_with_dates.GetStockDataWithDatesResponse, error) {
	return c.client.GetStockData(ctx, req)
}

// Close はgRPC接続を閉じる
func (c *getStockDataWithDatesClientImpl) Close() error {
	return c.conn.Close()
}

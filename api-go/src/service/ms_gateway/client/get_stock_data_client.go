package client

import (
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GetStockDataClient は株価データ取得サービスのgRPCクライアント
type GetStockDataClient interface {
	GetStockData(ctx context.Context, req *ms_gateway.GetStockDataRequest) (*ms_gateway.GetStockDataResponse, error)
	Close() error
}

// getStockClientImpl は GetStockClient インターフェースの実装
type getStockDataClientImpl struct {
	client ms_gateway.GetStockDataServiceClient
	conn   *grpc.ClientConn
}

// NewGetStockClient は GetStockClient の新しいインスタンスを作成
func NewGetStockDataClient(ctx context.Context) (GetStockDataClient, error) {
	port := os.Getenv("GET_STOCK_DATA_MS_PORT")  // .envを確認
	address := fmt.Sprintf("localhost:%s", port) // 環境変数からポート番号を取得

	// NewClient を使用して ClientConn を作成
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	// 接続を開始
	conn.Connect()

	// タイムアウトを設定 (例: 15秒)  必要に応じて調整
	connectCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 接続が確立されるまで待つ
	for conn.GetState() != connectivity.Ready { //ここを変更
		if !conn.WaitForStateChange(connectCtx, conn.GetState()) {
			return nil, fmt.Errorf("failed to connect to gRPC server: %w", connectCtx.Err())
		}
	}

	client := ms_gateway.NewGetStockDataServiceClient(conn)
	return &getStockDataClientImpl{client: client, conn: conn}, nil
}

// GetStockData は指定された銘柄と期間の株価データを取得
func (c *getStockDataClientImpl) GetStockData(ctx context.Context, req *ms_gateway.GetStockDataRequest) (*ms_gateway.GetStockDataResponse, error) {
	return c.client.GetStockData(ctx, req)
}

// Close はgRPC接続を閉じます。
func (c *getStockDataClientImpl) Close() error {
	return c.conn.Close()
}

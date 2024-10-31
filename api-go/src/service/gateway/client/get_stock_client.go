package client

import (
	"api-go/src/service/gateway"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetStockClient は株価データ取得サービスのgRPCクライアント
type GetStockClient interface {
	GetStockData(ctx context.Context, req *gateway.GetStockRequest) (*gateway.GetStockResponse, error)
	Close() error
}

// getStockClientImpl は GetStockClient インターフェースの実装
type getStockClientImpl struct {
	client gateway.GetStockServiceClient
	conn   *grpc.ClientConn
}

// NewGetStockClient は GetStockClient の新しいインスタンスを作成
func NewGetStockClient() (GetStockClient, error) {
	// err := godotenv.Load()
	// if err != nil {
	//     return nil, fmt.Errorf("failed to load .env file: %w", err)
	// }
	// port := os.Getenv("GET_STOCK_MS_PORT")
	// address := fmt.Sprintf("localhost:%s", port) // 環境変数からポート番号を取得
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial gRPC server: %w", err)
	}
	client := gateway.NewGetStockServiceClient(conn)
	return &getStockClientImpl{client: client, conn: conn}, nil
}

// GetStockData は指定された銘柄と期間の株価データを取得します。
func (c *getStockClientImpl) GetStockData(ctx context.Context, req *gateway.GetStockRequest) (*gateway.GetStockResponse, error) {
	return c.client.GetStockData(ctx, req)
}

// Close はgRPC接続を閉じます。
func (c *getStockClientImpl) Close() error {
	return c.conn.Close()
}

package client

import (
	ms_gateway "api-go/src/service/ms_gateway/get_stock_data"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
func NewGetStockDataClient() (GetStockDataClient, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	port := os.Getenv("GET_STOCK_DATA_MS_PORT")
	address := fmt.Sprintf("localhost:%s", port) // 環境変数からポート番号を取得
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial gRPC server: %w", err)
	}
	client := ms_gateway.NewGetStockDataServiceClient(conn)
	return &getStockDataClientImpl{client: client, conn: conn}, nil
}

// GetStockData は指定された銘柄と期間の株価データを取得します。
func (c *getStockDataClientImpl) GetStockData(ctx context.Context, req *ms_gateway.GetStockDataRequest) (*ms_gateway.GetStockDataResponse, error) {
	return c.client.GetStockData(ctx, req)
}

// Close はgRPC接続を閉じます。
func (c *getStockDataClientImpl) Close() error {
	return c.conn.Close()
}

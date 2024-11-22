// api-go\src\service\ms_gateway\client\get_stock_info_jq_client.go

package client

import (
	get_stock_info_jq "api-go/src/service/ms_gateway/get_stock_info_jq"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GetStockInfoJqClient は株式情報取得サービスのgRPCクライアント
type GetStockInfoJqClient interface {
	GetStockInfoJq(ctx context.Context, req *get_stock_info_jq.GetStockInfoJqRequest) (*get_stock_info_jq.GetStockInfoJqResponse, error)
	Close() error
}

// getStockInfoJqClientImpl は GetStockInfoJqClient インターフェースの実装
type getStockInfoJqClientImpl struct {
	client get_stock_info_jq.GetStockInfoJqServiceClient
	conn   *grpc.ClientConn
}

// NewGetStockInfoJqClient は GetStockInfoJqClient の新しいインスタンスを作成
func NewGetStockInfoJqClient(ctx context.Context) (GetStockInfoJqClient, error) {
	port := os.Getenv("GET_STOCK_INFO_JQ_MS_PORT") // .envを確認
	address := fmt.Sprintf("localhost:%s", port)   // 環境変数からポート番号を取得

	// NewClient を使用して ClientConn を作成
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	// 接続を開始
	conn.Connect()

	// タイムアウトを設定 (15秒)  必要に応じて調整
	connectCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 接続が確立されるまで待つ
	for conn.GetState() != connectivity.Ready {
		if !conn.WaitForStateChange(connectCtx, conn.GetState()) {
			return nil, fmt.Errorf("failed to connect to gRPC server: %w", connectCtx.Err())
		}
	}

	client := get_stock_info_jq.NewGetStockInfoJqServiceClient(conn)
	return &getStockInfoJqClientImpl{client: client, conn: conn}, nil
}

// GetStockInfoJq は指定された国の株式情報を取得
func (c *getStockInfoJqClientImpl) GetStockInfoJq(ctx context.Context, req *get_stock_info_jq.GetStockInfoJqRequest) (*get_stock_info_jq.GetStockInfoJqResponse, error) {
	return c.client.GetStockInfoJq(ctx, req)
}

// Close はgRPC接続を閉じます。
func (c *getStockInfoJqClientImpl) Close() error {
	return c.conn.Close()
}

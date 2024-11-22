// api-go\src\service\ms_gateway\client\get_stocks_datalist_with_dates_client.go

package client

import (
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GetStocksDatalistWithDatesClient は株価情報取得サービスのgRPCクライアント
type GetStocksDatalistWithDatesClient interface {
	GetStocksDatalist(ctx context.Context, req *gsdwd.GetStocksDatalistWithDatesRequest) (*gsdwd.GetStocksDatalistWithDatesResponse, error)
	Close() error
}

// getStocksDatalistWithDatesClientImpl は GetStocksDatalistWithDatesClient インターフェースの実装
type getStocksDatalistWithDatesClientImpl struct {
	client gsdwd.GetStocksDatalistWithDatesServiceClient
	conn   *grpc.ClientConn
}

// NewGetStocksDatalistWithDatesClient は GetStocksDatalistWithDatesClient の新しいインスタンスを作成
func NewGetStocksDatalistWithDatesClient(ctx context.Context) (GetStocksDatalistWithDatesClient, error) {
	port := os.Getenv("GET_STOCK_DATALIST_WITH_DATES_MS_PORT") // .envを確認
	address := fmt.Sprintf("localhost:%s", port)               // 環境変数からポート番号を取得

	// NewClient を使用して ClientConn を作成
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	client := gsdwd.NewGetStocksDatalistWithDatesServiceClient(conn)
	return &getStocksDatalistWithDatesClientImpl{client: client, conn: conn}, nil
}

// GetStocksDatalist は指定された銘柄コードと日付範囲の株価情報を取得
func (c *getStocksDatalistWithDatesClientImpl) GetStocksDatalist(ctx context.Context, req *gsdwd.GetStocksDatalistWithDatesRequest) (*gsdwd.GetStocksDatalistWithDatesResponse, error) {
	return c.client.GetStocksDatalist(ctx, req)
}

// Close はgRPC接続を閉じます。
func (c *getStocksDatalistWithDatesClientImpl) Close() error {
	return c.conn.Close()
}

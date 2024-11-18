// api-go\src\service\ms_gateway\client\simple_moving_average_client.go
package client

import (
	"context"
	"fmt"
	"os"
	"time"

	sma "api-go/src/service/ms_gateway/calculate_indicator/simple_moving_average"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// SimpleMovingAverageClient は単純移動平均計算サービスのgRPCクライアント
type SimpleMovingAverageClient interface {
	CalculateSimpleMovingAverage(ctx context.Context, req *sma.SimpleMovingAverageRequest) (*sma.SimpleMovingAverageResponse, error)
	Close() error
}

// simpleMovingAverageClientImpl は SimpleMovingAverageClient インターフェースの実装
type simpleMovingAverageClientImpl struct {
	client sma.SimpleMovingAverageServiceClient
	conn   *grpc.ClientConn
}

// NewSimpleMovingAverageClient は SimpleMovingAverageClient の新しいインスタンスを作成
func NewSimpleMovingAverageClient(ctx context.Context) (SimpleMovingAverageClient, error) {
	port := os.Getenv("SIMPLE_MOVING_AVERAGE_MS_PORT")
	address := fmt.Sprintf("localhost:%s", port)

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

	client := sma.NewSimpleMovingAverageServiceClient(conn)
	return &simpleMovingAverageClientImpl{client: client, conn: conn}, nil
}

// CalculateSimpleMovingAverage は指定された株価データと期間の単純移動平均を計算
func (c *simpleMovingAverageClientImpl) CalculateSimpleMovingAverage(ctx context.Context, req *sma.SimpleMovingAverageRequest) (*sma.SimpleMovingAverageResponse, error) {
	return c.client.CalculateSimpleMovingAverage(ctx, req)
}

// Close はgRPC接続を閉じる
func (c *simpleMovingAverageClientImpl) Close() error {
	return c.conn.Close()
}

// api-go\src\service\ms_gateway\client\generate_chart_client.go
package client

import (
	"context"
	"fmt"
	"os"
	"time"

	gc "api-go/src/service/ms_gateway/generate_chart" // generate_chart package の alias を gc に変更

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GenerateChartClient はチャート生成サービスのgRPCクライアント
type GenerateChartClient interface {
	GenerateChart(ctx context.Context, req *gc.GenerateChartRequest) (*gc.GenerateChartResponse, error)
	Close() error
}

// generateChartClientImpl は GenerateChartClient インターフェースの実装
type generateChartClientImpl struct {
	client gc.GenerateChartServiceClient
	conn   *grpc.ClientConn
}

// NewGenerateChartClient は GenerateChartClient の新しいインスタンスを作成
func NewGenerateChartClient(ctx context.Context) (GenerateChartClient, error) {
	port := os.Getenv("GENERATE_CHART_MS_PORT") // 環境変数名も修正
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
	for conn.GetState() != connectivity.Ready {
		if !conn.WaitForStateChange(connectCtx, conn.GetState()) {
			return nil, fmt.Errorf("failed to connect to gRPC server: %w", connectCtx.Err())
		}
	}

	client := gc.NewGenerateChartServiceClient(conn)
	return &generateChartClientImpl{client: client, conn: conn}, nil
}

// GenerateChart はチャート生成リクエストを送信し、レスポンスを受け取る
func (c *generateChartClientImpl) GenerateChart(ctx context.Context, req *gc.GenerateChartRequest) (*gc.GenerateChartResponse, error) {
	return c.client.GenerateChart(ctx, req)
}

// Close はgRPC接続を閉じる
func (c *generateChartClientImpl) Close() error {
	return c.conn.Close()
}

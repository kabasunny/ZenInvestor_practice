// api-go\src\service\ms_gateway\client\generate_chart_lc_sim_client.go
package client

import (
	generate_chart_lc_sim "api-go/src/service/ms_gateway/generate_chart_lc_sim"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GenerateChartLCClient はチャート生成サービスのgRPCクライアント
type GenerateChartLCClient interface {
	GenerateChart(ctx context.Context, req *generate_chart_lc_sim.GenerateChartLCRequest) (*generate_chart_lc_sim.GenerateChartLCResponse, error)
	Close() error
}

// generateChartLCClientImpl は GenerateChartLCClient インターフェースの実装
type generateChartLCClientImpl struct {
	client generate_chart_lc_sim.GenerateChartLCServiceClient
	conn   *grpc.ClientConn
}

// NewGenerateChartLCClient は GenerateChartLCClient の新しいインスタンスを作成
func NewGenerateChartLCClient(ctx context.Context) (GenerateChartLCClient, error) {
	port := os.Getenv("GENERATE_CHART_LC_SIM_MS_PORT") // .envを確認
	address := fmt.Sprintf("localhost:%s", port)       // 環境変数からポート番号を取得

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

	client := generate_chart_lc_sim.NewGenerateChartLCServiceClient(conn)
	return &generateChartLCClientImpl{client: client, conn: conn}, nil
}

// GenerateChart はチャート生成リクエストを送信
func (c *generateChartLCClientImpl) GenerateChart(ctx context.Context, req *generate_chart_lc_sim.GenerateChartLCRequest) (*generate_chart_lc_sim.GenerateChartLCResponse, error) {
	fmt.Println("In GenerateChart")

	response, err := c.client.GenerateChart(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Close はgRPC接続を閉じる
func (c *generateChartLCClientImpl) Close() error {
	return c.conn.Close()
}

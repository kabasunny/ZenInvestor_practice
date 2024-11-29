// api-go\src\service\ms_gateway\client\get_trading_calendar_jq_client.go

package client

import (
	get_trading_calendar_jq "api-go/src/service/ms_gateway/get_trading_calendar_jq"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// GetTradingCalendarJqClient は取引カレンダー取得サービスのgRPCクライアント
type GetTradingCalendarJqClient interface {
	GetTradingCalendarJq(ctx context.Context, req *get_trading_calendar_jq.GetTradingCalendarJqRequest) (*get_trading_calendar_jq.GetTradingCalendarJqResponse, error)
	Close() error
}

// getTradingCalendarJqClientImpl は GetTradingCalendarJqClient インターフェースの実装
type getTradingCalendarJqClientImpl struct {
	client get_trading_calendar_jq.GetTradingCalendarJqServiceClient
	conn   *grpc.ClientConn
}

// NewGetTradingCalendarJqClient は GetTradingCalendarJqClient の新しいインスタンスを作成
func NewGetTradingCalendarJqClient(ctx context.Context) (GetTradingCalendarJqClient, error) {
	port := os.Getenv("GET_TRADING_CALENDAR_JQ_MS_PORT") // .envを確認
	address := fmt.Sprintf("localhost:%s", port)         // 環境変数からポート番号を取得

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

	client := get_trading_calendar_jq.NewGetTradingCalendarJqServiceClient(conn)
	return &getTradingCalendarJqClientImpl{client: client, conn: conn}, nil
}

// GetTradingCalendarJq は取引カレンダーを取得
func (c *getTradingCalendarJqClientImpl) GetTradingCalendarJq(ctx context.Context, req *get_trading_calendar_jq.GetTradingCalendarJqRequest) (*get_trading_calendar_jq.GetTradingCalendarJqResponse, error) {
	fmt.Println("In GetTradingCalendarJq")
	return c.client.GetTradingCalendarJq(ctx, req)
}

// Close はgRPC接続を閉じる
func (c *getTradingCalendarJqClientImpl) Close() error {
	return c.conn.Close()
}

package infra

// MSClients holds the gRPC clients for microservices
type MSClients struct {
	MSClients map[string]interface{}
}

// アプリケーション用とバッチ処理用を分離し、ここは構造体のみ定義

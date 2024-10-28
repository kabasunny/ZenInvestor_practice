# MyProject

## インストール
1. protoc のダウンロード
https://github.com/protocolbuffers/protobuf/releases

2. zip ファイルの解凍
3. protoc.exe のコピー
C:\Windows
C:\Windows\System32
どちらかに管理者権限で貼り付け

4. 確認
コマンドプロンプトを開き
protoc --version


## protoc がGoコードを生成するためのプラグイン
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

## gRPCサービスのGoコードを生成するためのプラグイン
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## 依存パッケージのインストール
go mod init myproject
go get google.golang.org/grpc
go get google.golang.org/protobuf



## プロトコルバッファファイルのコンパイル
## 以下のコマンドを実行して、Go用のgRPCクライアントコードを生成：
## 株価データの取得
protoc --proto_path=../data-analysis-python/src/get_stock --go_out=./src/service/gateway --go_opt=paths=source_relative --go-grpc_out=./src/service/gateway --go-grpc_opt=paths=source_relative ../data-analysis-python/src/get_stock/get_stock.proto
## 指標：移動平均データの取得
protoc --proto_path=../data-analysis-python/src/calculate_indicator/moving_average --go_out=./src/service/gateway/moving_average --go_opt=paths=source_relative --go-grpc_out=./src/service/gateway/moving_average --go-grpc_opt=paths=source_relative ../data-analysis-python/src/calculate_indicator/moving_average/calculate_indicator.proto
## 可視化データの取得
protoc --proto_path=../data-analysis-python/src/generate_chart --go_out=./src/service/gateway --go_opt=paths=source_relative --go-grpc_out=./src/service/gateway --go-grpc_opt=paths=source_relative ../data-analysis-python/src/generate_chart/generate_chart.proto

├── src/
│   ├── controller/
│   │   └── controller.go
│   ├── service/
│   │   ├── service.go
│   │   └── gateway/  # マイクロサービスゲートウェイディレクトリ
│   │       └── grpc_client.go
│   └── repository/
│       └── repository.go
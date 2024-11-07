## テストディレクトリだけtest

# 一括でテスト
go test -v ./test/...


## 株価取得のテスト
# Python用のターミナル
cd .\data-analysis-python\
python src\get_stock_data\get_stock_data_grpc.py
# Go用の別のターミナル
cd .\api-go\
# クライアントテスト単体テスト
go test -v ./test/service/gateway/client/get_stock_data_client_test.go
# サービス層のテスト
go test -v ./test/service/get_stock_service_test.go

# 株価取得のクライアントテスト単体テスト
# Python用のターミナル
cd .\data-analysis-python\
python src\get_stock_data\get_stock_data_grpc.py
# Go用の別のターミナル
cd .\api-go\
go test -v ./test/service/gateway/client/simple_moving_average_client_test.go


VSコードで開いているファイルを閉じる
Windows用：Ctrl + K を押してから W を押す
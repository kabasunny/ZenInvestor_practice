# data-analysis-python\src\get_stock_data_with_dates\test_get_stock_data_with_dates_grpc.py
import unittest  # 標準的なテストフレームワークをインポート
import grpc
from concurrent import futures
import time
from get_stock_data_with_dates_grpc import (
    GetStockDataWithDatesService,
    serve,
)  # サービスとサーバー起動関数をインポート
import get_stock_data_with_dates_pb2  # Protocol Buffersコンパイラによって生成されるメッセージの定義を含むPythonモジュール
import get_stock_data_with_dates_pb2_grpc  # Protocol Buffersコンパイラによって生成されるgRPCサービスに関連するコードを含むPythonモジュール
from get_stock_data_with_dates_service import (
    get_stock_data_with_dates,
)  # 株価データ取得関数をインポート


class TestGetStockDataWithDatesGRPC(unittest.TestCase):  # テストクラスの定義

    @classmethod
    def setUpClass(cls):  # クラス全体で一度だけ実行されるセットアップメソッド
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_stock_data_with_dates_pb2_grpc.add_GetStockDataWithDatesServiceServicer_to_server(
            GetStockDataWithDatesService(), cls.server
        )
        cls.server.add_insecure_port("[::]:50104")
        cls.server.start()
        time.sleep(2)  # サーバーの起動待ち時間を追加

    @classmethod
    def tearDownClass(cls):  # クラス全体で一度だけ実行されるクリーンアップメソッド
        cls.server.stop(None)

    def test_get_stock_data_with_dates_grpc(self):  # 実際のテストメソッド
        print("gRPCサーバーtest")
        with grpc.insecure_channel(
            "localhost:50104"
        ) as channel:  # gRPCチャンネルを作成
            stub = get_stock_data_with_dates_pb2_grpc.GetStockDataWithDatesServiceStub(
                channel
            )  # スタブを作成
            start_date = "2023-01-01"
            end_date = "2023-01-31"
            response = stub.GetStockData(
                get_stock_data_with_dates_pb2.GetStockDataWithDatesRequest(
                    ticker="AAPL", start_date=start_date, end_date=end_date
                )
            )  # リクエストを送信
            print(f"Stock Name: {response.stock_name}")  # 銘柄名を表示
            print(
                "gRPCレスポンスは Protocol Buffers の形式:", response
            )  # レスポンスをターミナルに表示

            # 数値データの検証
            print("各日付のデータをループして日足データを表示")
            for date, stock_data in response.stock_data.items():
                print(
                    f"{date}: open: {stock_data.open} close: {stock_data.close} high: {stock_data.high} low: {stock_data.low} volume: {stock_data.volume}"
                )
                self.assertIsNotNone(stock_data.close)
                self.assertIsNotNone(stock_data.open)
                self.assertIsNotNone(stock_data.high)
                self.assertIsNotNone(stock_data.low)
                self.assertIsNotNone(stock_data.volume)


if __name__ == "__main__":  # スクリプトが直接実行された場合にテストを実行
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stock_data_with_dates -p 'test_get_stock_data_with_dates_grpc.py'

# test_get_stock_data_grpc.py
import unittest  # 標準的なテストフレームワークをインポート
import grpc
from concurrent import futures
from get_stock_data_grpc import (
    GetStockDataService,
    serve,
)  # サービスとサーバー起動関数をインポート
import get_stock_data_pb2 # Protocol Buffersコンパイラによって生成されるメッセージの定義を含むPythonモジュール
import get_stock_data_pb2_grpc # Protocol Buffersコンパイラによって生成されるgRPCサービスに関連するコードを含むPythonモジュール
from get_stock_data_service import get_stock_data  # 株価データ取得関数をインポート


class TestGetStockDataGRPC(unittest.TestCase):  # テストクラスの定義

    @classmethod
    def setUpClass(cls):  # クラス全体で一度だけ実行されるセットアップメソッド
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_stock_data_pb2_grpc.add_GetStockDataServiceServicer_to_server(
            GetStockDataService(), cls.server
        )
        cls.server.add_insecure_port("[::]:50051")
        cls.server.start()

    @classmethod
    def tearDownClass(cls):  # クラス全体で一度だけ実行されるクリーンアップメソッド
        cls.server.stop(None)

    def test_get_stock_data_grpc(self):  # 実際のテストメソッド
        print("gRPCサーバーtest")
        with grpc.insecure_channel(
            "localhost:50051"
        ) as channel:  # gRPCチャンネルを作成
            stub = get_stock_data_pb2_grpc.GetStockDataServiceStub(
                channel
            )  # スタブを作成
            period = "5d"
            response = stub.GetStockData(
                get_stock_data_pb2.GetStockDataRequest(ticker="^GSPC", period=period)
                # 期間の引数のリスト
                # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）
            )  # リクエストを送信
            print(
                "gRPCレスポンスは Protocol Buffers の形式:", response
            )  # レスポンスをターミナルに表示

            # 数値データの検証
            print("各日付のデータをループして" f"{period}分表示")
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
# python -m unittest discover -s src/get_stock_data  -p 'test_get_stock_data_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_stock_data  -p 'test*.py'

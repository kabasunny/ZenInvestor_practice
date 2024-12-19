# data-analysis-python\src\generate_chart_lc_sim\test_generate_chart_lc_sim_grpc.py
import unittest
import grpc
from concurrent import futures
import generate_chart_lc_sim_pb2
import generate_chart_lc_sim_pb2_grpc
from generate_chart_lc_sim_grpc import GenerateChartLCService
import pandas as pd
import os
import base64


class TestGenerateChartLCService(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        """gRPCサーバーをセットアップ"""
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        generate_chart_lc_sim_pb2_grpc.add_GenerateChartLCServiceServicer_to_server(
            GenerateChartLCService(), cls.server
        )
        cls.port = 50406
        cls.server.add_insecure_port(f"[::]:{cls.port}")
        cls.server.start()
        print(f"Test gRPC Server started, listening on port {cls.port}")

    @classmethod
    def tearDownClass(cls):
        """gRPCサーバーを停止"""
        cls.server.stop(None)
        print("Test gRPC Server stopped")

    def test_generate_chart(self):
        """gRPCサービスのGenerateChartメソッドをテスト"""
        # テストデータの作成
        data = {
            "Date": pd.date_range(start="2023-01-01", periods=10, freq="D"),
            "Close": [100 + i for i in range(10)],
        }

        dates = [date.strftime("%Y-%m-%d") for date in data["Date"]]
        close_prices = list(data["Close"])
        purchase_date = dates[2]
        purchase_price = close_prices[2]
        end_date = dates[7]
        end_price = close_prices[7]

        # gRPCリクエストメッセージの作成
        request = generate_chart_lc_sim_pb2.GenerateChartLCRequest(
            dates=dates,
            close_prices=close_prices,
            purchase_date=purchase_date,
            purchase_price=purchase_price,
            end_date=end_date,
            end_price=end_price,
        )

        # gRPCクライアントの設定
        with grpc.insecure_channel(f"localhost:{self.port}") as channel:
            stub = generate_chart_lc_sim_pb2_grpc.GenerateChartLCServiceStub(channel)
            response = stub.GenerateChart(request)

        # レスポンスの検証
        self.assertTrue(response.success)
        self.assertEqual(response.message, "チャート生成に成功しました")

        # 可視化データをoutput_testsディレクトリに保存
        output_dir = "src/generate_chart_lc_sim/output_tests"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        output_path = os.path.join(output_dir, "test_generate_chart_lc_sim_grpc.png")

        with open(output_path, "wb") as f:
            f.write(base64.b64decode(response.chart_data))

        self.assertTrue(os.path.exists(output_path))
        print(f"可視化データが {output_path} に保存されました")


if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/generate_chart_lc_sim -p 'test_generate_chart_lc_sim_grpc.py'

# 一括テスト
# python -m unittest discover -s src/generate_chart_lc_sim -p 'test*.py'

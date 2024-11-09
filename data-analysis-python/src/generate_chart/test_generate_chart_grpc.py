# test_generate_chart_grpc.py
import unittest
import grpc
from concurrent import futures
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_grpc import ChartGenerationService

class TestGenerateChartGRPC(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # gRPCサーバーのセットアップ
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        generate_chart_pb2_grpc.add_ChartGenerationServiceServicer_to_server(ChartGenerationService(), cls.server)
        cls.server.add_insecure_port('[::]:50052')
        cls.server.start()
        print("gRPCサーバー起動")

    @classmethod
    def tearDownClass(cls):
        # gRPCサーバーの停止
        cls.server.stop(None)
        print("gRPCサーバー停止")

    def test_generate_chart(self):
        # テストデータを作成
        stock_data = {
            "2023-01-01": generate_chart_pb2.StockDataForChart(
                open=100, close=110, high=115, low=95, volume=1000
            ),
            "2023-01-02": generate_chart_pb2.StockDataForChart(
                open=110, close=120, high=125, low=105, volume=1100
            )
        }

        indicators = [
            generate_chart_pb2.IndicatorData(
                type="Test Indicator",
                values={
                    "2023-01-01": 105,
                    "2023-01-02": 115
                }
            )
        ]

        with grpc.insecure_channel('localhost:50052') as channel:
            stub = generate_chart_pb2_grpc.ChartGenerationServiceStub(channel)
            request = generate_chart_pb2.GenerateChartRequest(
                stock_data=stock_data,
                indicators=indicators
            )
            response = stub.GenerateChart(request)
            chart_data = response.chart_data

            print("生成されたチャートデータ:", chart_data)

            self.assertTrue(chart_data)  # チャートデータが存在することを確認

if __name__ == '__main__':
    unittest.main()


# 本ファイル単体テスト
# python -m unittest discover -s src/generate_chart  -p 'test_generate_chart_grpc.py'

# 一括テスト
# python -m unittest discover -s src/generate_chart  -p 'test*.py'

import unittest
import grpc
from concurrent import futures
import calculate_indicator_pb2
import calculate_indicator_pb2_grpc
from calculate_indicator_grpc import IndicatorService
import matplotlib.pyplot as plt
import os

class TestCalculateIndicatorGRPC(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        calculate_indicator_pb2_grpc.add_IndicatorServiceServicer_to_server(IndicatorService(), cls.server)
        cls.server.add_insecure_port('[::]:50053')
        cls.server.start()

    @classmethod
    def tearDownClass(cls):
        cls.server.stop(None)

    def test_calculate_moving_average(self):
        with grpc.insecure_channel('localhost:50053') as channel:
            stub = calculate_indicator_pb2_grpc.IndicatorServiceStub(channel)
            response = stub.CalculateMovingAverage(calculate_indicator_pb2.IndicatorRequest(ticker="^GSPC", window_size=5))
            moving_average = response.moving_average

            # 移動平均のデータをプロット
            plt.figure()
            plt.plot(moving_average, label="^GSPC 5-day Moving Average")
            plt.title("^GSPC 5-day Moving Average")
            plt.legend()

            output_dir = "src/calculate_indicator/test_output"
            if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
                os.makedirs(output_dir)

            plt.savefig(f"{output_dir}/grpc_test_moving_average_chart.png")  # チャートを保存
            plt.close()

            self.assertTrue(moving_average)  # 移動平均データが存在することを確認
            print(f"Chart saved as {output_dir}/grpc_test_moving_average_chart.png")

if __name__ == "__main__":
    unittest.main()

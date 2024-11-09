import unittest
import grpc
from concurrent import futures
import simple_moving_average_pb2
import simple_moving_average_pb2_grpc
from simple_moving_average_grpc import SimpleMovingAverageService
import yfinance as yf
import matplotlib.pyplot as plt
import os

# 株価データを取得する関数
def fetch_stock_data(ticker):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="1y")
    # StockDataメッセージに適合するデータ形式に変換
    return {
        str(index.date()): simple_moving_average_pb2.StockData(
            open=row["Open"],
            close=row["Close"],
            high=row["High"],
            low=row["Low"],
            volume=row["Volume"]
        )
        for index, row in stock_data.iterrows()
    }

class SimpleMovingAverageServiceForTest(SimpleMovingAverageService):
    def CalculateSimpleMovingAverage(self, request, context):
        return super().CalculateSimpleMovingAverage(request, context)

class TestCalculateIndicatorGRPC(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        simple_moving_average_pb2_grpc.add_SimpleMovingAverageServiceServicer_to_server(
            SimpleMovingAverageServiceForTest(), cls.server
        )
        cls.server.add_insecure_port("[::]:50053")
        cls.server.start()

    @classmethod
    def tearDownClass(cls):
        cls.server.stop(None)

    def test_calculate_simple_moving_average(self):
        stock_data = fetch_stock_data("^GSPC")
        with grpc.insecure_channel("localhost:50053") as channel:
            stub = simple_moving_average_pb2_grpc.SimpleMovingAverageServiceStub(channel)
            # gRPCリクエストの送信
            response = stub.CalculateSimpleMovingAverage(
                simple_moving_average_pb2.SimpleMovingAverageRequest(
                    stock_data=stock_data,
                    window_size=5
                )
            )
            simple_moving_average = response.moving_average

            # プロットして画像を保存
            plt.figure()
            close_prices = [stock_data[date].close for date in stock_data]
            moving_averages = [simple_moving_average[date] for date in sorted(simple_moving_average.keys())]

            plt.plot(close_prices, label="^GSPC Close Prices")
            plt.plot(moving_averages, label="^GSPC 5-day Moving Average")
            plt.title("^GSPC 5-day Moving Average")
            plt.legend()

            output_dir = "src/calculate_indicator/simple_moving_average/test_output"
            if not os.path.exists(output_dir):
                os.makedirs(output_dir)
            plt.savefig(f"{output_dir}/grpc_simple_moving_average_chart.png")
            plt.close()

            self.assertTrue(moving_averages)
            print(f"Chart saved as {output_dir}/grpc_simple_moving_average_chart.png")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/calculate_indicator/simple_moving_average -p 'test_simple_moving_average_grpc.py'

# 一括テスト
# python -m unittest discover -s src/calculate_indicator/simple_moving_average -p 'test*.py'

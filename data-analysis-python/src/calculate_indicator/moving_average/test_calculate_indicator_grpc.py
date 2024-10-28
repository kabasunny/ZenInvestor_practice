import unittest
import grpc
from concurrent import futures
import calculate_indicator_pb2
import calculate_indicator_pb2_grpc
from calculate_indicator_grpc import MovingAverageService
import yfinance as yf
import matplotlib.pyplot as plt
import os

# 株価データを取得する関数
def fetch_stock_data(ticker):
    stock = yf.Ticker(ticker)  # Yahoo Finance から指定されたティッカーシンボルのデータを取得
    stock_data = stock.history(period="1y")  # 過去1年分の株価データを取得
    return stock_data["Close"].tolist()  # 終値のリストを返す

class TestCalculateIndicatorGRPC(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        # サーバーに MovingAverageService を追加
        calculate_indicator_pb2_grpc.add_MovingAverageServiceServicer_to_server(MovingAverageService(), cls.server)
        # サーバーにポート 50053 を追加
        cls.server.add_insecure_port('[::]:50053')
        # サーバーを起動
        cls.server.start()

    @classmethod
    def tearDownClass(cls):
        # サーバーを停止
        cls.server.stop(None)

    def test_calculate_moving_average(self):
        stock_data = fetch_stock_data("^GSPC")  # 株価データを取得
        with grpc.insecure_channel('localhost:50053') as channel:  # gRPC チャンネルを作成
            stub = calculate_indicator_pb2_grpc.MovingAverageServiceStub(channel)  # スタブを作成
            # gRPC リクエストを送信し、移動平均を計算
            response = stub.CalculateMovingAverage(calculate_indicator_pb2.MovingAverageRequest(stock_data=stock_data, window_size=5))
            moving_average = response.values  # 移動平均の結果を取得

            # プロットして画像を保存
            plt.figure()
            plt.plot(stock_data, label="^GSPC Close Prices")
            plt.plot(moving_average, label="^GSPC 5-day Moving Average")
            plt.title("^GSPC 5-day Moving Average")
            plt.legend()

            output_dir = "src/calculate_indicator/moving_average/test_output"
            if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
                os.makedirs(output_dir)

            plt.savefig(f"{output_dir}/grpc_moving_average_chart.png")  # チャートを保存
            plt.close()

            self.assertTrue(moving_average)  # 移動平均データが存在することを確認
            print(f"Chart saved as {output_dir}/grpc_moving_average_chart.png")

if __name__ == "__main__":
    unittest.main()


# python -m unittest discover -s src/calculate_indicator/moving_average  -p 'test*.py'
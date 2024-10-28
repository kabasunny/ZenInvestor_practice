import unittest
import grpc
from concurrent import futures
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_grpc import ChartService
import yfinance as yf
import base64
import matplotlib.pyplot as plt
import os
import matplotlib
matplotlib.use('Agg')# GUIバックエンドを使用せず、ファイル出力専用のバックエンドに変更


# 株価データを取得する関数
def fetch_stock_data(ticker):
    stock = yf.Ticker(ticker)  # Yahoo Finance から指定されたティッカーシンボルのデータを取得
    stock_data = stock.history(period="1y")  # 過去1年分の株価データを取得
    return stock_data["Close"].tolist()  # 終値のリストを返す

# gRPC テストクラス
class TestGenerateChartGRPC(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        # サーバーに ChartService を追加
        generate_chart_pb2_grpc.add_ChartServiceServicer_to_server(ChartService(), cls.server)
        # サーバーにポート 50052 を追加
        cls.server.add_insecure_port('[::]:50052')
        # サーバーを起動
        cls.server.start()

    @classmethod
    def tearDownClass(cls):
        # サーバーを停止
        cls.server.stop(None)

    def test_generate_chart(self):
        stock_data = fetch_stock_data("^GSPC")  # 株価データを取得
        indicator_data = fetch_stock_data("^GSPC")  # 仮に同じデータを指標データとする

        with grpc.insecure_channel('localhost:50052') as channel:  # gRPC チャンネルを作成
            stub = generate_chart_pb2_grpc.ChartServiceStub(channel)  # スタブを作成
            # gRPC リクエストを送信し、チャートを生成
            response = stub.GenerateChart(generate_chart_pb2.ChartRequest(stock_data=stock_data, indicator_data=indicator_data, ticker="^GSPC"))
            chart_data = base64.b64decode(response.chart_data)  # BASE64デコード

            output_dir = "src/generate_chart/test_output"
            if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
                os.makedirs(output_dir)

            # チャートデータをファイルに書き込む
            with open(f"{output_dir}/grpc_test_chart.png", "wb") as f:
                f.write(chart_data)
            print(f"Chart saved as {output_dir}/grpc_test_chart.png")

            self.assertTrue(chart_data)  # チャートデータが存在することを確認

if __name__ == "__main__":
    unittest.main()  # テストを実行


# サーバーを起動し、クライアントとして gRPC チャンネルを作成してリクエストを送信し、レスポンスとして返された BASE64 エンコードされたチャートデータをデコードしてファイルに保存
# python -m unittest discover -s src/generate_chart  -p 'test*.py'
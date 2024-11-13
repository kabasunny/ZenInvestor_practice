# test_generate_chart_grpc.py
import unittest
import os
import base64
import grpc
import threading
import time
import yfinance as yf

import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_grpc import GenerateChartService
from concurrent import futures

class TestGenerateChartGrpc(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # サーバーを別スレッドで起動
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        generate_chart_pb2_grpc.add_GenerateChartServiceServicer_to_server(
            GenerateChartService(), cls.server)
        cls.port = '51052' # test時は常時ポート+1000
        cls.server.add_insecure_port(f'[::]:{cls.port}')
        cls.server.start()
        print(f'Server started on port {cls.port}')
        time.sleep(1)  # サーバーが起動するまで待機

    @classmethod
    def tearDownClass(cls):
        cls.server.stop(0)
        print('Server stopped')

    def setUp(self):
        # gRPCクライアントを作成
        self.channel = grpc.insecure_channel(f'localhost:{self.port}')
        self.stub = generate_chart_pb2_grpc.GenerateChartServiceStub(self.channel)

        # 仮の株価データをyfinanceから取得
        ticker = "^GSPC"  # S&P 500のティッカーシンボル
        period = "1y"  # 1年分のデータを取得
        data = yf.download(ticker, period=period)

        # 株価データをプロトコルバッファの形式に変換
        self.stock_data = {}
        for date, row in data.iterrows():
            date_str = date.strftime("%Y-%m-%d")
            self.stock_data[date_str] = generate_chart_pb2.StockDataForChart(
                open=row['Open'],
                close=row['Close'],
                high=row['High'],
                low=row['Low'],
                volume=row['Volume']
            )

        # 指標データを作成
        self.indicators = []
        percentages = [-10, -20, 15]
        for percent in percentages:
            indicator_values = {}
            for date_str, stock_data_pb in self.stock_data.items():
                adjusted_value = stock_data_pb.close * (1 + percent / 100)
                indicator_values[date_str] = adjusted_value
            indicator = generate_chart_pb2.IndicatorData(
                type=f"Indicator_{percent}",
                values=indicator_values
            )
            self.indicators.append(indicator)

    def test_generate_chart_with_volume(self):
        # リクエストオブジェクトを作成（出来高を含める）
        request = generate_chart_pb2.GenerateChartRequest(
            stock_data=self.stock_data,
            indicators=self.indicators,
            include_volume=True
        )

        # サービスを呼び出し
        response = self.stub.GenerateChart(request)

        # チャートデータをデコードして画像として保存
        chart_data_base64 = response.chart_data
        chart_data = base64.b64decode(chart_data_base64)

        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)

        # チャートデータをファイルに書き込む
        with open(f"{output_dir}/grpc_generate_chart_with_volume.png", "wb") as f:
            f.write(chart_data)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/grpc_generate_chart_with_volume.png")

    def test_generate_chart_without_volume(self):
        # リクエストオブジェクトを作成（出来高を含めない）
        request = generate_chart_pb2.GenerateChartRequest(
            stock_data=self.stock_data,
            indicators=self.indicators,
            include_volume=False
        )

        # サービスを呼び出し
        response = self.stub.GenerateChart(request)

        # チャートデータをデコードして画像として保存
        chart_data_base64 = response.chart_data
        chart_data = base64.b64decode(chart_data_base64)

        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)

        # チャートデータをファイルに書き込む
        with open(f"{output_dir}/grpc_generate_chart_without_volume.png", "wb") as f:
            f.write(chart_data)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/grpc_generate_chart_without_volume.png")

if __name__ == '__main__':
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/generate_chart  -p 'test_generate_chart_grpc.py'

# 一括テスト
# python -m unittest discover -s src/generate_chart  -p 'test*.py'

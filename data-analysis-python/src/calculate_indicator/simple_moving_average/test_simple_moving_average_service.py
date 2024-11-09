# test_simple_moving_average_service.py
import unittest
from simple_moving_average_service import calculate_simple_moving_average
import yfinance as yf
import matplotlib.pyplot as plt
import os
from google.protobuf.timestamp_pb2 import Timestamp
import simple_moving_average_pb2

# 株価データを取得する関数
def fetch_stock_data(ticker):
    stock = yf.Ticker(ticker)  # Yahoo Finance から指定されたティッカーシンボルのデータを取得
    stock_data = stock.history(period="1y")  # 過去1年分の株価データを取得
    return stock_data.to_dict('index')  # 株価データを辞書形式で返す

# 株価データをprotobufメッセージに変換する関数
def convert_to_protobuf(stock_data):
    stock_data_pb = {}
    for date, data in stock_data.items():
        stock_data_pb[date.strftime('%Y-%m-%d')] = simple_moving_average_pb2.StockData(
            open=data["Open"], close=data["Close"], high=data["High"], low=data["Low"], volume=data["Volume"]
        )
    return stock_data_pb

# 単体テスト用クラス
class TestCalculateIndicatorService(unittest.TestCase):
    def test_calculate_simple_moving_average(self):
        ticker = "^GSPC"  # テストするティッカーシンボル
        window_size = 30  # 移動平均のウィンドウサイズ
        stock_data = fetch_stock_data(ticker)  # 株価データを取得
        stock_data_pb = convert_to_protobuf(stock_data)  # protobufメッセージに変換
        
        # サービス関数の呼び出し
        simple_moving_average = calculate_simple_moving_average(
            stock_data_pb, window_size
        )  # 移動平均を計算

        # プロットして画像を保存
        plt.figure()  # 新しい図を作成
        close_prices = [data.close for date, data in stock_data_pb.items()]
        moving_averages = [simple_moving_average[date] for date in sorted(simple_moving_average.keys())]

        plt.plot(close_prices, label=f"{ticker} Close Prices")  # 終値をプロット
        plt.plot(
            moving_averages,
            label=f"{ticker} {window_size}-day Simple Moving Average",
        )  # 移動平均をプロット
        plt.title(
            f"{ticker} {window_size}-day Simple Moving Average"
        )  # グラフのタイトルを設定
        plt.legend()  # 凡例を表示
        output_dir = "src/calculate_indicator/simple_moving_average/test_output"
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)
        plt.savefig(
            f"{output_dir}/service_simple_moving_average_chart.png"
        )  # チャートを保存
        plt.close()

        self.assertTrue(moving_averages)  # 移動平均データが存在することを確認
        print(f"Chart saved as {output_dir}/service_simple_moving_average_chart.png")

if __name__ == "__main__":
    unittest.main()  # 単体テストを実行

# 本ファイル単体テスト
# python -m unittest discover -s src/calculate_indicator/simple_moving_average -p 'test_simple_moving_average_service.py'

# 一括テスト
# python -m unittest discover -s src/calculate_indicator/simple_moving_average -p 'test*.py'

import unittest
from calculate_indicator_service import calculate_moving_average
import yfinance as yf
import matplotlib.pyplot as plt
import os

# 株価データを取得する関数
def fetch_stock_data(ticker):
    stock = yf.Ticker(ticker)  # Yahoo Finance から指定されたティッカーシンボルのデータを取得
    stock_data = stock.history(period="1y")  # 過去1年分の株価データを取得
    return stock_data["Close"].tolist()  # 終値のリストを返す

# 単体テスト用クラス
class TestCalculateIndicatorService(unittest.TestCase):
    def test_calculate_moving_average(self):
        ticker = "^GSPC"  # テストするティッカーシンボル
        window_size = 30  # 移動平均のウィンドウサイズ
        stock_data = fetch_stock_data(ticker)  # 株価データを取得
        moving_average = calculate_moving_average(stock_data, window_size)  # 移動平均を計算

        # プロットして画像を保存
        plt.figure()  # 新しい図を作成
        plt.plot(stock_data, label=f"{ticker} Close Prices")  # 終値をプロット
        plt.plot(moving_average, label=f"{ticker} {window_size}-day Moving Average")  # 移動平均をプロット
        plt.title(f"{ticker} {window_size}-day Moving Average")  # グラフのタイトルを設定
        plt.legend()  # 凡例を表示
        output_dir = "src/calculate_indicator/moving_average/test_output"
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)
        plt.savefig(f"{output_dir}/service_moving_average_chart.png")  # チャートを保存
        plt.close()

        self.assertTrue(moving_average)  # 移動平均データが存在することを確認
        print(f"Chart saved as {output_dir}/service_moving_average_chart.png")

if __name__ == "__main__":
    unittest.main()  # 単体テストを実行

# python -m unittest discover -s src/calculate_indicator/moving_average  -p 'test*.py'
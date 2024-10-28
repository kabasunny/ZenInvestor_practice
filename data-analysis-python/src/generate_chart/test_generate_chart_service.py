import unittest
from generate_chart_service import generate_chart
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

class TestGenerateChartService(unittest.TestCase):
    def test_generate_chart(self):
        ticker = "^GSPC"  # テストするティッカーシンボル
        stock_data = fetch_stock_data(ticker)  # 株価データを取得
        indicator_data = fetch_stock_data(ticker)  # 仮に同じデータを指標データとして使用

        # generate_chart 関数を呼び出してチャートデータを生成
        chart_data = generate_chart(stock_data, indicator_data, ticker)
        # BASE64 デコードしてバイナリデータに変換
        chart_data = base64.b64decode(chart_data)

        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)

        # チャートデータをファイルに書き込む
        with open(f"{output_dir}/service_test_chart.png", "wb") as f:
            f.write(chart_data)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/service_test_chart.png")

if __name__ == "__main__":
    unittest.main()  # テストを実行

# 直接 generate_chart 関数を呼び出してチャートデータを取得し、BASE64 デコードしてファイルに保存
# python -m unittest discover -s src/generate_chart  -p 'test*.py'
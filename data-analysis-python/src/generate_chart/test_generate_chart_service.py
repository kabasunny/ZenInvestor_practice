# test_generate_chart_service.py
import unittest
from generate_chart_service import generate_chart, handle_generate_chart_request
import yfinance as yf
import base64
import os
import matplotlib
matplotlib.use('Agg')  # GUIバックエンドを使用せず、ファイル出力専用のバックエンドに変更

class MockRequest:
    def __init__(self, stock_data, indicators):
        self.stock_data = stock_data
        self.indicators = indicators

class MockIndicatorData:
    def __init__(self, indicator_type, values):
        self.indicator_type = indicator_type
        self.values = values

# 株価データを取得する関数
def fetch_stock_data(ticker, period):
    stock = yf.Ticker(ticker)  # Yahoo Finance から指定されたティッカーシンボルのデータを取得
    stock_data = stock.history(period)  # 指定された期間の株価データを取得
    return {
        date.strftime("%Y-%m-%d"): {
            "open": row["Open"],
            "close": row["Close"],
            "high": row["High"],
            "low": row["Low"],
            "volume": row["Volume"]
        }
        for date, row in stock_data.iterrows()
    }

# テスト用指標データを10%減算して作成する関数
def create_test_indicator_data(stock_data, percentage):
    test_indicator_data = {
        date: values["close"] * (1 - percentage / 100)
        for date, values in stock_data.items()
    }
    return test_indicator_data

class TestGenerateChartService(unittest.TestCase):
    def test_generate_chart(self):
        ticker = "^GSPC"  # テストするティッカーシンボル
        period = "1y"  # 過去1年分の株価データを指定
        stock_data = fetch_stock_data(ticker, period)  # 株価データを取得
        test_indicator_data = create_test_indicator_data(stock_data, 10)  # テスト用指標データを10%減算して作成

        indicator_data = {
            "type": "Test Indicator",
            "values": test_indicator_data,
        }

        # generate_chart 関数を呼び出してチャートデータを生成
        chart_data = generate_chart(stock_data, [indicator_data])
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

    def test_handle_generate_chart_request(self):
        ticker = "^GSPC"
        period = "1y"
        stock_data = fetch_stock_data(ticker, period)

        request_stock_data = {
            date: MockStockData(
                open=values["open"],
                close=values["close"],
                high=values["high"],
                low=values["low"],
                volume=values["volume"]
            ) for date, values in stock_data.items()
        }

        test_indicator_data = create_test_indicator_data(stock_data, 10) # テスト用指標データを10%減算して作成

         # MockIndicatorDataを使用して指標データを作成
        indicator_data = MockIndicatorData("Test Indicator", test_indicator_data) #  変更点


        request = MockRequest(request_stock_data, [indicator_data]) #  変更点
        chart_data = handle_generate_chart_request(request)
        # BASE64 デコードしてバイナリデータに変換
        chart_data = base64.b64decode(chart_data)

        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)

        # チャートデータをファイルに書き込む
        with open(f"{output_dir}/handle_service_test_chart.png", "wb") as f:
            f.write(chart_data)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/handle_service_test_chart.png")

    def test_generate_chart(self):
        ticker = "^GSPC"  # テストするティッカーシンボル
        period = "1y"  # 過去1年分の株価データを指定
        stock_data = fetch_stock_data(ticker, period)  # 株価データを取得
        test_indicator_data = create_test_indicator_data(stock_data, 10)  # テスト用指標データを10%減算して作成

        # generate_chart 関数に渡す指標データの形式を修正
        indicator_data_list = [
            MockIndicatorData("Test Indicator", test_indicator_data)
        ]


        # generate_chart 関数を呼び出してチャートデータを生成
        chart_data_base64 = generate_chart(stock_data, indicator_data_list)  # 修正点

        chart_data = base64.b64decode(chart_data_base64)


        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        # ... (ファイル出力処理)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/service_test_chart.png")

class MockStockData:
    def __init__(self, open, close, high, low, volume):
        self.open = open
        self.close = close
        self.high = high
        self.low = low
        self.volume = volume

if __name__ == "__main__":
    unittest.main()  # テストを実行


# 本ファイル単体テスト
# python -m unittest discover -s src/generate_chart  -p 'test_generate_chart_service.py'

# 一括テスト
# python -m unittest discover -s src/generate_chart  -p 'test*.py'

# test_generate_chart_service.py
import unittest
import os
import base64
from generate_chart_service import handle_generate_chart_request
import generate_chart_pb2
import yfinance as yf


class TestGenerateChartService(unittest.TestCase):
    def setUp(self):
        # 仮の株価データをyfinanceから取得
        ticker = "^GSPC"  # S&P 500のティッカーシンボル
        period = "1y"  # 1年分のデータを取得
        data = yf.download(ticker, period=period)

        # 株価データをプロトコルバッファの形式に変換
        self.stock_data = {}
        for date, row in data.iterrows():
            date_str = date.strftime("%Y-%m-%d")
            self.stock_data[date_str] = generate_chart_pb2.StockDataForChart(
                open=row["Open"],
                close=row["Close"],
                high=row["High"],
                low=row["Low"],
                volume=row["Volume"],
            )

        # 指標データを作成（株価の終値を基に-10%、-20%、+15%した値）
        self.indicators = []
        percentages = [-10, -20, 15]
        for percent in percentages:
            indicator_values = {}
            for date_str, stock_data_pb in self.stock_data.items():
                adjusted_value = stock_data_pb.close * (1 + percent / 100)
                indicator_values[date_str] = adjusted_value
            indicator = generate_chart_pb2.IndicatorData(
                type=f"Indicator_{percent}", legend_name=f"Indicator_{percent}", values=indicator_values
            )
            self.indicators.append(indicator)

    def test_handle_generate_chart_request_with_volume(self):
        # リクエストオブジェクトを作成（出来高を含める）
        request = generate_chart_pb2.GenerateChartRequest(
            stock_data=self.stock_data, indicators=self.indicators, include_volume=True
        )

        # handle_generate_chart_requestを呼び出し
        chart_data_base64 = handle_generate_chart_request(request)

        # チャートデータをデコードしてPNGファイルとして保存
        chart_data = base64.b64decode(chart_data_base64)

        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)

        # チャートデータをファイルに書き込む
        with open(
            f"{output_dir}/service_generate_chart_with_volume.png", "wb"
        ) as f:
            f.write(chart_data)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/service_generate_chart_with_volume.png")

    def test_handle_generate_chart_request_without_volume(self):
        # リクエストオブジェクトを作成（出来高を含めない）
        request = generate_chart_pb2.GenerateChartRequest(
            stock_data=self.stock_data, indicators=self.indicators, include_volume=False
        )

        # handle_generate_chart_requestを呼び出し
        chart_data_base64 = handle_generate_chart_request(request)

        # チャートデータをデコードしてPNGファイルとして保存
        chart_data = base64.b64decode(chart_data_base64)

        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)

        # チャートデータをファイルに書き込む
        with open(
            f"{output_dir}/service_generate_chart_without_volume.png", "wb"
        ) as f:
            f.write(chart_data)

        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/service_generate_chart_without_volume.png")


if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/generate_chart  -p 'test_generate_chart_service.py'

# 一括テスト
# python -m unittest discover -s src/generate_chart  -p 'test*.py'

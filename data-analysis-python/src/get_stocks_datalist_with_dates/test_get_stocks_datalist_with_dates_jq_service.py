# data-analysis-python\src\get_stocks_datalist_with_dates\test_get_stocks_datalist_with_dates_jq_service.py
import unittest
import os
import pandas as pd

from get_stocks_datalist_with_dates_jq_service import get_stocks_datalist_with_dates_jq  # 関数名を変更

class TestGetStocksDatalistWithDatesJqService(unittest.TestCase):

    def test_get_stocks_datalist_with_dates_jq(self):
        """株価情報取得メソッドのテスト"""
        # CSVファイルから銘柄リストを読み込む
        # csv_file_path = "./src/get_stocks_datalist_with_dates/test_input/tickers1.csv"
        # if not os.path.exists(csv_file_path):
        #     self.fail(f"CSVファイルが存在しません: {csv_file_path}")

        # codes_df = pd.read_csv(csv_file_path)
        # codes = codes_df['Ticker'].tolist()  # 銘柄コードをリスト形式に変換
        # codes = [str(code) for code in codes]  # 銘柄コードを文字列に変換

        codes = ["131A", "6932"] # 適当な2銘柄   1311 131A で実験中
        print(codes)

        from_date = "2024-08-01"  # 開始日
        to_date = "2024-08-01"    # 終了日

        # 株価情報を取得
        stock_quotes = get_stocks_datalist_with_dates_jq(codes=codes, from_date=from_date, to_date=to_date)

        # テストの確認
        self.assertIsNotNone(stock_quotes)
        self.assertGreater(len(stock_quotes), 0, "No stock quotes were retrieved.")

        # 全データを一つのDataFrameにまとめる
        df = pd.DataFrame(stock_quotes)

        # 最初の行に銘柄コードを追加
        df.insert(0, 'Ticker', df.pop('symbol'))

        # CSV出力
        output_dir = "src/get_stocks_datalist_with_dates/test_output"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)

        output_file = os.path.join(output_dir, f"stock_quotes_{from_date}_{to_date}.csv")
        df.to_csv(output_file, index=False)
        self.assertTrue(os.path.exists(output_file))
        print(f"CSVファイルが保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates -p 'test_get_stocks_datalist_with_dates_jq_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates -p 'test*.py'

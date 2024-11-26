# data-analysis-python\src\get_stocks_datalist_with_dates_jq\test_get_stocks_datalist_with_dates_jq_service.py
import unittest
import os
import pandas as pd

from get_stocks_datalist_with_dates_jq_service import fetch_stock_quotes

class TestGetStocksDatalistWithDatesJqService(unittest.TestCase):

    def test_fetch_stock_quotes(self):
        """株価情報取得メソッドのテスト"""
        # CSVファイルから銘柄リストを読み込む
        csv_file_path = "./src/get_stocks_datalist_with_dates_jq/listed_tickers.csv"
        if not os.path.exists(csv_file_path):
            self.fail(f"CSVファイルが存在しません: {csv_file_path}")

        codes_df = pd.read_csv(csv_file_path)
        codes = codes_df['Ticker'].tolist()  # 銘柄コードをリスト形式に変換
        codes = [str(code) for code in codes]  # 銘柄コードを文字列に変換

        from_date = "2024-08-01"  # 開始日
        to_date = "2024-08-01"    # 終了日

        # 株価情報を取得
        stock_quotes = fetch_stock_quotes(codes=codes, from_date=from_date, to_date=to_date)

        # テストの確認
        self.assertIsNotNone(stock_quotes)
        self.assertGreater(len(stock_quotes), 0, "No stock quotes were retrieved.")
        all_data = []

        for quote in stock_quotes:
            if not quote or not quote.get("daily_quotes"):
                continue  # データがない場合をスキップ
            self.assertIn('daily_quotes', quote)
            self.assertGreater(len(quote['daily_quotes']), 0, f"No data found for code: {quote.get('code', 'unknown')}")
            for q in quote['daily_quotes']:
                self.assertIn('Date', q)
                self.assertIn('Code', q)
                self.assertIn('Open', q)
                self.assertIn('High', q)
                self.assertIn('Low', q)
                self.assertIn('Close', q)
            all_data.extend(quote['daily_quotes'])

        # 全データを一つのDataFrameにまとめる
        df = pd.DataFrame(all_data)

        # 最初の行に銘柄コードを追加
        df.insert(0, 'Ticker', df.pop('Code'))

        # CSV出力
        output_dir = "src/get_stocks_datalist_with_dates_jq/test_output"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)

        output_file = os.path.join(output_dir, f"stock_quotes_{from_date}_{to_date}.csv")
        df.to_csv(output_file, index=False)
        self.assertTrue(os.path.exists(output_file))
        print(f"CSVファイルが保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates_jq -p 'test_get_stocks_datalist_with_dates_jq_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates_jq -p 'test*.py'

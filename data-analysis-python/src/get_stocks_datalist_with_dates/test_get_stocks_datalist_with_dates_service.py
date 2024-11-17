# data-analysis-python\src\get_stocks_datalist_with_dates\test_get_stocks_datalist_with_dates_service.py
import unittest
import pandas as pd
import os
from get_stocks_datalist_with_dates_service import get_stocks_datalist_with_dates

class TestGetStocksDatalistWithDatesService(unittest.TestCase):
    def test_get_stocks_datalist_with_dates_japan(self):
        symbols = ["7203.T", "6758.T"]  # トヨタ自動車、ソニー
        start_date = "2024-11-9"
        end_date = "2024-11-15"
        self._test_get_stocks_datalist_with_dates(symbols, start_date, end_date, 'jp_stock_prices_with_dates.csv')

    def _test_get_stocks_datalist_with_dates(self, symbols, start_date, end_date, output_file_name):
        # 株価情報を取得
        stock_prices_list = get_stocks_datalist_with_dates(symbols, start_date, end_date)
        
        # 株価情報がリストであり、非空であることを確認
        self.assertIsInstance(stock_prices_list, list)
        self.assertTrue(len(stock_prices_list) > 0, "株価情報が取得できませんでした")
        
        # データフレームとして保存
        df = pd.DataFrame(stock_prices_list)
        output_dir = "src/get_stocks_datalist_with_dates/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, output_file_name)
        df.to_csv(output_file, index=False)
        
        print(f"株価情報がCSVファイルとして保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates -p 'test_get_stocks_datalist_with_dates_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates -p 'test*.py'

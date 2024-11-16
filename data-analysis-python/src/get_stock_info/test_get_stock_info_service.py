# data-analysis-python\src\get_stock_info\test_get_stock_info_service.py

import unittest
import pandas as pd
import os
from get_stock_info_service import get_stock_info

class TestGetStockInfoService(unittest.TestCase):
    def test_get_stock_info_japan(self):
        self._test_get_stock_info_by_country('japan', 'jp_stock_info.csv')

    def test_get_stock_info_us(self):
        self._test_get_stock_info_by_country('united states', 'us_stock_info.csv')

    def test_get_stock_info_germany(self):
        self._test_get_stock_info_by_country('germany', 'de_stock_info.csv')

    def _test_get_stock_info_by_country(self, country, output_file_name):
        # 株式情報を取得
        stock_info_list = get_stock_info(country)
        
        # 株式情報がリストであり、非空であることを確認
        self.assertIsInstance(stock_info_list, list)
        self.assertTrue(len(stock_info_list) > 0, f"{country} の株式情報が取得できませんでした")
        
        # データフレームとして保存
        df = pd.DataFrame(stock_info_list)
        output_dir = "src/get_stock_info/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, output_file_name)
        df.to_csv(output_file, index=False)
        
        print(f"{country} の株式情報がCSVファイルとして保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stock_info -p 'test_get_stock_info_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stock_info -p 'test*.py'

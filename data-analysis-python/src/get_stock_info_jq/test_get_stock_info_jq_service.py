# data-analysis-python\src\get_stock_info_jq\test_get_stock_info_jq_service.py
import unittest
import os
import pandas as pd
from get_stock_info_jq_service import fetch_stock_info

class TestGetStockInfoJqService(unittest.TestCase):

    def test_service_methods(self):
        """サービスメソッドのテスト"""
        # 株式情報を取得
        df = fetch_stock_info()
        self.assertIsNotNone(df)
        self.assertGreater(len(df), 0)
        self.assertEqual(df.columns.tolist(), ['ticker', 'name', 'sector', 'industry', 'date'])

        # CSV出力
        output_dir = "src/get_stock_info_jq/test_output"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        output_file = os.path.join(output_dir, "listed_companies.csv")
        df.to_csv(output_file, index=False)
        self.assertTrue(os.path.exists(output_file))
        print(f"CSVファイルが保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stock_info_jq -p 'test_get_stock_info_jq_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stock_info_jq -p 'test*.py'

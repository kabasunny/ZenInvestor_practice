# data-analysis-python\src\get_trading_calendar_jq\test_get_trading_calendar_jq.py
import unittest
import os
import pandas as pd
from get_trading_calendar_jq import fetch_trading_calendar

class TestGetTradingCalendarService(unittest.TestCase):

    def test_service_methods(self):
        """サービスメソッドのテスト"""
        from_date = "2023-01-01"
        to_date = "2023-12-31"

        # 取引カレンダーを取得
        df = fetch_trading_calendar(from_date, to_date)
        self.assertIsNotNone(df)
        self.assertGreater(len(df), 0)
        self.assertEqual(df.columns.tolist(), ['Date', 'HolidayDivision'])

        # CSV出力
        output_dir = "src/get_trading_calendar_jq/test_output"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        output_file = os.path.join(output_dir, "trading_calendar_jq.csv")
        df.to_csv(output_file, index=False)
        self.assertTrue(os.path.exists(output_file))
        print(f"CSVファイルが保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_trading_calendar -p 'test_get_trading_calendar.py'

# 一括テスト
# python -m unittest discover -s src/get_trading_calendar -p 'test*.py'

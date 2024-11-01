# tests/test_services/test_get_stock_data_service.py
import unittest
from get_stock_data_service import get_stock_data


class TestGetStockDataService(unittest.TestCase):
    def test_get_stock_data(self):
        print("サービスメソッドtest")
        data = get_stock_data("AAPL", "5d")
        # 期間の引数のリスト
        # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）
        print("取得データはPython の辞書形式:")
        for date, values in data.items():
            print(f"{date}: {values}")

        # サンプルの日付をチェック
        sample_date = list(data.keys())[0]
        self.assertIn("Close", data[sample_date])
        self.assertIn("Open", data[sample_date])
        self.assertIn("High", data[sample_date])
        self.assertIn("Low", data[sample_date])
        self.assertIn("Volume", data[sample_date])


if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stock_data  -p 'test_get_stock_data_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stock_data  -p 'test*.py'

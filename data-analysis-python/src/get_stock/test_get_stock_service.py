# tests/test_services/test_get_stock_service.py
import unittest
from get_stock_service import get_stock_data


class TestGetStockService(unittest.TestCase):
    def test_get_stock_data(self):
        print("サービスメソッドtest")
        data, date = get_stock_data("AAPL")
        print(f"取得データはPython の辞書形式: {data}, 日付: {date}")
        self.assertIn("Close", data)  # 辞書部分を確認
        self.assertIn("Open", data)
        self.assertIn("High", data)
        self.assertIn("Low", data)
        self.assertIn("Volume", data)


if __name__ == "__main__":
    unittest.main()

# python -m unittest discover -s src/get_stock  -p 'test*.py'

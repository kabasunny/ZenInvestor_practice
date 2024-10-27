# tests/test_services/test_get_stock_service.py
import unittest
from get_stock_service import get_stock_data

class TestGetStockService(unittest.TestCase):
    def test_get_stock_data(self):
        data = get_stock_data("^GSPC")
        print("取得データはPython の辞書形式:", data)  # データをターミナルに表示
        self.assertIn("Close", data)
        self.assertGreater(data["Close"], 0) 

if __name__ == "__main__":
    unittest.main()

# python -m unittest discover -s src/get_stock  -p 'test*.py'
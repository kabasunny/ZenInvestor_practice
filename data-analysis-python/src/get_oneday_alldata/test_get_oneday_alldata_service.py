# test_get_oneday_alldata_service.py
import unittest
from get_oneday_alldata_service import get_oneday_alldata

class TestGetOneDayAllDataService(unittest.TestCase):
    def test_get_oneday_alldata(self):
        print("Testing get_oneday_alldata service method")
        # Specify a date for which the market was open
        date_str = "2023-10-05"  # Update to a recent trading date
        data = get_oneday_alldata(date_str)
        self.assertIsNotNone(data)
        self.assertTrue(len(data) > 0)
        
        # Print the retrieved data for verification
        for ticker, values in data.items():
            print(f"Ticker: {ticker}")
            print(f"Data: {values}")
            # Verify that all expected fields are present
            self.assertIn('open', values)
            self.assertIn('close', values)
            self.assertIn('high', values)
            self.assertIn('low', values)
            self.assertIn('volume', values)
            # Ensure that none of the values are None
            self.assertIsNotNone(values['open'])
            self.assertIsNotNone(values['close'])
            self.assertIsNotNone(values['high'])
            self.assertIsNotNone(values['low'])
            self.assertIsNotNone(values['volume'])

if __name__ == "__main__":
    unittest.main()


# 本ファイル単体テスト
# python -m unittest discover -s src/get_stock_data  -p 'test_get_oneday_alldata_service.py'

# 一括テスト
# python -m unittest discover -s src/get_stock_data  -p 'test*.py'
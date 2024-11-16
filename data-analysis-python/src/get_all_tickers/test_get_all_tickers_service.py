import unittest
import pandas as pd
import os
from get_all_tickers_service import get_all_tickers

class TestGetAllTickersService(unittest.TestCase):
    def test_get_all_tickers(self):
        # ティッカーシンボルを取得
        tickers = get_all_tickers()
        
        # ティッカーシンボルがリストであり、非空であることを確認
        self.assertIsInstance(tickers, list)
        self.assertTrue(len(tickers) > 0, "ティッカーシンボルが取得できませんでした")
        
        # データフレームとして保存
        df = pd.DataFrame(tickers, columns=['Ticker'])
        output_dir = "src/get_all_tickers/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, "all_tickers.csv")
        df.to_csv(output_file, index=False)
        
        print(f"ティッカーシンボルがCSVファイルとして保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_all_tickers -p 'test_get_all_tickers_service.py'

# 一括テスト
# python -m unittest discover -s src/get_all_tickers -p 'test*.py'

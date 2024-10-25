# tests/test_get_stock_data.py
import unittest
from flask import Flask
from src.routes.get_stock_data import bp
import time  # timeモジュールをインポート

app = Flask(__name__)
app.register_blueprint(bp)


class TestGetStockData(unittest.TestCase):
    def test_get_stock_data(self):
        with app.test_client() as client:  # 自動でクローズされる。通常の構文だとクローズメソッドが必要
            start_time = time.time()  # 測定開始時間
            response = client.post("/get-stock-data/", json={"ticker": "^GSPC"})
            end_time = time.time()  # 測定終了時間

            self.assertEqual(response.status_code, 200)
            data = response.get_json()
            # print(data)  # 取得したデータを表示
            self.assertIn("Close", data)
            self.assertGreater(len(data["Close"]), 0)

            elapsed_time = end_time - start_time  # 経過時間を計算
            print(
                f"Request-Response Time: {elapsed_time:.4f} seconds"
            )  # 経過時間を表示


if __name__ == "__main__":
    unittest.main()

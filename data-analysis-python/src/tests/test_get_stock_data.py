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

# class TestExample(unittest.TestCase):
#     def test_example(self):
#         self.assertEqual(1, 1)  # a と b が等しいことを確認
#         self.assertNotEqual(1, 2)  # a と b が等しくないことを確認
#         self.assertTrue(True)  # x が True であることを確認
#         self.assertFalse(False)  # x が False であることを確認
#         self.assertIs(None, None)  # a と b が同一のオブジェクトであることを確認
#         self.assertIsNot(None, 1)  # a と b が同一のオブジェクトでないことを確認
#         self.assertIsNone(None)  # x が None であることを確認
#         self.assertIsNotNone(1)  # x が None でないことを確認
#         self.assertIn(1, [1, 2, 3])  # a が b に含まれていることを確認
#         self.assertNotIn(4, [1, 2, 3])  # a が b に含まれていないことを確認
#         self.assertAlmostEqual(1.0001, 1.0002, places=3)  # a と b がほぼ等しいことを確認（少数点以下の桁数指定可能）
#         self.assertNotAlmostEqual(1.0001, 1.0002, places=4)  # a と b がほぼ等しくないことを確認
#         self.assertGreater(2, 1)  # a が b より大きいことを確認
#         self.assertGreaterEqual(2, 2)  # a が b 以上であることを確認
#         self.assertLess(1, 2)  # a が b より小さいことを確認
#         self.assertLessEqual(2, 2)  # a が b 以下であることを確認

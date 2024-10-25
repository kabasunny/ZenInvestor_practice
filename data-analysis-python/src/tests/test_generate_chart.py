# tests/test_generate_chart.py
import unittest
from flask import Flask
from src.routes.generate_chart import bp
import base64  # 追加
import os  # 追加

app = Flask(__name__)
app.register_blueprint(bp)

class TestGenerateChart(unittest.TestCase):
    def test_generate_chart(self):
        with app.test_client() as client:
            response = client.post("/generate-chart/", json={"ticker": "^GSPC"})
            self.assertEqual(response.status_code, 200)
            data = response.get_json()
            self.assertIn("chart_url", data)

            # 生成されたチャートをデコードして保存
            chart_data_base64 = data["chart_url"]
            chart_path = os.path.join(os.path.dirname(__file__), "chart.png")  # テストディレクトリに保存
            with open(chart_path, "wb") as f:
                f.write(base64.b64decode(chart_data_base64))
            print(f"Chart saved as {chart_path}")

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

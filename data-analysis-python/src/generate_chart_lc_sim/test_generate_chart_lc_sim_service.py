# data-analysis-python\src\generate_chart_lc_sim\test_generate_chart_lc_sim_service.py
import unittest
import pandas as pd
from datetime import datetime
import os
import base64  # base64 モジュールをインポート
from generate_chart_lc_sim_service import plot_stop_results


class TestGenerateChartLCService(unittest.TestCase):
    def test_plot_stop_results(self):
        print("plot_stop_results 関数のテスト")

        # テストデータの作成
        data = {
            "Date": pd.date_range(start="2023-01-01", periods=10, freq="D"),
            "Close": [100 + i for i in range(10)],
        }
        df = pd.DataFrame(data)

        # テスト用の購入日、購入価格、売却日、売却価格を設定
        purchase_date = df["Date"].iloc[2]
        purchase_price = df["Close"].iloc[2]
        end_date = df["Date"].iloc[7]
        end_price = df["Close"].iloc[7]

        # プロット関数の呼び出し
        image_base64 = plot_stop_results(
            df, purchase_date, purchase_price, end_date, end_price
        )

        # 可視化データをoutput_testsディレクトリに保存
        output_dir = "src/generate_chart_lc_sim/output_tests"
        os.makedirs(output_dir, exist_ok=True)
        output_path = os.path.join(output_dir, "test_generate_chart_lc_sim_service.png")

        with open(output_path, "wb") as f:
            f.write(base64.b64decode(image_base64))

        print(f"可視化データが {output_path} に保存されました")

        # 検証
        self.assertEqual(purchase_price, 102)
        self.assertEqual(end_price, 107)


if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/generate_chart_lc_sim -p 'test_generate_chart_lc_sim_service.py'

# 一括テスト
# python -m unittest discover -s src/generate_chart_lc_sim -p 'test*.py'

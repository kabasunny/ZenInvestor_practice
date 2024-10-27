import unittest  # 標準的なテストフレームワークをインポート
from generate_chart_service import generate_chart  # チャート生成サービスをインポート
import base64  # base64 モジュールをインポート
import os  # ファイル操作用のモジュールをインポート

class TestGenerateChartService(unittest.TestCase):
    def test_generate_chart(self):
        chart_data = generate_chart("^GSPC")  # チャートデータを生成
        chart_data = base64.b64decode(chart_data)  # base64 モジュールを使用してデコード
        output_dir = "src/generate_chart/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)
        with open(f"{output_dir}/service_test_chart.png", "wb") as f:
            f.write(chart_data)  # チャートデータをファイルに書き込む
        self.assertTrue(chart_data)  # チャートデータが存在することを確認
        print(f"Chart saved as {output_dir}/service_test_chart.png")

if __name__ == "__main__":
    unittest.main()  # テストを実行

# 直接 generate_chart 関数を呼び出してチャートデータを取得し、BASE64 デコードしてファイルに保存
# python -m unittest discover -s src/generate_chart  -p 'test*.py'
import unittest
from calculate_indicator_service import calculate_moving_average
import matplotlib.pyplot as plt
import os

class TestCalculateIndicatorService(unittest.TestCase):
    def test_calculate_moving_average(self):
        ticker = "^GSPC"
        window_size = 30
        moving_average = calculate_moving_average(ticker, window_size)
        
        # 移動平均のデータをプロット
        plt.figure()
        plt.plot(moving_average, label=f"{ticker} {window_size}-day Moving Average")
        plt.title(f"{ticker} {window_size}-day Moving Average")
        plt.legend()
        
        output_dir = "src/calculate_indicator/test_output"
        if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
            os.makedirs(output_dir)
        
        plt.savefig(f"{output_dir}/setvice_test_moving_average_chart.png")  # チャートを保存
        plt.close()
        
        self.assertTrue(moving_average)  # 移動平均データが存在することを確認
        print(f"Chart saved as {output_dir}/setvice_test_moving_average_chart.png")

if __name__ == "__main__":
    unittest.main()

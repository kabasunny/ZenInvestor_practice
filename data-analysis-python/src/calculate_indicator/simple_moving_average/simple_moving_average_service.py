import numpy as np

def calculate_simple_moving_average(stock_data, window_size):
    # 株価データ (close) を取り出し、日付順にソート
    sorted_dates = sorted(stock_data.keys())
    close_prices = [stock_data[date].close for date in sorted_dates]
    
    # NumPy 配列に変換
    close_prices = np.array(close_prices)

    # 移動平均を計算
    simple_moving_average = np.convolve(
        close_prices, np.ones(window_size) / window_size, mode="valid"
    )

    # 結果を日付順のマップに変換して返す
    sma_map = {}
    for i in range(len(simple_moving_average)):
        date = sorted_dates[i + window_size - 1]  # 移動平均の結果に対応する日付
        sma_map[date] = simple_moving_average[i]

    return sma_map

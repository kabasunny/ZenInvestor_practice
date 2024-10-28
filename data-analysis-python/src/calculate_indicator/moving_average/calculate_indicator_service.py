import numpy as np

def calculate_moving_average(stock_data, window_size):
    # 入力された株価データをNumPy配列に変換
    stock_data = np.array(stock_data)
    
    # 移動平均を計算するために畳み込み演算を使用
    # np.ones(window_size)/window_size は指定されたウィンドウサイズで平均化するためのカーネルを作成
    moving_average = np.convolve(stock_data, np.ones(window_size)/window_size, mode='valid')
    
    # 結果の配列をリスト形式に変換して返す
    return moving_average.tolist()

# data-analysis-python\src\check_symbols\check_symbols_service.py
import yfinance as yf
import pandas as pd

# CSVファイルから銘柄リストを読み込む
csv_file_path = "./src/check_symbols/jp_tickers_2.csv"
df = pd.read_csv(csv_file_path)
symbols = df['Ticker'].astype(str) + ".T"  # 'Ticker' カラムの末尾に ".T" を追加

# データをダウンロード
data = yf.download(symbols.tolist(), start="2024-11-19", end="2024-11-22", group_by='ticker')
print(data)


# 実行コマンド
# python ./src/check_symbols/check_symbols_service.py
# data-analysis-python\src\get_stock_info\investpy_yfinance_test\investpy_test.py
import investpy

# 日本の株式情報を取得
stocks = investpy.stocks.get_stocks(country='japan')

# 100番目の銘柄の全情報を取得
if len(stocks) >= 100:
    stock_100 = stocks.iloc[99]  # インデックスは0から始まるので99を指定
    print(stock_100)
else:
    print("日本の株式情報のデータが100件未満です。")

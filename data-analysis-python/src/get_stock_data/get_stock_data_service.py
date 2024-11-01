# services/get_stock_service.py
import yfinance as yf
import pandas as pd


def get_stock_data(ticker, period):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period=period)  # 指定された期間のデータを取得
    # stock_data = stock.history(period="1mo")  # 直近1ヶ月のデータを取得
    # 期間の引数のリスト
    # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）

    stock_dict = {}  # 結果を格納する辞書を初期化
    for date, row in stock_data.iterrows():
        stock_dict[date.strftime("%Y-%m-%d")] = {
            "Open": row["Open"],  # 始値
            "Close": row["Close"],  # 終値
            "High": row["High"],  # 高値
            "Low": row["Low"],  # 安値
            "Volume": row["Volume"],  # 取引量
        }

    return stock_dict  # すべての日付のデータを返す

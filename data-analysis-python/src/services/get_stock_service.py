# services/get_stock_service.py
import yfinance as yf
import pandas as pd


def get_stock_data(ticker):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="10y")  # 直近1年のデータを取得
    # stock_data = stock.history(period="1mo")  # 直近1ヶ月のデータを取得
    # 期間の引数のリスト
    # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）

    # JSONシリアライズ可能な形式に変換
    stock_data.reset_index(
        inplace=True
    )  # インデックスをリセットして、Timestampを列に変換
    stock_data["Date"] = stock_data["Date"].astype(str)  # 日付を文字列に変換
    stock_dict = stock_data.to_dict(orient="list")  # リスト形式の辞書に変換
    return stock_dict

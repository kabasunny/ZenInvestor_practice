# services/get_stock_service.py
import yfinance as yf
import pandas as pd


def get_stock_data(ticker):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="1mo")  # 直近1ヶ月のデータを取得

    # JSONシリアライズ可能な形式に変換
    stock_data.reset_index(
        inplace=True
    )  # インデックスをリセットして、Timestampを列に変換
    stock_data["Date"] = stock_data["Date"].astype(str)  # 日付を文字列に変換
    stock_dict = stock_data.to_dict(orient="list")  # リスト形式の辞書に変換
    return stock_dict

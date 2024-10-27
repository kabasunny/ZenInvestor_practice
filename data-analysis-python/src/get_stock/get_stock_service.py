# services/get_stock_service.py
import yfinance as yf
import pandas as pd


def get_stock_data(ticker):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="1y")  # 直近～のデータを取得
    # stock_data = stock.history(period="1mo")  # 直近1ヶ月のデータを取得
    # 期間の引数のリスト
    # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）

    # JSONシリアライズ可能な形式に変換
    stock_dict = {
        "Open": stock_data["Open"].iloc[-1],
        "Close": stock_data["Close"].iloc[-1]
    }
    return stock_dict

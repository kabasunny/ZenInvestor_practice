import yfinance as yf
import numpy as np

def calculate_moving_average(ticker, window_size):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="1y")["Close"]
    moving_average = stock_data.rolling(window=window_size).mean().dropna()
    return moving_average.tolist()

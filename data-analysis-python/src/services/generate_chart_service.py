# services/generate_chart_service.py
import matplotlib.pyplot as plt
import io
import base64
import yfinance as yf


def generate_chart(ticker):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="1y")  # 直近～のデータを取得
    # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）

    plt.figure()
    plt.plot(stock_data["Close"], label=f"{ticker} Close Prices")
    plt.title(f"{ticker} Stock Data")
    plt.legend()

    img = io.BytesIO()
    plt.savefig(img, format="png")
    img.seek(0)
    return base64.b64encode(img.getvalue()).decode()

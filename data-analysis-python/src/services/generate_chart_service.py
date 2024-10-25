# services/generate_chart_service.py
import matplotlib.pyplot as plt
import io
import base64
import yfinance as yf


def generate_chart(ticker):
    stock = yf.Ticker(ticker)
    stock_data = stock.history(period="1mo")  # 直近1ヶ月のデータを取得

    plt.figure()
    plt.plot(stock_data["Close"], label=f"{ticker} Close Prices")
    plt.title(f"{ticker} Stock Data")
    plt.legend()

    img = io.BytesIO()
    plt.savefig(img, format="png")
    img.seek(0)
    return base64.b64encode(img.getvalue()).decode()

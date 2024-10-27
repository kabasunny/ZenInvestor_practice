import matplotlib.pyplot as plt
import io
import base64
import yfinance as yf

def generate_chart(ticker):
    stock = yf.Ticker(ticker)  # Yahoo Financeから指定されたティッカーシンボルのデータを取得
    stock_data = stock.history(period="1y")  # 過去1年分の株価データを取得# 期間の引数のリスト
    # "1d": 1日, "5d": 5日, "1mo": 1ヶ月, "3mo": 3ヶ月, "6mo": 6ヶ月, "1y": 1年, "2y": 2年, "5y": 5年. "10y": 10年, "ytd": 年初から現在まで, "max": 最大期間（可能な限り最長）

    
    plt.figure()  # 新しい図を作成
    plt.plot(stock_data["Close"], label=f"{ticker} Close Prices")  # 株価データの終値をプロット
    plt.title(f"{ticker} Stock Data")  # グラフのタイトルを設定
    plt.legend()  # 凡例を表示
    
    img = io.BytesIO()  # バイトストリームのオブジェクトを作成
    plt.savefig(img, format="png")  # 図をPNG形式でバイトストリームに保存
    img.seek(0)  # バイトストリームの先頭に移動
    
    return base64.b64encode(img.getvalue()).decode()  # バイナリデータをBASE64エンコードして返す
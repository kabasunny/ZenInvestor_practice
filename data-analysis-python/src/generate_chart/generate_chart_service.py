import matplotlib.pyplot as plt
import io
import base64


def generate_chart(stock_data, indicator_data, ticker):
    # 新しい図を作成
    plt.figure()
    # 株価データをプロット
    plt.plot(stock_data, label=f"{ticker} Close Prices")
    # 指標データをプロット
    plt.plot(indicator_data, label=f"{ticker} Indicator", linestyle="--")
    # グラフのタイトルを設定
    plt.title(f"{ticker} Stock Data with Indicator")
    # 凡例を表示
    plt.legend()

    # 画像をバイトストリームに保存
    img = io.BytesIO()
    plt.savefig(img, format="png")
    img.seek(0)

    # 画像データをBase64エンコードして返す
    return base64.b64encode(img.getvalue()).decode()

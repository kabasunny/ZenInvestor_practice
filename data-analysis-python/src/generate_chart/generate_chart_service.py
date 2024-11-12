import matplotlib  # matplotlib をインポート (pyplotの前に)

# GUIバックエンドを使用しないように設定 (matplotlib のインポート後)
matplotlib.use("Agg")

import os
import matplotlib.pyplot as plt  # matplotlib.pyplot をインポート (matplotlibの後)
import base64
from io import BytesIO
from datetime import datetime
import matplotlib.font_manager as fm

# フォントへの相対パス
font_dir = os.path.join(os.path.dirname(__file__), "fonts")
font_path = os.path.join(font_dir, "ipaexg.ttf")
jp_font = fm.FontProperties(fname=font_path)


def generate_chart(stock_data, indicators_data):
    # stock_dataのキーはdatetimeオブジェクト
    sorted_dates = sorted(stock_data.keys())
    close_prices = [stock_data[date]["close"] for date in sorted_dates]

    plt.figure(figsize=(16, 9), facecolor="#ADD8E6")  # 図全体の背景色を水色に設定
    ax = plt.gca()
    ax.set_facecolor("#EAEAEA")  # プロット部分の背景色を灰色に設定
    plt.plot(sorted_dates, close_prices, label="終値")

    for indicator_data in indicators_data:
        legend_name = indicator_data.get(
            "legend_name", indicator_data["type"]
        )  # 凡例の名称を取得

        indicator_values = indicator_data["values"]

        # 日付をソートすることを確認
        sorted_indicator_dates = sorted(indicator_values.keys())
        values = [indicator_values[date] for date in sorted_indicator_dates]

        plt.plot(sorted_indicator_dates, values, label=legend_name)

    plt.xlabel("日付", fontproperties=jp_font)
    plt.ylabel("$ or \\", fontproperties=jp_font)
    plt.legend(
        loc="upper left", frameon=True, facecolor="#ADD8E6", prop=jp_font
    )  # 凡例を左上に表示し、背景色を水色に設定

    # グリッド線の色を白に設定
    plt.grid(color="white")

    # 余白を調整して16:9の描画部分を確保
    plt.tight_layout(rect=[0, 0, 1, 1])

    # プロットを画像として保存
    buffer = BytesIO()
    plt.savefig(buffer, format="png")
    buffer.seek(0)
    plt.close()  # 図を閉じる

    # 画像をBase64エンコード
    image_base64 = base64.b64encode(buffer.read()).decode("utf-8")

    return image_base64


def handle_generate_chart_request(request):
    # リクエストからstock_dataを処理
    stock_data = {}
    for date_str, stock_data_pb in request.stock_data.items():
        # 日付文字列をdatetimeオブジェクトに変換
        date = datetime.strptime(date_str, "%Y-%m-%d")
        stock_data[date] = {
            "open": stock_data_pb.open,
            "close": stock_data_pb.close,
            "high": stock_data_pb.high,
            "low": stock_data_pb.low,
            "volume": stock_data_pb.volume,
        }

    # リクエストから指標データを処理
    indicators_data = []
    for indicator_pb in request.indicators:
        indicator_type = indicator_pb.type
        legend_name = f"{indicator_pb.legend_name}"  # WindowSizeを凡例に含める
        indicator_values_pb = indicator_pb.values

        # 日付文字列をdatetimeオブジェクトに変換
        indicator_values = {}
        for date_str, value in indicator_values_pb.items():
            date = datetime.strptime(date_str, "%Y-%m-%d")
            indicator_values[date] = value

        indicators_data.append(
            {
                "type": indicator_type,
                "legend_name": legend_name,  # 新しいフィールドを追加
                "values": indicator_values,
            }
        )

    return generate_chart(stock_data, indicators_data)

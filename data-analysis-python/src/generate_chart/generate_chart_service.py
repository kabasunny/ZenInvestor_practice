# data-analysis-python\src\generate_chart\generate_chart_service.py
import matplotlib  # matplotlib をインポート (pyplotの前に)

# GUIバックエンドを使用しないように設定 (matplotlib のインポート後)
matplotlib.use("Agg")

import os
import matplotlib.pyplot as plt  # matplotlib.pyplot をインポート (matplotlibの後)
import base64
from io import BytesIO
from datetime import datetime
import matplotlib.font_manager as fm
import itertools  # itertools をインポート

# フォントへの相対パス
font_dir = os.path.join(os.path.dirname(__file__), "fonts")
font_path = os.path.join(font_dir, "ipaexg.ttf")
jp_font = fm.FontProperties(fname=font_path)

def generate_chart(stock_data, indicators_data, include_volume):
    # 株価データの日付をソートして取得
    sorted_dates = sorted(stock_data.keys())
    # 終値データを取得
    close_prices = [stock_data[date]["close"] for date in sorted_dates]

    # プロットの作成と背景色の設定
    fig, ax1 = plt.subplots(figsize=(16, 9), facecolor="#ADD8E6")
    ax1.set_facecolor("#EAEAEA")

    # 株価チャートのプロット（青色）
    ax1.plot(sorted_dates, close_prices, label="終値", color='tab:blue')
    ax1.set_xlabel("日付", fontproperties=jp_font)
    ax1.set_ylabel("価格", fontproperties=jp_font)

    # 出来高チャートのプロット
    if include_volume:
        volumes = [stock_data[date]["volume"] for date in sorted_dates]
        ax2 = ax1.twinx()
        ax2.bar(sorted_dates, volumes, alpha=0.3, label="出来高", color='tab:orange')
        ax2.set_ylabel("出来高", fontproperties=jp_font)

    # カラーマップから異なる色を生成（青色を除外）
    colors = itertools.cycle(plt.cm.tab10.colors)
    # `tab:blue` を除外するために next(colors) を呼び出す
    next(colors)

    # 指標チャートのプロット
    for indicator_data in indicators_data:
        legend_name = indicator_data.get("legend_name", indicator_data["type"])
        indicator_values = indicator_data["values"]

        sorted_indicator_dates = sorted(indicator_values.keys())
        values = [indicator_values[date] for date in sorted_indicator_dates]

        # 指標チャートを異なる色でプロット
        ax1.plot(sorted_indicator_dates, values, label=legend_name, color=next(colors))

    # 凡例の作成
    legend = ax1.legend(loc="upper left", frameon=True, facecolor="#ADD8E6", prop=jp_font)

    # 凡例のテキストをデバッグ用に表示
    for text in legend.get_texts():
        print(text.get_text())

    # グリッドの設定
    ax1.grid(color="white")

    # レイアウトの調整
    fig.tight_layout(rect=[0, 0, 1, 1])

    # 画像をバッファに保存し、Base64エンコード
    buffer = BytesIO()
    plt.savefig(buffer, format="png")
    buffer.seek(0)
    plt.close()

    image_base64 = base64.b64encode(buffer.read()).decode("utf-8")

    return image_base64

def handle_generate_chart_request(request):
    stock_data = {}
    # リクエストから株価データを読み込み、辞書に格納
    for date_str, stock_data_pb in request.stock_data.items():
        date = datetime.strptime(date_str, "%Y-%m-%d")
        stock_data[date] = {
            "open": stock_data_pb.open,
            "close": stock_data_pb.close,
            "high": stock_data_pb.high,
            "low": stock_data_pb.low,
            "volume": stock_data_pb.volume,
        }

    indicators_data = []
    # リクエストから指標データを読み込み、リストに格納
    for indicator_pb in request.indicators:
        indicator_type = indicator_pb.type
        legend_name = f"{indicator_pb.legend_name}"
        indicator_values_pb = indicator_pb.values

        indicator_values = {}
        for date_str, value in indicator_values_pb.items():
            date = datetime.strptime(date_str, "%Y-%m-%d")
            indicator_values[date] = value

        indicators_data.append({
            "type": indicator_type,
            "legend_name": legend_name,
            "values": indicator_values,
        })

    return generate_chart(stock_data, indicators_data, request.include_volume)

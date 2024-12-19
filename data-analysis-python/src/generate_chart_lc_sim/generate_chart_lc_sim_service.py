# data-analysis-python\src\generate_chart_lc_sim\generate_chart_lc_sim_service.py
import matplotlib

matplotlib.use("Agg")  # 追加：GUIバックエンドを使用しない設定

import matplotlib.pyplot as plt
import base64
from io import BytesIO


def plot_stop_results(data, purchase_date, purchase_price, end_date, end_price):
    plt.figure(figsize=(14, 7))

    # 株価の推移をプロット
    plt.plot(data["Date"], data["Close"], label="Close Price")

    # 購入点をプロット
    plt.scatter(
        purchase_date, purchase_price, color="green", label="Purchase", zorder=5
    )
    plt.text(
        purchase_date,
        purchase_price,
        f" Purchase\n{purchase_price}",
        color="green",
        fontsize=12,
        ha="left",
    )

    # 売却点をプロット
    plt.scatter(end_date, end_price, color="red", label="Sell", zorder=5)
    plt.text(
        end_date, end_price, f" Sell\n{end_price}", color="red", fontsize=12, ha="left"
    )

    plt.xlabel("Date")
    plt.ylabel("Price")
    plt.title("Loss Cut Simulation Results")
    plt.legend()
    plt.xticks(rotation=45)
    plt.grid(True)
    plt.tight_layout()

    # プロットを画像として保存し、Base64エンコードする
    buffer = BytesIO()
    plt.savefig(buffer, format="png")
    buffer.seek(0)
    image_base64 = base64.b64encode(buffer.read()).decode("utf-8")
    buffer.close()

    # プロットをクローズしてリソースを解放
    plt.close()

    return image_base64

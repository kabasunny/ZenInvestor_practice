# generate_chart_service.py
import matplotlib.pyplot as plt
import base64
from io import BytesIO

def generate_chart(stock_data, indicators_data):
    # 株価データの日付をソート
    sorted_dates = sorted(stock_data.keys())
    
    # 株価データを準備
    close_prices = [stock_data[date]["close"] for date in sorted_dates]
    
    # プロットの作成
    plt.figure(figsize=(14, 7))
    plt.plot(sorted_dates, close_prices, label='Close Prices')
    
     # 各指標をプロット
    for indicator_data in indicators_data:
        indicator_type = indicator_data.indicator_type  # 属性アクセス
        indicator_values = indicator_data.values        # 属性アクセス

        plt.plot(indicator_values.keys(), indicator_values.values(), label=indicator_type)

    plt.xlabel('Date')
    plt.ylabel('Price')
    plt.title('Stock Prices and Indicators')
    plt.legend()
    
    # プロットを画像として保存
    buffer = BytesIO()
    plt.savefig(buffer, format='png')
    buffer.seek(0)
    
    # 画像データをBase64エンコード
    image_base64 = base64.b64encode(buffer.read()).decode('utf-8')
    
    return image_base64

def handle_generate_chart_request(request):
    stock_data = {
        date: {
            "open": values.open,
            "close": values.close,
            "high": values.high,
            "low": values.low,
            "volume": values.volume
        } for date, values in request.stock_data.items()
    }

    # indicators_data を request.indicators から直接作成
    indicators_data = request.indicators # 変更点

    return generate_chart(stock_data, indicators_data)

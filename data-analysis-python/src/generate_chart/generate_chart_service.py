import matplotlib.pyplot as plt

import base64
from io import BytesIO
from datetime import datetime

# GUIバックエンドを使用しないように設定 
import matplotlib
matplotlib.use('Agg')

def generate_chart(stock_data, indicators_data):
    # stock_dataのキーはdatetimeオブジェクトです
    sorted_dates = sorted(stock_data.keys())
    close_prices = [stock_data[date]["close"] for date in sorted_dates]

    plt.figure(figsize=(14, 7))
    plt.plot(sorted_dates, close_prices, label='終値')
    
    for indicator_data in indicators_data:
        indicator_type = indicator_data['type']
        indicator_values = indicator_data['values']

        # 日付をソートすることを確認
        sorted_indicator_dates = sorted(indicator_values.keys())
        values = [indicator_values[date] for date in sorted_indicator_dates]

        plt.plot(sorted_indicator_dates, values, label=indicator_type)

    plt.xlabel('日付')
    plt.ylabel('価格')
    plt.title('株価と指標')
    plt.legend()

    # プロットを画像として保存
    buffer = BytesIO()
    plt.savefig(buffer, format='png')
    buffer.seek(0)
    plt.close()  # 図を閉じる

    # 画像をBase64エンコード
    image_base64 = base64.b64encode(buffer.read()).decode('utf-8')

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
            "volume": stock_data_pb.volume
        }

    # リクエストから指標データを処理
    indicators_data = []
    for indicator_pb in request.indicators:
        indicator_type = indicator_pb.type
        indicator_values_pb = indicator_pb.values

        # 日付文字列をdatetimeオブジェクトに変換
        indicator_values = {}
        for date_str, value in indicator_values_pb.items():
            date = datetime.strptime(date_str, "%Y-%m-%d")
            indicator_values[date] = value

        indicators_data.append({
            'type': indicator_type,
            'values': indicator_values
        })

    return generate_chart(stock_data, indicators_data)

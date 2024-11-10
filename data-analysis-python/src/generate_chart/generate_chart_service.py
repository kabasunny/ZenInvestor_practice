# generate_chart_service.py
import matplotlib.pyplot as plt
import base64
from io import BytesIO
from datetime import datetime

def generate_chart(stock_data, indicators_data):
    # stock_data keys are datetime objects
    sorted_dates = sorted(stock_data.keys())
    close_prices = [stock_data[date]["close"] for date in sorted_dates]

    plt.figure(figsize=(14, 7))
    plt.plot(sorted_dates, close_prices, label='Close Prices')
    
    for indicator_data in indicators_data:
        indicator_type = indicator_data['type']
        indicator_values = indicator_data['values']

        # Ensure dates are sorted
        sorted_indicator_dates = sorted(indicator_values.keys())
        values = [indicator_values[date] for date in sorted_indicator_dates]

        plt.plot(sorted_indicator_dates, values, label=indicator_type)

    plt.xlabel('Date')
    plt.ylabel('Price')
    plt.title('Stock Prices and Indicators')
    plt.legend()

    # Save plot to image
    buffer = BytesIO()
    plt.savefig(buffer, format='png')
    buffer.seek(0)
    plt.close()  # Close the figure

    # Base64 encode the image
    image_base64 = base64.b64encode(buffer.read()).decode('utf-8')

    return image_base64

def handle_generate_chart_request(request):
    # Process stock_data from request
    stock_data = {}
    for date_str, stock_data_pb in request.stock_data.items():
        # Convert date string to datetime object
        date = datetime.strptime(date_str, "%Y-%m-%d")
        stock_data[date] = {
            "open": stock_data_pb.open,
            "close": stock_data_pb.close,
            "high": stock_data_pb.high,
            "low": stock_data_pb.low,
            "volume": stock_data_pb.volume
        }

    # Process indicators data from request
    indicators_data = []
    for indicator_pb in request.indicators:
        indicator_type = indicator_pb.type
        indicator_values_pb = indicator_pb.values

        # Convert date strings to datetime objects
        indicator_values = {}
        for date_str, value in indicator_values_pb.items():
            date = datetime.strptime(date_str, "%Y-%m-%d")
            indicator_values[date] = value

        indicators_data.append({
            'type': indicator_type,
            'values': indicator_values
        })

    return generate_chart(stock_data, indicators_data)
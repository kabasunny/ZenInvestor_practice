# data-analysis-python\src\get_stocks_datalist_with_dates\get_stocks_datalist_with_dates_service.py
import yfinance as yf
import time
from datetime import datetime, timedelta

def get_stocks_datalist_with_dates(symbols, start_date, end_date):
    stock_prices_list = []

    print(symbols, start_date, end_date)

    try:
        # シンボルに .T を追加の処理時間
        start_time_ticker = time.time()
        symbols_with_t = [symbol + ".T" for symbol in symbols]
        end_time_ticker = time.time()
        elapsed_time_ticker = end_time_ticker - start_time_ticker
        print(f"シンボルに .T を追加の処理時間: {elapsed_time_ticker:.2f}秒")

        # 終了日に1日加算
        end_date_dt = datetime.strptime(end_date, "%Y-%m-%d")
        end_date_plus_one = end_date_dt + timedelta(days=1)
        end_date_str = end_date_plus_one.strftime("%Y-%m-%d")

        # 複数のシンボルのデータを一度に取得の処理時間
        start_time_download = time.time()
        tickers = " ".join(symbols_with_t)
        data = yf.download(tickers, start=start_date, end=end_date_str, group_by='ticker')
        end_time_download = time.time()
        elapsed_time_download = end_time_download - start_time_download
        print(f"複数のシンボルのデータを一度に取得の処理時間: {elapsed_time_download:.2f}秒")

        # シンボルごとにデータを処理の処理時間
        start_time_processing = time.time()
        for symbol in symbols_with_t:
            if symbol in data:
                hist = data[symbol]
                for date, row in hist.iterrows():
                    stock_price = {
                        'symbol': symbol.replace(".T", ""),  # .T を取り除く
                        'date': date.strftime('%Y-%m-%d'),  # 日付フィールドを追加
                        'open': row['Open'],
                        'close': row['Close'],
                        'high': row['High'],
                        'low': row['Low'],
                        'volume': int(row['Volume']),  # intに変換
                        'turnover': row['Close'] * int(row['Volume'])  # 売買代金（終値 * 出来高）を追加
                    }
                    stock_prices_list.append(stock_price)
            else:
                print(f"Failed to get data for ticker: {symbol}")

        end_time_processing = time.time()
        elapsed_time_processing = end_time_processing - start_time_processing
        print(f"シンボルごとにデータを処理の処理時間: {elapsed_time_processing:.2f}秒")

        print(stock_prices_list)
        return stock_prices_list
    except Exception as e:
        print(f"Error fetching stock prices: {e}")
        return []

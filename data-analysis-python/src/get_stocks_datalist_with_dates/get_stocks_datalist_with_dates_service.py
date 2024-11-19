# data-analysis-python\src\get_stocks_datalist_with_dates\get_stocks_datalist_with_dates_service.py
import yfinance as yf

def get_stocks_datalist_with_dates(symbols, start_date, end_date):
    stock_prices_list = []
    try:
        for symbol in symbols:
            ticker = yf.Ticker(symbol)
            hist = ticker.history(start=start_date, end=end_date)
            if not hist.empty:
                for date, row in hist.iterrows():
                    stock_price = {
                        'symbol': symbol,
                        'date': date.strftime('%Y-%m-%d'),       # 日付フィールドを追加
                        'open': row['Open'],
                        'close': row['Close'],
                        'high': row['High'],
                        'low': row['Low'],
                        'volume': int(row['Volume']),            # intに変換
                        'turnover': row['Close'] * int(row['Volume'])  # 売買代金（終値 * 出来高）を追加
                    }
                    stock_prices_list.append(stock_price)
        return stock_prices_list
    except Exception as e:
        print(f"Error fetching stock prices: {e}")
        return []
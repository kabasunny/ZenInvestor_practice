# data-analysis-python\src\get_stocks_datalist_5d\get_stocks_datalist_5d_service.py
import yfinance as yf

def get_stocks_datalist_5d(symbols):
    stock_prices_list = []
    try:
        for symbol in symbols:
            ticker = yf.Ticker(symbol)
            hist = ticker.history(period="5d")
            if not hist.empty:
                for date, row in hist.iterrows():
                    stock_price = {
                        'symbol': symbol,
                        'date': date.strftime('%Y-%m-%d'),
                        'open': row['Open'],
                        'close': row['Close'],
                        'high': row['High'],
                        'low': row['Low'],
                        'volume': int(row['Volume']),              # ここでintに変換
                        'turnover': row['Close'] * int(row['Volume'])  # 売買代金（取引金額）を追加
                    }
                    stock_prices_list.append(stock_price)
        return stock_prices_list
    except Exception as e:
        print(f"Error fetching stock prices: {e}")
        return []

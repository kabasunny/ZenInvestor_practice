# data-analysis-python\src\get_stocks_datalist\get_stocks_datalist_service.py
import yfinance as yf

def get_stocks_datalist(symbols):
    stock_prices_list = []
    try:
        for symbol in symbols:
            ticker = yf.Ticker(symbol)
            hist = ticker.history(period="1d")
            if not hist.empty:
                date = hist.index[0].strftime('%Y-%m-%d')
                stock_price = {
                    'symbol': symbol,
                    'date': date,                             # 日付を追加
                    'open': hist['Open'].iloc[0],             # 始値
                    'close': hist['Close'].iloc[0],           # 終値
                    'high': hist['High'].iloc[0],             # 高値
                    'low': hist['Low'].iloc[0],               # 安値
                    'volume': int(hist['Volume'].iloc[0]),    # 出来高
                    'turnover': hist['Close'].iloc[0] * int(hist['Volume'].iloc[0])  # 売買代金（終値 * 出来高）
                }
                stock_prices_list.append(stock_price)
        return stock_prices_list
    except Exception as e:
        print(f"Error fetching stock prices: {e}")
        return []

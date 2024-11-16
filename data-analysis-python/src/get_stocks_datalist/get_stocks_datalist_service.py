# data-analysis-python\src\get_stocks_datalist\get_stocks_datalist_service.py
import yfinance as yf

def get_stocks_datalist(symbols):
    stock_prices_list = []
    try:
        for symbol in symbols:
            ticker = yf.Ticker(symbol)
            hist = ticker.history(period="1d")
            if not hist.empty:
                stock_price = {
                    'symbol': symbol,
                    'open': hist['Open'].iloc[0],    # 修正: ser[0] -> ser.iloc[0]
                    'close': hist['Close'].iloc[0],  # 修正: ser[0] -> ser.iloc[0]
                    'high': hist['High'].iloc[0],    # 修正: ser[0] -> ser.iloc[0]
                    'low': hist['Low'].iloc[0],      # 修正: ser[0] -> ser.iloc[0]
                    'volume': int(hist['Volume'].iloc[0])  # 修正: ser[0] -> ser.iloc[0]
                }
                stock_prices_list.append(stock_price)
        return stock_prices_list
    except Exception as e:
        print(f"Error fetching stock prices: {e}")
        return []

# data-analysis-python\src\get_jp_tickers\get_jp_tickers_service.py

import investpy

def get_jp_tickers():
    try:
        # 日本株のティッカーシンボルを取得
        stocks = investpy.stocks.get_stocks(country='japan')
        # 取得されるデータは、Investing.comが提供する最新の情報に基づく
        tickers = stocks['symbol'].tolist()
        return tickers
    except Exception as e:
        print(f"Error fetching tickers: {e}")
        return []

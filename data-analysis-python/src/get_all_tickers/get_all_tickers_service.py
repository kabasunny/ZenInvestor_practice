# data-analysis-python\src\get_all_tickers\get_all_tickers_service.py

import investpy

def get_all_tickers():
    try:
        # 全世界の株式ティッカーシンボルを取得
        stocks = investpy.stocks.get_stocks()
        # 取得されるデータは、Investing.comが提供する最新の情報に基づく
        tickers = stocks['symbol'].tolist()
        return tickers
    except Exception as e:
        print(f"Error fetching tickers: {e}")
        return []

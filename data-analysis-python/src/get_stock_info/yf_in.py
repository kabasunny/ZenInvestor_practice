import investpy
import yfinance as yf
from exchange_codes import exchange_codes  # exchange_codesのインポート

def fetch_stock_info(country):
    # 指定された国の株式情報を取得
    stocks = investpy.stocks.get_stocks(country=country)

    # 100番目から103番目の銘柄の情報を取得
    for index in range(99, 103):
        if index < len(stocks):
            stock_info = stocks.iloc[index]  # インデックスを指定して銘柄情報を取得
            print(stock_info)
            
            # 銘柄のシンボルを取得し、対応する取引所コードを追加
            exchange_code = exchange_codes.get(country.lower(), '')
            ticker = stock_info['symbol'] + exchange_code
            
            try:
                # yfinance を使ってティッカー情報を取得
                stock = yf.Ticker(ticker)
                info = stock.info

                # sector と industry を取得
                sector = info.get('sector', 'N/A')
                industry = info.get('industry', 'N/A')

                # 結果を表示
                print(f"Ticker: {ticker}")
                print(f"Sector: {sector}")
                print(f"Industry: {industry}")
            except Exception as e:
                print(f"Error fetching data for ticker {ticker}: {e}")
        else:
            print(f"インデックス {index} に対応するデータが存在しません。")

if __name__ == "__main__":
    # 各バリエーションの呼び出し
    countries = ["japan", "germany", "united kingdom", "france", "canada", "australia", "hong kong", "united states"]
    
    for country in countries:
        print(f"\nFetching stock info for {country.capitalize()}:")
        fetch_stock_info(country)

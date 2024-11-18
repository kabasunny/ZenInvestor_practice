import investpy
import yfinance as yf
import pandas as pd
import time

def get_stock_info_by_country(country):
    try:
        # 指定された国の株式リストを取得
        stocks = investpy.stocks.get_stocks(country=country)
        
        # 結果を格納するリスト
        stock_info_list = []
        
        # シンボルのマッピング辞書（必要に応じて拡張）
        symbol_mapping = {
            # 'investpy_symbol': 'yfinance_symbol',
            # 例:
            # 'BASFN': 'BAS',
        }
        
        # 国に対応する取引所サフィックス
        exchange_suffix = {
            'japan': '.T',
            'germany': '.DE',
            'united kingdom': '.L',
            'france': '.PA',
            'canada': '.TO',
            'australia': '.AX',
            'hong kong': '.HK',
            # 必要に応じて追加
        }
        
        suffix = exchange_suffix.get(country.lower(), '')
        
        for index, row in stocks.iterrows():
            # investpyからのシンボルを取得
            investpy_symbol = row['symbol']
            
            # マッピングが存在する場合は適用
            yf_symbol_base = symbol_mapping.get(investpy_symbol, investpy_symbol)
            
            # yfinance用のシンボルを作成
            yf_symbol = yf_symbol_base + suffix
            
            # yfinanceから情報を取得
            try:
                time.sleep(0.5)  # リクエスト間の遅延を挿入
                ticker = yf.Ticker(yf_symbol)
                info = ticker.info
                
                stock_info = {
                    'ticker': yf_symbol,  # yfinanceのシンボル
                    'name': info.get('longName') or info.get('shortName'),
                    'sector': info.get('sector'),
                    'industry': info.get('industry'),
                    'market_cap': info.get('marketCap'),
                    'listing_date': None
                }
                
                # 上場日の取得（可能な場合）
                if 'ipoDate' in info and info['ipoDate']:
                    stock_info['listing_date'] = info['ipoDate']
                elif 'firstTradeDateEpochUtc' in info and info['firstTradeDateEpochUtc']:
                    from datetime import datetime
                    timestamp = info['firstTradeDateEpochUtc']
                    stock_info['listing_date'] = datetime.fromtimestamp(timestamp).strftime('%Y-%m-%d')
                
                stock_info_list.append(stock_info)
            except Exception as e:
                print(f"Error fetching data for {yf_symbol}: {e}")
                continue  # エラーが発生した場合は次の銘柄に進む
        
        return stock_info_list
    
    except Exception as e:
        print(f"Error fetching stocks for {country}: {e}")
        return []
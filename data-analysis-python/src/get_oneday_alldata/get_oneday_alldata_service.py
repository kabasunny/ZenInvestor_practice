# get_oneday_alldata_service.py
import yfinance as yf
import pandas as pd
from datetime import datetime

def load_japanese_tickers():
    # デモ用に、日本のティッカーシンボルの小リストを用意
    tickers = ['6758.T', '7203.T', '9984.T', '9432.T']  # Sony, Toyota, SoftBank, NTT
    return tickers

def get_stock_data_for_date(ticker, date_str):
    # 日付文字列をdatetimeオブジェクトに変換
    date = datetime.strptime(date_str, '%Y-%m-%d')
    # データ取得の開始日と終了日を定義
    start_date = date.strftime('%Y-%m-%d')
    end_date = (date + pd.Timedelta(days=1)).strftime('%Y-%m-%d')

    try:
        # 指定した日付のデータを取得
        df = yf.download(ticker, start=start_date, end=end_date)
        if not df.empty:
            data = df.iloc[0]
            stock_data = {
                'open': data['Open'],
                'close': data['Close'],
                'high': data['High'],
                'low': data['Low'],
                'volume': data['Volume']
            }
            return stock_data
        else:
            # 指定した日付にデータがない場合
            return None
    except Exception as e:
        print(f"{date_str}の{ticker}のデータ取得エラー: {e}")
        return None

def get_oneday_alldata(date_str):
    tickers = load_japanese_tickers()
    stock_data = {}
    
    date = datetime.strptime(date_str, '%Y-%m-%d')
    start_date = date.strftime('%Y-%m-%d')
    end_date = (date + pd.Timedelta(days=1)).strftime('%Y-%m-%d')
    
    try:
        # すべてのティッカーのデータを一度にダウンロード
        df = yf.download(tickers, start=start_date, end=end_date, group_by='ticker')
        if df.empty:
            print(f"{date_str}にはどのティッカーにもデータがありません")
            return stock_data  # 空の辞書を返す
        
        # ティッカーが1つだけの場合（異なるDataFrame構造）
        if len(tickers) == 1:
            ticker = tickers[0]
            data = df.iloc[0]
            stock_data[ticker] = {
                'open': data['Open'],
                'close': data['Close'],
                'high': data['High'],
                'low': data['Low'],
                'volume': data['Volume']
            }
        else:
            # 複数のティッカー; 各ティッカーのデータを反復処理
            for ticker in tickers:
                if ticker in df.columns.levels[0]:
                    ticker_df = df[ticker]
                    if not ticker_df.empty:
                        data = ticker_df.iloc[0]
                        stock_data[ticker] = {
                            'open': data['Open'],
                            'close': data['Close'],
                            'high': data['High'],
                            'low': data['Low'],
                            'volume': data['Volume']
                        }
                    else:
                        print(f"{date_str}には{ticker}のデータがありません")
                else:
                    print(f"{date_str}には{ticker}のデータがありません")
    except Exception as e:
        print(f"データダウンロードエラー: {e}")
    
    return stock_data

# get_oneday_alldata_service.py
import yfinance as yf
import pandas as pd
from datetime import datetime

def load_japanese_tickers():
    # For demonstration purposes, a small list of Japanese ticker symbols
    tickers = ['6758.T', '7203.T', '9984.T', '9432.T']  # Sony, Toyota, SoftBank, NTT
    return tickers

def get_stock_data_for_date(ticker, date_str):
    # Convert the date string to a datetime object
    date = datetime.strptime(date_str, '%Y-%m-%d')
    # Define the start and end dates for the data retrieval
    start_date = date.strftime('%Y-%m-%d')
    end_date = (date + pd.Timedelta(days=1)).strftime('%Y-%m-%d')

    try:
        # Fetch data for the specified date
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
            # No data available for this date
            return None
    except Exception as e:
        print(f"Error fetching data for {ticker} on {date_str}: {e}")
        return None

def get_oneday_alldata(date_str):
    tickers = load_japanese_tickers()
    stock_data = {}
    
    date = datetime.strptime(date_str, '%Y-%m-%d')
    start_date = date.strftime('%Y-%m-%d')
    end_date = (date + pd.Timedelta(days=1)).strftime('%Y-%m-%d')
    
    try:
        # Download data for all tickers in a single request
        df = yf.download(tickers, start=start_date, end=end_date, group_by='ticker')
        if df.empty:
            print(f"No data available for any tickers on {date_str}")
            return stock_data  # Return an empty dictionary
        
        # Check if there's only one ticker (different DataFrame structure)
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
            # Multiple tickers; iterate over each ticker's data
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
                        print(f"No data for {ticker} on {date_str}")
                else:
                    print(f"No data for {ticker} on {date_str}")
    except Exception as e:
        print(f"Error downloading data: {e}")
    
    return stock_data
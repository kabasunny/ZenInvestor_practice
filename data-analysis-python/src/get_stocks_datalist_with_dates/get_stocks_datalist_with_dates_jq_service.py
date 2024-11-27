# data-analysis-python\src\get_stocks_datalist_with_dates\get_stocks_datalist_with_dates_jq_service.py
import os
import requests
import time
from dotenv import load_dotenv

# 環境変数を読み込み
load_dotenv()

def get_refresh_token():
    """リフレッシュトークンを取得する"""
    email = os.getenv("JQUANTS_EMAIL")
    password = os.getenv("JQUANTS_PASSWORD")
    url = "https://api.jquants.com/v1/token/auth_user"
    headers = {"Content-Type": "application/json"}
    data = {"mailaddress": email, "password": password}

    response = requests.post(url, headers=headers, json=data)
    response.raise_for_status()  # ステータスコードが200番台以外の場合に例外を投げる
    return response.json().get("refreshToken")

def get_id_token(refresh_token):
    """リフレッシュトークンを使ってIDトークンを取得する"""
    url = f"https://api.jquants.com/v1/token/auth_refresh?refreshtoken={refresh_token}"
    response = requests.post(url)
    response.raise_for_status()  # ステータスコードが200番台以外の場合に例外を投げる
    return response.json().get("idToken")

def get_daily_quotes(id_token, code=None, date=None, from_date=None, to_date=None):
    """日次株価データを取得する"""
    url = "https://api.jquants.com/v1/prices/daily_quotes"
    headers = {"Authorization": f"Bearer {id_token}"}
    params = {}
    
    if code:
        params['code'] = code
    if date:
        params['date'] = date
    if from_date and to_date:
        params['from'] = from_date
        params['to'] = to_date

    data = []
    while True:
        response = requests.get(url, headers=headers, params=params)
        try:
            response.raise_for_status()  # ステータスコードが200番台以外の場合に例外を投げる
            result = response.json()
            data.extend(result.get("daily_quotes", []))
            if "pagination_key" not in result:
                break
            params["pagination_key"] = result["pagination_key"]
        except requests.exceptions.HTTPError as e:
            print(f"Error fetching data for code {code}: {e}")
            break

        # レートリミット対応のための待機時間
        time.sleep(1)  # 1秒の待機

    return data

def get_stocks_datalist_with_dates_jq(codes=None, date=None, from_date=None, to_date=None):
    """複数の株価情報を取得して返す"""
    refresh_token = get_refresh_token()
    id_token = get_id_token(refresh_token)

    stock_prices_list = []

    for code in codes:
        quotes = get_daily_quotes(id_token, code=code, date=date, from_date=from_date, to_date=to_date)
        for quote in quotes:
            open_price = quote.get("Open", 0.0)
            close_price = quote.get("Close", 0.0)
            high_price = quote.get("High", 0.0)
            low_price = quote.get("Low", 0.0)
            volume = quote.get("Volume") if quote.get("Volume") is not None else 0
            turnover = close_price * volume if close_price is not None and volume is not None else 0.0

            stock_price = {
                "symbol": code,
                "date": quote.get("Date", ""),
                "open": open_price,
                "close": close_price,
                "high": high_price,
                "low": low_price,
                "volume": int(volume) if volume else 0,  # int() 関数を呼び出す前にチェック
                "turnover": turnover,
            }
            stock_prices_list.append(stock_price)

    return stock_prices_list

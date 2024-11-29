# data-analysis-python\src\get_trading_calendar_jq\get_trading_calendar_jq_service.py
import os
import requests
import pandas as pd
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

def get_trading_calendar(id_token, from_date, to_date):
    """取引カレンダーを取得する"""
    url = "https://api.jquants.com/v1/markets/trading_calendar"
    headers = {"Authorization": f"Bearer {id_token}"}
    params = {
        "from": from_date,
        "to": to_date
    }

    response = requests.get(url, headers=headers, params=params)
    response.raise_for_status()  # ステータスコードが200番台以外の場合に例外を投げる
    data = response.json()["trading_calendar"]

    # データフレームに変換
    trading_calendar_df = pd.DataFrame(data)

    return trading_calendar_df

def fetch_trading_calendar(from_date, to_date):
    """取引カレンダーを取得して返す"""
    refresh_token = get_refresh_token() # リフレッシュトークンの有効期間は1週間
    id_token = get_id_token(refresh_token) # IDトークンの有効期間は24時間
    return get_trading_calendar(id_token, from_date, to_date)
    # 非営業日:0 営業日:1 東証半日立会日:2 非営業日(祝日取引あり):3

# 使用例
if __name__ == "__main__":
    from_date = "2023-01-01"
    to_date = "2023-12-31"

    trading_calendar = fetch_trading_calendar(from_date, to_date)
    print(trading_calendar)

# data-analysis-python\src\get_stocks_datalist_with_dates_jq\get_stock_info_only_tickers_jq_service.py

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

def get_listed_companies(id_token):
    """上場銘柄一覧を取得し、指定形式でデータフレームを作成する"""
    url = "https://api.jquants.com/v1/listed/info"
    headers = {"Authorization": f"Bearer {id_token}"}

    response = requests.get(url, headers=headers)
    response.raise_for_status()  # ステータスコードが200番台以外の場合に例外を投げる
    data = response.json()["info"]

    # データフレームに変換し、必要なカラムを選択
    listed_companies = pd.DataFrame(data).rename(columns={
        "Code": "ticker",
        "CompanyName": "name",
        "Sector17CodeName": "sector",
        "Sector33CodeName": "industry"
    })[["ticker", "name", "sector", "industry"]]

    # ティッカーの末尾の0を削除（5桁以上かつ末尾が0の場合）
    def trim_ticker(ticker):
        if len(ticker) > 4 and ticker.endswith('0'):
            return ticker[:-1]
        return ticker

    listed_companies['ticker'] = listed_companies['ticker'].apply(trim_ticker)

    return listed_companies

def fetch_stock_info():
    """株式情報を取得して返す"""
    refresh_token = get_refresh_token()
    id_token = get_id_token(refresh_token)
    return get_listed_companies(id_token)

def save_tickers_to_csv():
    """銘柄コードのみを抽出してCSVに保存する"""
    listed_companies = fetch_stock_info()
    tickers = listed_companies['ticker'].unique()
    tickers_df = pd.DataFrame(tickers, columns=['Ticker'])
    output_csv_path = "./src/get_stocks_datalist_with_dates_jq/listed_tickers.csv"
    tickers_df.to_csv(output_csv_path, index=False)
    print(f"銘柄コード一覧が保存されました: {output_csv_path}")

# 使用例
save_tickers_to_csv()


# 実行コマンド
# python src/get_stocks_datalist_with_dates_jq/get_stock_info_only_tickers_jq_service.py

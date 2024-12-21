# data-analysis-python\src\get_stock_data_with_dates\get_stock_data_with_dates_service.py
import yfinance as yf
import pandas as pd


def get_stock_data_with_dates(symbol, start_date, end_date):
    print(symbol, start_date, end_date)
    # 日足データの取得
    daily_data = yf.download(symbol, start=start_date, end=end_date, interval="1d")

    stock = yf.Ticker(symbol)
    stock_info = stock.info  # 銘柄情報を取得
    stock_name = stock_info.get(
        "longName", symbol
    )  # 銘柄名を取得、存在しない場合はシンボルを使用

    # 日付部分のみを取得
    daily_data.index = daily_data.index.strftime(
        "%Y-%m-%d"
    )  # これがないせいで、Go側で日付キーでデータ取得に泣かされた　こんな感じ　2021-12-31 00:00:00

    # DataFrameを辞書形式に変換
    stock_data_dict = daily_data.to_dict("index")

    return stock_name, stock_data_dict  # 銘柄名とすべての日付のデータを返す


# API（今回の例では yfinance）が、デフォルトでデータを pandas.DataFrame 形式で返す
# gRPC は、通信プロトコルとしてプロトコルバッファ（Protocol Buffers）を使用し
# プロトコルバッファは、データを効率的にシリアライズ（直列化）し、ネットワーク越しにデータを送受信するための形式ですが、基本的なデータ型にしか対応していない
# 具体的には、文字列、整数、浮動小数点数、ブール値、およびこれらのデータ型のリストやマップなどが使用できる
# 問題点: pandas.DataFrame は、gRPC のメッセージフィールドとしては使用できない。基本的なデータ型に変換する必要がある
# 解決方法: pandas.DataFrame を辞書型に変換し、その辞書型データをプロトコルバッファのメッセージフィールドにマッピングする

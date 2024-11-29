# StartForRankingDataMicroservices.ps1

# スクリプトの開始ディレクトリを設定 (相対パス)
$BasePath = "."
Set-Location $BasePath

# 各サービスの相対パスを設定
# 日本株:銘柄データ一括取得(J-Quantsの制約あり)
$GetStockInfoJqServicePath = "./data-analysis-python/src/get_stock_info_jq/get_stock_info_jq_grpc.py"
# 日本株:株価データ一括取得(J-Quantsの制約あり)
$GetStocksDatalistWithDatesServicePath = "./data-analysis-python/src/get_stocks_datalist_with_dates/get_stocks_datalist_with_dates_grpc.py"
# JPX営業日データ(J-Quantsの制約あり)
$GetTradingCalendarJqServicePath = "./data-analysis-python/src/get_trading_calendar_jq/get_trading_calendar_jq_grpc.py"

# 各サービスを新しいターミナルウィンドウで起動

Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockInfoJqServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStocksDatalistWithDatesServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetTradingCalendarJqServicePath

# サービスの起動メッセージを表示
Write-Output "Ranking services have been started in new terminal windows."

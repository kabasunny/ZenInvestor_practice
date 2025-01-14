# StartMicroservices.ps1

# スクリプトの開始ディレクトリを設定 (相対パス)
$BasePath = "."
Set-Location $BasePath

# 各サービスの相対パスを設定
# 指標SAM
$SimpleMovingAverageServicePath = "./data-analysis-python/src/calculate_indicator/simple_moving_average/simple_moving_average_grpc.py"
# チャート生成
$GenerateChartServicePath = "./data-analysis-python/src/generate_chart/generate_chart_grpc.py"
# 株価データ単一銘柄
$GetStockDataServicePath = "./data-analysis-python/src/get_stock_data/get_stock_data_grpc.py"

# # 日本株:銘柄全データ一括取得(J-Quantsの制約あり)
# $GetStockInfoJqServicePath = "./data-analysis-python/src/get_stock_info_jq/get_stock_info_jq_grpc.py"
# # 日本株:株価全データ一括取得(J-Quantsの制約あり)
# $GetStocksDatalistWithDatesServicePath = "./data-analysis-python/src/get_stocks_datalist_with_dates/get_stocks_datalist_with_dates_grpc.py"
# # JPX営業日データ(J-Quantsの制約あり)
# $GetTradingCalendarJqServicePath = "./data-analysis-python/src/get_trading_calendar_jq/get_trading_calendar_jq_grpc.py"

# 株価データ単一銘柄を取得する期間を日付で指定して取得するサービス
$GetStockDataWithDatesServicePath = "./data-analysis-python/src/get_stock_data_with_dates/get_stock_data_with_dates_grpc.py"
# ロスカットシミュレーションのチャート生成
$GenerateLC_SIM_ChartServicePath = "./data-analysis-python/src/generate_chart_lc_sim/generate_chart_lc_sim_grpc.py"

# バッチ処理は無くても良いな
# 各サービスを新しいターミナルウィンドウで起動
Start-Process "powershell" -ArgumentList "-NoExit", "python", $SimpleMovingAverageServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GenerateChartServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockDataServicePath
# Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockInfoJqServicePath
# Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStocksDatalistWithDatesServicePath
# Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetTradingCalendarJqServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockDataWithDatesServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GenerateLC_SIM_ChartServicePath

# 3秒待機
Start-Sleep -Seconds 3

# バックエンドAPIサーバーをターミナルウィンドウで起動
Start-Process powershell -ArgumentList "Start-Process powershell -ArgumentList 'cd ./api-go; air; Pause'"

# さらに3秒待機
Start-Sleep -Seconds 3

# フロントエンドサーバーをターミナルウィンドウで起動
Start-Process powershell -ArgumentList "Start-Process powershell -ArgumentList 'cd ./frontend-react; npm start; Pause'"

# サービスの起動メッセージを表示
Write-Output "All services have been started in new terminal windows."

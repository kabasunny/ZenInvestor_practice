# StartMicroservices.ps1

# スクリプトの開始ディレクトリを設定 (相対パス)
$BasePath = "."
Set-Location $BasePath

# 各サービスの相対パスを設定
$SimpleMovingAverageServicePath = "./data-analysis-python/src/calculate_indicator/simple_moving_average/simple_moving_average_grpc.py"
$GenerateChartServicePath = "./data-analysis-python/src/generate_chart/generate_chart_grpc.py"
$GetStockDataServicePath = "./data-analysis-python/src/get_stock_data/get_stock_data_grpc.py"
$GetStockInfoJqServicePath = "./data-analysis-python/src/get_stock_info_jq/get_stock_info_jq_grpc.py"
$GetStocksDatalistWithDatesServicePath = "./data-analysis-python/src/get_stocks_datalist_with_dates/get_stocks_datalist_with_dates_grpc.py"

# 各サービスを新しいターミナルウィンドウで起動
Start-Process "powershell" -ArgumentList "-NoExit", "python", $SimpleMovingAverageServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GenerateChartServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockDataServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockInfoJqServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStocksDatalistWithDatesServicePath

# サービスの起動メッセージを表示
Write-Output "All services have been started in new terminal windows."

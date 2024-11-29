# StartForStockAndChartMicroservices.ps1

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

# 各サービスを新しいターミナルウィンドウで起動
Start-Process "powershell" -ArgumentList "-NoExit", "python", $SimpleMovingAverageServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GenerateChartServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockDataServicePath

# サービスの起動メッセージを表示
Write-Output "Stock and Chart services have been started in new terminal windows."

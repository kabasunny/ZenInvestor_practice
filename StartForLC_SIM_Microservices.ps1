# StartForLC_SIM_Microservices.ps1.ps1

# スクリプトの開始ディレクトリを設定 (相対パス)
$BasePath = "."
Set-Location $BasePath

# 各サービスの相対パスを設定
# 株価データ単一銘柄を取得する期間を日付で指定して取得するサービス
$GetStockDataWithDatesServicePath = "./data-analysis-python/src/get_stock_data_with_dates/get_stock_data_with_dates_grpc.py"
# ロスカットシミュレーションのチャート生成
$GenerateLC_SIM_ChartServicePath = "./data-analysis-python/src/generate_chart_lc_sim/generate_chart_lc_sim_grpc.py"

# 各サービスを新しいターミナルウィンドウで起動
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GetStockDataWithDatesServicePath
Start-Process "powershell" -ArgumentList "-NoExit", "python", $GenerateLC_SIM_ChartServicePath

# サービスの起動メッセージを表示
Write-Output "Stock and Chart services have been started in new terminal windows."

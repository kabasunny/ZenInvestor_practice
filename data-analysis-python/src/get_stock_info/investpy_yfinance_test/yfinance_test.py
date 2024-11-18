import yfinance as yf

# トヨタ自動車のティッカーシンボル
ticker = "6502.T"

# yfinance を使ってティッカー情報を取得
stock = yf.Ticker(ticker)
info = stock.info

# すべてのフィールドを表示
for key, value in info.items():
    print(f"{key}: {value}")

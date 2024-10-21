from flask import Flask, send_file
import yfinance as yf
import matplotlib.pyplot as plt
import io

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "Hello, World!"

@app.route("/plot")
def plot():
    stock = yf.Ticker("AAPL")
    data = stock.history(period="1mo")
    plt.figure(figsize=(10, 5))
    plt.plot(data.index, data["Close"], label="Close Price", color="blue")
    plt.title("AAPL Stock Price")
    plt.xlabel("Date")
    plt.ylabel("Price (USD)")
    plt.legend()
    img = io.BytesIO()
    plt.savefig(img, format="png")
    img.seek(0)
    plt.close()
    return send_file(img, mimetype="image/png")

if __name__ == "__main__": 
    app.run(host="0.0.0.0", port=5006, debug=True)
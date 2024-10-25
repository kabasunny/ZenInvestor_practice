# routes/get_stock_data.py
from flask import Blueprint, jsonify, request
from src.services.get_stock_service import get_stock_data

bp = Blueprint("get_stock_data", __name__, url_prefix="/get-stock-data")


@bp.route("/", methods=["POST"])
def get_data():
    data = request.json
    ticker = data.get("ticker")

    if not ticker:
        return jsonify({"error": "Ticker is required"}), 400

    stock_data = get_stock_data(ticker)
    return jsonify(stock_data)

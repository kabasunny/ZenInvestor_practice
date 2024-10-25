# routes/generate_chart.py
from flask import Blueprint, jsonify, request
from src.services.generate_chart_service import generate_chart

bp = Blueprint("generate_chart", __name__, url_prefix="/generate-chart")


@bp.route("/", methods=["POST"])
def generate():
    data = request.json
    ticker = data.get("ticker")

    if not ticker:
        return jsonify({"error": "Ticker is required"}), 400

    chart_url = generate_chart(ticker)
    return jsonify({"chart_url": chart_url})

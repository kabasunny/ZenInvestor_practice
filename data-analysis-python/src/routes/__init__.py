# routes/__init__.py
from flask import Blueprint

bp = Blueprint("routes", __name__)

from . import get_stock_data, generate_chart

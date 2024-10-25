# tests/test_get_stock_data.py
import unittest
from flask import Flask
from src.routes.get_stock_data import bp

app = Flask(__name__)
app.register_blueprint(bp)


class TestGetStockData(unittest.TestCase):
    def test_get_stock_data(self):
        with app.test_client() as client:
            response = client.post("/get-stock-data/", json={"ticker": "^GSPC"})
            self.assertEqual(response.status_code, 200)
            self.assertIn("Close", response.get_json())


if __name__ == "__main__":
    unittest.main()

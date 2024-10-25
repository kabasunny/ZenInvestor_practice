# tests/test_generate_chart.py
import unittest
from flask import Flask
from src.routes.generate_chart import bp

app = Flask(__name__)
app.register_blueprint(bp)


class TestGenerateChart(unittest.TestCase):
    def test_generate_chart(self):
        with app.test_client() as client:
            response = client.post("/generate-chart/", json={"ticker": "^GSPC"})
            self.assertEqual(response.status_code, 200)
            self.assertIn("chart_url", response.get_json())


if __name__ == "__main__":
    unittest.main()

# data-analysis-python\src\get_stocks_datalist_with_dates\test_get_stocks_datalist_with_dates_grpc.py
import unittest
import grpc
import time
import pandas as pd
import os
import get_stocks_datalist_with_dates_pb2
import get_stocks_datalist_with_dates_pb2_grpc
from concurrent import futures
from get_stocks_datalist_with_dates_grpc import GetStocksDatalistWithDatesService

class TestGetStocksDatalistWithDatesGrpc(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # サーバーを別スレッドで起動
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_stocks_datalist_with_dates_pb2_grpc.add_GetStocksDatalistWithDatesServiceServicer_to_server(
            GetStocksDatalistWithDatesService(), cls.server)
        cls.port = '51066'  # テスト時にはポートを50066から51066に変更
        cls.server.add_insecure_port(f'[::]:{cls.port}')
        cls.server.start()
        print(f'Server started on port {cls.port}')
        time.sleep(1)  # サーバーが起動するまで待機

    @classmethod
    def tearDownClass(cls):
        cls.server.stop(0)
        print('Server stopped')

    def setUp(self):
        # gRPCクライアントを作成
        self.channel = grpc.insecure_channel(f'localhost:{self.port}')
        self.stub = get_stocks_datalist_with_dates_pb2_grpc.GetStocksDatalistWithDatesServiceStub(self.channel)

    def test_get_stocks_datalist_with_dates_us(self):
        symbols = ["AAPL", "MSFT", "GOOGL"]  # Apple, Microsoft, Google
        start_date = "2023-11-9"
        end_date = "2023-11-15"
        self._test_get_stocks_datalist_with_dates(symbols, start_date, end_date, 'us_stock_prices_with_dates_grpc.csv')

    def _test_get_stocks_datalist_with_dates(self, symbols, start_date, end_date, output_file_name):
        # リクエストを作成
        request = get_stocks_datalist_with_dates_pb2.GetStocksDatalistWithDatesRequest(
            symbols=symbols, start_date=start_date, end_date=end_date)
        # サービスを呼び出し
        response = self.stub.GetStocksDatalist(request)

        # レスポンスを確認
        self.assertTrue(len(response.stock_prices) > 0, "株価情報が取得できませんでした")
        print(f"取得した株価情報の件数: {len(response.stock_prices)}")

        # データフレームとして保存
        stock_prices_list = []
        for stock_price in response.stock_prices:
            stock_prices_list.append({
                'symbol': stock_price.symbol,
                'date': stock_price.date,
                'open': stock_price.open,
                'close': stock_price.close,
                'high': stock_price.high,
                'low': stock_price.low,
                'volume': stock_price.volume,
                'turnover': stock_price.turnover      # 売買代金（取引金額）を追加
            })
        
        df = pd.DataFrame(stock_prices_list)
        output_dir = "src/get_stocks_datalist_with_dates/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, output_file_name)
        df.to_csv(output_file, index=False)
        
        print(f"株価情報がCSVファイルとして保存されました: {output_file}")

if __name__ == '__main__':
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates -p 'test_get_stocks_datalist_with_dates_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_stocks_datalist_with_dates -p 'test*.py'

# data-analysis-python\src\get_stocks_datalist\test_get_stocks_datalist_grpc.py

import unittest
import grpc
import time
import pandas as pd
import os
import get_stocks_datalist_pb2
import get_stocks_datalist_pb2_grpc
from concurrent import futures
from get_stocks_datalist_grpc import GetStocksDatalistService

class TestGetStocksDatalistGrpc(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # サーバーを別スレッドで起動
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_stocks_datalist_pb2_grpc.add_GetStocksDatalistServiceServicer_to_server(
            GetStocksDatalistService(), cls.server)
        cls.port = '51065' # テスト時にはポートを50065から51065に変更
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
        self.stub = get_stocks_datalist_pb2_grpc.GetStocksDatalistServiceStub(self.channel)

    def test_get_stocks_datalist_germany(self):
        symbols = ["DAI.DE", "SAP.DE", "BAYN.DE"]  # Daimler, SAP, Bayer
        self._test_get_stocks_datalist(symbols, 'de_stock_prices_grpc.csv')

    def test_get_stocks_datalist_india(self):
        symbols = ["TCS.NS", "RELIANCE.NS", "INFY.NS"]  # TCS, Reliance, Infosys
        self._test_get_stocks_datalist(symbols, 'in_stock_prices_grpc.csv')

    def _test_get_stocks_datalist(self, symbols, output_file_name):
        # リクエストを作成
        request = get_stocks_datalist_pb2.GetStocksDatalistRequest(symbols=symbols)
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
                'open': stock_price.open,
                'close': stock_price.close,
                'high': stock_price.high,
                'low': stock_price.low,
                'volume': stock_price.volume
            })
        
        df = pd.DataFrame(stock_prices_list)
        output_dir = "src/get_stocks_datalist/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, output_file_name)
        df.to_csv(output_file, index=False)
        
        print(f"株価情報がCSVファイルとして保存されました: {output_file}")

if __name__ == '__main__':
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stocks_datalist -p 'test_get_stocks_datalist_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_stocks_datalist -p 'test*.py'

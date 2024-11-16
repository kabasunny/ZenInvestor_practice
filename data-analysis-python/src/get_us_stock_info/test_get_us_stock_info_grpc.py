# data-analysis-python\src\get_us_stock_info\test_get_us_stock_info_grpc.py

import unittest
import grpc
import time
import pandas as pd
import os
import get_us_stock_info_pb2
import get_us_stock_info_pb2_grpc
from concurrent import futures
from get_us_stock_info_grpc import GetUsStockInfoService

class TestGetUsStockInfoGrpc(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # サーバーを別スレッドで起動
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_us_stock_info_pb2_grpc.add_GetUsStockInfoServiceServicer_to_server(
            GetUsStockInfoService(), cls.server)
        cls.port = '51063' # テスト時にはポートを50063から51063に変更
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
        self.stub = get_us_stock_info_pb2_grpc.GetUsStockInfoServiceStub(self.channel)

    def test_get_us_stock_info(self):
        # リクエストを作成
        request = get_us_stock_info_pb2.GetUsStockInfoRequest()
        # サービスを呼び出し
        response = self.stub.GetUsStockInfo(request)

        # レスポンスを確認
        self.assertTrue(len(response.stocks) > 0, "株式情報が取得できませんでした")
        print(f"取得した株式情報の件数: {len(response.stocks)}")  # 取得件数を表示

        # データフレームとして保存
        stock_info_list = []
        for stock_info in response.stocks:
            stock_info_list.append({
                'country': stock_info.country,
                'symbol': stock_info.symbol,
                'name': stock_info.name,
                'full_name': stock_info.full_name,
                'isin': stock_info.isin,
                'currency': stock_info.currency,
                'stock_exchange': stock_info.stock_exchange,
                'sector': stock_info.sector,
                'industry': stock_info.industry
            })
        
        df = pd.DataFrame(stock_info_list)
        output_dir = "src/get_us_stock_info/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, "us_stock_info_grpc.csv")
        df.to_csv(output_file, index=False)
        
        print(f"株式情報がCSVファイルとして保存されました: {output_file}")

if __name__ == '__main__':
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_us_stock_info -p 'test_get_us_stock_info_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_us_stock_info -p 'test*.py'

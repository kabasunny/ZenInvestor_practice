import unittest
import grpc
import time
import pandas as pd
import os
import get_all_tickers_pb2
import get_all_tickers_pb2_grpc
from concurrent import futures
from get_all_tickers_grpc import GetAllTickersService

class TestGetAllTickersGrpc(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # サーバーを別スレッドで起動
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_all_tickers_pb2_grpc.add_GetAllTickersServiceServicer_to_server(
            GetAllTickersService(), cls.server)
        cls.port = '51054' # テスト時にはポートを50054から51054に変更
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
        self.stub = get_all_tickers_pb2_grpc.GetAllTickersServiceStub(self.channel)

    def test_get_all_tickers(self):
        # リクエストを作成
        request = get_all_tickers_pb2.GetAllTickersRequest()
        # サービスを呼び出し
        response = self.stub.GetAllTickers(request)

        # レスポンスを確認
        self.assertTrue(len(response.tickers) > 0, "ティッカーシンボルが取得できませんでした")
        print(f"取得したティッカーシンボル: {response.tickers[:5]}...")  # 最初の5件を表示

        # データフレームとして保存
        tickers = list(response.tickers)
        df = pd.DataFrame(tickers, columns=['Ticker'])
        output_dir = "src/get_all_tickers/test_output"  # 出力ディレクトリを指定
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        
        # CSVファイルとして保存
        output_file = os.path.join(output_dir, "all_tickers_grpc.csv")
        df.to_csv(output_file, index=False)
        
        print(f"ティッカーシンボルがCSVファイルとして保存されました: {output_file}")

if __name__ == '__main__':
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_all_tickers -p 'test_get_all_tickers_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_all_tickers -p 'test*.py'

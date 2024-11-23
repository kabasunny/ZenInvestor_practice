# data-analysis-python\src\get_stock_info_jq\test_get_stock_info_jq_grpc.py
import unittest
import grpc
from concurrent import futures
import get_stock_info_jq_pb2
import get_stock_info_jq_pb2_grpc
from get_stock_info_jq_grpc import GetStockInfoJqService, serve
from get_stock_info_jq_service import fetch_stock_info
import pandas as pd
import os
import threading

class TestGetStockInfoJqGrpcService(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        """gRPCサーバーをセットアップ"""
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_stock_info_jq_pb2_grpc.add_GetStockInfoJqServiceServicer_to_server(GetStockInfoJqService(), cls.server)
        cls.port = 50404
        cls.server.add_insecure_port(f'[::]:{cls.port}')
        cls.server.start()
        print(f"Test gRPC Server started, listening on port {cls.port}")

    @classmethod
    def tearDownClass(cls):
        """gRPCサーバーを停止"""
        cls.server.stop(None)
        print("Test gRPC Server stopped")

    def test_get_stock_info_jq(self):
        """gRPCサービスのGetStockInfoJqメソッドをテスト"""
        # gRPCクライアントの設定
        with grpc.insecure_channel(f'localhost:{self.port}') as channel:
            stub = get_stock_info_jq_pb2_grpc.GetStockInfoJqServiceStub(channel)
            request = get_stock_info_jq_pb2.GetStockInfoJqRequest() #(country="Japan") J-QUANTSは日本株なので無し11/23
            response = stub.GetStockInfoJq(request)

        # レスポンスの検証
        self.assertGreater(len(response.stocks), 0)
        
        # CSV出力
        output_dir = "src/get_stock_info_jq/test_output"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        output_file = os.path.join(output_dir, "grpc_listed_companies.csv")
        df = pd.DataFrame([{
            "ticker": stock.ticker,
            "name": stock.name,
            "sector": stock.sector,
            "industry": stock.industry
        } for stock in response.stocks])
        df.to_csv(output_file, index=False)
        self.assertTrue(os.path.exists(output_file))
        print(f"CSVファイルが保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_stock_info_jq -p 'test_get_stock_info_jq_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_stock_info_jq -p 'test*.py'
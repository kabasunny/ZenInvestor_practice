# data-analysis-python\src\get_trading_calendar_jq\test_get_trading_calendar_jq_grpc.py

import unittest
import grpc
from concurrent import futures
import get_trading_calendar_jq_pb2
import get_trading_calendar_jq_pb2_grpc
from get_trading_calendar_jq_grpc import GetTradingCalendarJqService, serve
from get_trading_calendar_jq_service import fetch_trading_calendar
import pandas as pd
import os
import threading

class TestGetTradingCalendarJqGrpcService(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        """gRPCサーバーをセットアップ"""
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        get_trading_calendar_jq_pb2_grpc.add_GetTradingCalendarJqServiceServicer_to_server(GetTradingCalendarJqService(), cls.server)
        cls.port = 50003
        cls.server.add_insecure_port(f'[::]:{cls.port}')
        cls.server.start()
        print(f"Test gRPC Server started, listening on port {cls.port}")

    @classmethod
    def tearDownClass(cls):
        """gRPCサーバーを停止"""
        cls.server.stop(None)
        print("Test gRPC Server stopped")

    def test_get_trading_calendar_jq(self):
        """gRPCサービスのGetTradingCalendarJqメソッドをテスト"""
        # gRPCクライアントの設定
        with grpc.insecure_channel(f'localhost:{self.port}') as channel:
            stub = get_trading_calendar_jq_pb2_grpc.GetTradingCalendarJqServiceStub(channel)
            request = get_trading_calendar_jq_pb2.GetTradingCalendarJqRequest(
                from_date="2023-01-01",
                to_date="2023-12-31"
            )
            response = stub.GetTradingCalendarJq(request)

        # レスポンスの検証
        self.assertGreater(len(response.trading_calendar), 0)
        
        # CSV出力
        output_dir = "src/get_trading_calendar_jq/test_output"
        if not os.path.exists(output_dir):
            os.makedirs(output_dir)
        output_file = os.path.join(output_dir, "grpc_trading_calendar.csv")
        df = pd.DataFrame([{
            "Date": trading_calendar.Date,
            "HolidayDivision": trading_calendar.HolidayDivision
        } for trading_calendar in response.trading_calendar])
        df.to_csv(output_file, index=False)
        self.assertTrue(os.path.exists(output_file))
        print(f"CSVファイルが保存されました: {output_file}")

if __name__ == "__main__":
    unittest.main()

# 本ファイル単体テスト
# python -m unittest discover -s src/get_trading_calendar_jq -p 'test_get_trading_calendar_jq_grpc.py'

# 一括テスト
# python -m unittest discover -s src/get_trading_calendar_jq -p 'test*.py'

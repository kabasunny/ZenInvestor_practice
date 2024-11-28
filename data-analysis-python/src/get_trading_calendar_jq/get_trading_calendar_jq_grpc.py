# data-analysis-python\src\get_trading_calendar_jq\get_trading_calendar_jq_grpc.py

from concurrent import futures
import grpc
import get_trading_calendar_jq_pb2
import get_trading_calendar_jq_pb2_grpc
from get_trading_calendar_jq_service import fetch_trading_calendar
import time

class GetTradingCalendarJqService(get_trading_calendar_jq_pb2_grpc.GetTradingCalendarJqServiceServicer):
    def GetTradingCalendarJq(self, request, context):
        
        print("gRPCサーバー : get_trading_calendar_jq_service リクエスト")

        # 処理開始時刻の記録
        start_time = time.time()

        # サービス呼び出し
        trading_calendar_df = fetch_trading_calendar(request.from_date, request.to_date)

        # レスポンス作成
        response = get_trading_calendar_jq_pb2.GetTradingCalendarJqResponse()
        for index, row in trading_calendar_df.iterrows():
            trading_calendar_pb = get_trading_calendar_jq_pb2.TradingCalendar(
                Date=row['Date'],
                HolidayDivision=row['HolidayDivision']
            )
            response.trading_calendar.append(trading_calendar_pb)
        
        # 処理終了時刻の記録
        end_time = time.time()
        elapsed_time = end_time - start_time
        
        print(f"gRPCサーバー : get_trading_calendar_jq_service レスポンス - 処理時間: {elapsed_time:.2f}秒")
            
        return response

def serve():
    port = '50003'  # ポート番号を指定
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_trading_calendar_jq_pb2_grpc.add_GetTradingCalendarJqServiceServicer_to_server(GetTradingCalendarJqService(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    print(f"GetTradingCalendarJq gRPC Server started, listening on port {port}")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

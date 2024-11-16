# data-analysis-python\src\get_jp_tickers\get_jp_tickers_grpc.py

from concurrent import futures
import grpc
import get_jp_tickers_pb2
import get_jp_tickers_pb2_grpc
from get_jp_tickers_service import get_jp_tickers

class GetJpTickersService(get_jp_tickers_pb2_grpc.GetJpTickersServiceServicer):
    def GetJpTickers(self, request, context):
        tickers = get_jp_tickers()
        response = get_jp_tickers_pb2.GetJpTickersResponse()
        response.tickers.extend(tickers)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_jp_tickers_pb2_grpc.add_GetJpTickersServiceServicer_to_server(GetJpTickersService(), server)
    server.add_insecure_port('[::]:50055')
    server.start()
    print("GetJpTickers gRPCServer started, listening on port 50055")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

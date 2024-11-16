# data-analysis-python\src\get_all_tickers\get_all_tickers_grpc.py

from concurrent import futures
import grpc
import get_all_tickers_pb2
import get_all_tickers_pb2_grpc
from get_all_tickers_service import get_all_tickers

class GetAllTickersService(get_all_tickers_pb2_grpc.GetAllTickersServiceServicer):
    def GetAllTickers(self, request, context):
        tickers = get_all_tickers()
        response = get_all_tickers_pb2.GetAllTickersResponse()
        response.tickers.extend(tickers)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_all_tickers_pb2_grpc.add_GetAllTickersServiceServicer_to_server(GetAllTickersService(), server)
    server.add_insecure_port('[::]:50054')
    server.start()
    print("GetAllTickers gRPCServer started, listening on port 50054")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

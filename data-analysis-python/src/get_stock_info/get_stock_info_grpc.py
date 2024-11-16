# data-analysis-python\src\get_stock_info\get_stock_info_grpc.py

from concurrent import futures
import grpc
import get_stock_info_pb2
import get_stock_info_pb2_grpc
from get_stock_info_service import get_stock_info

class GetStockInfoService(get_stock_info_pb2_grpc.GetStockInfoServiceServicer):
    def GetStockInfo(self, request, context):
        stock_info_list = get_stock_info(request.country)
        response = get_stock_info_pb2.GetStockInfoResponse()
        for stock_info in stock_info_list:
            stock_info_pb = get_stock_info_pb2.StockInfo(
                country=stock_info['country'],
                symbol=stock_info['symbol'],
                name=stock_info['name'],
                full_name=stock_info['full_name'],
                isin=stock_info['isin'],
                currency=stock_info['currency'],
                stock_exchange=stock_info['stock_exchange'],
                sector=stock_info['sector'],
                industry=stock_info['industry']
            )
            response.stocks.append(stock_info_pb)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_stock_info_pb2_grpc.add_GetStockInfoServiceServicer_to_server(GetStockInfoService(), server)
    server.add_insecure_port('[::]:50404')
    server.start()
    print("GetStockInfo gRPCServer started, listening on port 50404")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

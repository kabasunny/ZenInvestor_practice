# data-analysis-python\src\get_jp_stock_info\get_jp_stock_info_grpc.py

from concurrent import futures
import grpc
import get_jp_stock_info_pb2
import get_jp_stock_info_pb2_grpc
from get_jp_stock_info_service import get_jp_stock_info

class GetJpStockInfoService(get_jp_stock_info_pb2_grpc.GetJpStockInfoServiceServicer):
    def GetJpStockInfo(self, request, context):
        stock_info_list = get_jp_stock_info()
        response = get_jp_stock_info_pb2.GetJpStockInfoResponse()
        for stock_info in stock_info_list:
            stock_info_pb = get_jp_stock_info_pb2.StockInfo(
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
    get_jp_stock_info_pb2_grpc.add_GetJpStockInfoServiceServicer_to_server(GetJpStockInfoService(), server)
    server.add_insecure_port('[::]:50402')
    server.start()
    print("GetJpStockInfo gRPCServer started, listening on port 50402")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

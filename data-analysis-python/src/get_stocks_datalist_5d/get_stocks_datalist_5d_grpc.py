# data-analysis-python\src\get_stocks_datalist_5d\get_stocks_datalist_5d_grpc.py

from concurrent import futures
import grpc
import get_stocks_datalist_5d_pb2
import get_stocks_datalist_5d_pb2_grpc
from get_stocks_datalist_5d_service import get_stocks_datalist_5d

class GetStocksDatalist5dService(get_stocks_datalist_5d_pb2_grpc.GetStocksDatalist5dServiceServicer):
    def GetStocksDatalist(self, request, context):
        stock_prices_list = get_stocks_datalist_5d(request.symbols)
        response = get_stocks_datalist_5d_pb2.GetStocksDatalist5dResponse()
        for stock_price in stock_prices_list:
            stock_price_pb = get_stocks_datalist_5d_pb2.StockPrice(
                symbol=stock_price['symbol'],
                date=stock_price['date'],
                open=stock_price['open'],
                close=stock_price['close'],
                high=stock_price['high'],
                low=stock_price['low'],
                volume=stock_price['volume']
            )
            response.stock_prices.append(stock_price_pb)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_stocks_datalist_5d_pb2_grpc.add_GetStocksDatalist5dServiceServicer_to_server(GetStocksDatalist5dService(), server)
    server.add_insecure_port('[::]:50103')
    server.start()
    print("GetStocksDatalist5d gRPCServer started, listening on port 50103")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

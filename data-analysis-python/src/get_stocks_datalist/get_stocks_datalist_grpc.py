# data-analysis-python\src\get_stocks_datalist\get_stocks_datalist_grpc.py

from concurrent import futures
import grpc
import get_stocks_datalist_pb2
import get_stocks_datalist_pb2_grpc
from get_stocks_datalist_service import get_stocks_datalist

class GetStocksDatalistService(get_stocks_datalist_pb2_grpc.GetStocksDatalistServiceServicer):
    def GetStocksDatalist(self, request, context):
        stock_prices_list = get_stocks_datalist(request.symbols)
        response = get_stocks_datalist_pb2.GetStocksDatalistResponse()
        for stock_price in stock_prices_list:
            stock_price_pb = get_stocks_datalist_pb2.StockPrice(
                symbol=stock_price['symbol'],
                date=stock_price['date'],             # 日付を追加
                open=stock_price['open'],
                close=stock_price['close'],
                high=stock_price['high'],
                low=stock_price['low'],
                volume=stock_price['volume'],
                turnover=stock_price['turnover']      # 売買代金（取引金額）を追加
            )
            response.stock_prices.append(stock_price_pb)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_stocks_datalist_pb2_grpc.add_GetStocksDatalistServiceServicer_to_server(GetStocksDatalistService(), server)
    server.add_insecure_port('[::]:50102')
    server.start()
    print("GetStocksDatalist gRPCServer started, listening on port 50102")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

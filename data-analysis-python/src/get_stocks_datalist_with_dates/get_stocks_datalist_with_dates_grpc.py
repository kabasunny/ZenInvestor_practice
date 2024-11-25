# data-analysis-python\src\get_stocks_datalist_with_dates\get_stocks_datalist_with_dates_grpc.py
from concurrent import futures
import grpc
import get_stocks_datalist_with_dates_pb2
import get_stocks_datalist_with_dates_pb2_grpc
from get_stocks_datalist_with_dates_service import get_stocks_datalist_with_dates
import time

class GetStocksDatalistWithDatesService(get_stocks_datalist_with_dates_pb2_grpc.GetStocksDatalistWithDatesServiceServicer):
    def GetStocksDatalist(self, request, context):
        
        print("gRPCサーバー : get_stocks_datalist_with_datesサービス リクエスト")

        # 処理開始時刻の記録
        start_time = time.time()

        stock_prices_list = get_stocks_datalist_with_dates(request.symbols, request.start_date, request.end_date)
        response = get_stocks_datalist_with_dates_pb2.GetStocksDatalistWithDatesResponse()
        for stock_price in stock_prices_list:
            stock_price_pb = get_stocks_datalist_with_dates_pb2.StockPrice(
                symbol=stock_price['symbol'],
                date=stock_price['date'],
                open=stock_price['open'],
                close=stock_price['close'],
                high=stock_price['high'],
                low=stock_price['low'],
                volume=stock_price['volume'],
                turnover=stock_price['turnover']      # 売買代金（取引金額）を追加
            )
            response.stock_prices.append(stock_price_pb)

        # 処理終了時刻の記録
        end_time = time.time()
        elapsed_time = end_time - start_time

        print(f"gRPCサーバー : get_stocks_datalist_with_datesサービス レスポンス - 処理時間: {elapsed_time:.2f}秒")

        return response


def serve():
    port = '50002'  # ここでポート番号を変数として定義
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_stocks_datalist_with_dates_pb2_grpc.add_GetStocksDatalistWithDatesServiceServicer_to_server(GetStocksDatalistWithDatesService(), server)
    server.add_insecure_port(f'[::]:{port}')  # 変数を使用
    server.start()
    print(f"GetStocksDatalistWithDates gRPCServer started, listening on port {port}")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

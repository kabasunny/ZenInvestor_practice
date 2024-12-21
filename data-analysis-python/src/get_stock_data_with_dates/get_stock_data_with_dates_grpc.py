# data-analysis-python\src\get_stock_data_with_dates\get_stock_data_with_dates_grpc.py
import grpc
from concurrent import futures
import time
import get_stock_data_with_dates_pb2
import get_stock_data_with_dates_pb2_grpc
from get_stock_data_with_dates_service import get_stock_data_with_dates


class GetStockDataWithDatesService(
    get_stock_data_with_dates_pb2_grpc.GetStockDataWithDatesServiceServicer
):
    def GetStockData(self, request, context):
        start_time = time.time()
        print("gRPCサーバー : get_stock_data_with_datesサービス リクエスト")
        ticker = request.ticker
        start_date = request.start_date
        end_date = request.end_date

        try:
            stock_name, stock_data_dict = get_stock_data_with_dates(
                ticker, start_date, end_date
            )

            # # デバッグログ
            # print("デバッグ: stock_data_dict =", stock_data_dict)

            stock_data = {
                str(date): get_stock_data_with_dates_pb2.StockDataWithDates(
                    open=float(values["Open"]),
                    close=float(values["Close"]),
                    high=float(values["High"]),
                    low=float(values["Low"]),
                    volume=float(values["Volume"]),
                )
                for date, values in stock_data_dict.items()
            }

            # # デバッグログ
            # print("デバッグ: stock_data =", stock_data)

            end_time = time.time()
            processing_time = end_time - start_time
            print(
                f"gRPCサーバー : get_stock_data_with_datesサービス レスポンス (処理時間: {processing_time:.2f}秒)"
            )
            return get_stock_data_with_dates_pb2.GetStockDataWithDatesResponse(
                stock_name=stock_name, stock_data=stock_data
            )
        except Exception as e:
            print(f"Error: {e}")
            context.set_details(str(e))
            context.set_code(grpc.StatusCode.INTERNAL)
            return get_stock_data_with_dates_pb2.GetStockDataWithDatesResponse()


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_stock_data_with_dates_pb2_grpc.add_GetStockDataWithDatesServiceServicer_to_server(
        GetStockDataWithDatesService(), server
    )
    server.add_insecure_port("[::]:50104")
    server.start()
    print("GetStockDataWithDates gRPCServer started, listening on port 50104")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()

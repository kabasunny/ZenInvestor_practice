# data-analysis-python\src\get_stock_info_jq\get_stock_info_jq_grpc.py
from concurrent import futures
import grpc
import get_stock_info_jq_pb2
import get_stock_info_jq_pb2_grpc
from get_stock_info_jq_service import fetch_stock_info
import time

class GetStockInfoJqService(get_stock_info_jq_pb2_grpc.GetStockInfoJqServiceServicer):
    def GetStockInfoJq(self, request, context):
        
        print("gRPCサーバー : get_stock_info_jq_service リクエスト")

        # 処理開始時刻の記録
        start_time = time.time()

        # サービス呼び出し
        stock_info_list = fetch_stock_info()

        # レスポンス作成
        response = get_stock_info_jq_pb2.GetStockInfoJqResponse()
        for index, row in stock_info_list.iterrows():
            stock_info_pb = get_stock_info_jq_pb2.StockInfo(
                ticker=row['ticker'],
                name=row['name'],
                sector=row['sector'],
                industry=row['industry'],
                date=row['date']
            )
            response.stocks.append(stock_info_pb)
        
        # 処理終了時刻の記録
        end_time = time.time()
        elapsed_time = end_time - start_time
        
        print(f"gRPCサーバー : get_stock_info_jq_service レスポンス - 処理時間: {elapsed_time:.2f}秒")
        
        # レスポンスの詳細を表示
        # for stock in response.stocks:
        #     print(f"Ticker: {stock.ticker}, Name: {stock.name}, Sector: {stock.sector}, Industry: {stock.industry}, Date: {stock.date}")
            
        return response

def serve():
    port = '50405'  # 異なるポート番号を使用
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    get_stock_info_jq_pb2_grpc.add_GetStockInfoJqServiceServicer_to_server(GetStockInfoJqService(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    print(f"GetStockInfoJq gRPC Server started, listening on port {port}")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

from concurrent import futures
import grpc
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_service import generate_chart

class ChartService(generate_chart_pb2_grpc.ChartServiceServicer):
    def GenerateChart(self, request, context):
        # リクエストから株価データと指標データを取得
        stock_data = [float(price) for price in request.stock_data]
        indicator_data = [float(value) for value in request.indicator_data]
        # チャートを生成
        chart_data = generate_chart(stock_data, indicator_data, request.ticker)
        # チャートデータをレスポンスとして返す
        return generate_chart_pb2.ChartResponse(chart_data=chart_data)

def serve():
    # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # サーバーに ChartService を追加
    generate_chart_pb2_grpc.add_ChartServiceServicer_to_server(ChartService(), server)
    # サーバーにポート 50052 を追加
    server.add_insecure_port('[::]:50052')
    # サーバーを起動
    server.start()
    # サーバー終了まで待機
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
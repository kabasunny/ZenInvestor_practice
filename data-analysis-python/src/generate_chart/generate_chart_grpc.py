from concurrent import futures
import grpc
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_service import generate_chart

class ChartService(generate_chart_pb2_grpc.ChartServiceServicer):
    def GenerateChart(self, request, context):
        chart_data = generate_chart(request.ticker)  # チャートデータを生成
        return generate_chart_pb2.ChartResponse(chart_data=chart_data)  # チャートデータをレスポンスとして返す

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))  # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
    generate_chart_pb2_grpc.add_ChartServiceServicer_to_server(ChartService(), server)  # ChartService サービスをサーバーに追加
    server.add_insecure_port('[::]:50052')  # サーバーにポート 50052 を追加
    server.start()  # サーバーを起動
    server.wait_for_termination()  # サーバー終了まで待機
    
if __name__ == '__main__':
    serve()

# data-analysis-python\src\generate_chart\generate_chart_grpc.py
import grpc
from concurrent import futures
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_service import handle_generate_chart_request

class GenerateChartService(generate_chart_pb2_grpc.GenerateChartServiceServicer):
    def GenerateChart(self, request, context):
        # サービスハンドラー関数を呼び出す
        chart_data = handle_generate_chart_request(request)

        # レスポンスを作成
        response = generate_chart_pb2.GenerateChartResponse(chart_data=chart_data)

        print("gRPCサーバーが、generate_chartサービスを呼び出し")
        return response

def serve():
    # gRPCサーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generate_chart_pb2_grpc.add_GenerateChartServiceServicer_to_server(
        GenerateChartService(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    print('GenerateChart gRPCServer started, listening on port:50052')
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

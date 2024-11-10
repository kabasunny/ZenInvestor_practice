# generate_chart_grpc.py
import grpc
from concurrent import futures
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_service import handle_generate_chart_request

class ChartGenerationService(generate_chart_pb2_grpc.ChartGenerationServiceServicer):
    def GenerateChart(self, request, context):
        # Call the service handler function
        chart_data = handle_generate_chart_request(request)

        # Create response
        response = generate_chart_pb2.GenerateChartResponse(chart_data=chart_data)

        print("gRPCサーバーが、generate_chartサービスを呼び出しました。")
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generate_chart_pb2_grpc.add_ChartGenerationServiceServicer_to_server(
        ChartGenerationService(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    print('チャート生成サーバーが起動しました。ポート:50052')
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
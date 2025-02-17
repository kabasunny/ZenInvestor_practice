# data-analysis-python\src\generate_chart\generate_chart_grpc.py
import grpc
from concurrent import futures
import generate_chart_pb2
import generate_chart_pb2_grpc
from generate_chart_service import handle_generate_chart_request
import time  # タイムスタンプを取得するためのモジュール


class GenerateChartService(generate_chart_pb2_grpc.GenerateChartServiceServicer):
    def GenerateChart(self, request, context):
        start_time = time.time()  # 開始時刻を記録
        print("gRPCサーバー : generate_chartサービス リクエスト")

        # サービスハンドラー関数を呼び出す
        chart_data = handle_generate_chart_request(request)

        # レスポンスを作成
        response = generate_chart_pb2.GenerateChartResponse(chart_data=chart_data)

        end_time = time.time()  # 終了時刻を記録
        processing_time = end_time - start_time  # 処理時間を計算
        print(
            f"gRPCサーバー : generate_chartサービス レスポンス (処理時間: {processing_time:.2f}秒)"
        )

        return response


def serve():
    # gRPCサーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generate_chart_pb2_grpc.add_GenerateChartServiceServicer_to_server(
        GenerateChartService(), server
    )
    server.add_insecure_port("[::]:50001")
    server.start()
    print("GenerateChart gRPCServer started, listening on port:50001")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()

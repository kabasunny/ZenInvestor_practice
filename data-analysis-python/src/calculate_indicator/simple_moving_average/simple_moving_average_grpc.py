# simple_moving_average_grpc.py
from concurrent import futures
import grpc
import simple_moving_average_pb2
import simple_moving_average_pb2_grpc
from simple_moving_average_service import calculate_simple_moving_average

class SimpleMovingAverageService(
    simple_moving_average_pb2_grpc.SimpleMovingAverageServiceServicer
):
    def CalculateSimpleMovingAverage(self, request, context):
        # 移動平均を計算する関数を呼び出し、結果を moving_average に格納
        moving_average = calculate_simple_moving_average(request.stock_data, request.window_size)

        # 計算結果をレスポンスとして返す
        response = simple_moving_average_pb2.SimpleMovingAverageResponse(
            moving_average=moving_average
        )

        print("gRPCサーバーが、calculate_simple_moving_averageサービスを呼び出し")
        return response

def serve():
    # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # サーバーに SimpleMovingAverageService を追加
    simple_moving_average_pb2_grpc.add_SimpleMovingAverageServiceServicer_to_server(
        SimpleMovingAverageService(), server
    )
    # サーバーにポート 50053 を追加
    server.add_insecure_port("[::]:50201")
    # サーバーを起動
    server.start()
    print("CalculateIndicater_SMA gRPCServer started, listening on port:50201")
    # サーバー終了まで待機
    server.wait_for_termination()

if __name__ == "__main__":
    serve()

from concurrent import futures
import grpc
import simple_moving_average_pb2
import simple_moving_average_pb2_grpc
from simple_moving_average_service import calculate_simple_moving_average


class SimpleMovingAverageService(
    simple_moving_average_pb2_grpc.SimpleMovingAverageServiceServicer
):
    def CalculateSimpleMovingAverage(self, request, context):
        # リクエストから株価データを取り出し、float 型に変換
        stock_data = [float(price) for price in request.stock_data]
        # 移動平均を計算する関数を呼び出し、結果を values に格納
        values = calculate_simple_moving_average(stock_data, request.window_size)
        # 計算結果をレスポンスとして返す

        print("grpcサーバーが、calculate_simple_moving_averageサービスを呼び出し")
        return simple_moving_average_pb2.SimpleMovingAverageResponse(values=values)


def serve():
    # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # サーバーに MovingAverageService を追加
    simple_moving_average_pb2_grpc.add_SimpleMovingAverageServiceServicer_to_server(
        SimpleMovingAverageService(), server
    )
    # サーバーにポート 50053 を追加
    server.add_insecure_port("[::]:50053")
    # サーバーを起動
    server.start()
    print("Server started, listening on port 50053")
    # サーバー終了まで待機
    server.wait_for_termination()


if __name__ == "__main__":
    serve()

# simple_moving_average_grpc.py
from concurrent import futures
import grpc
import simple_moving_average_pb2
import simple_moving_average_pb2_grpc
from simple_moving_average_service import calculate_simple_moving_average
import time  # タイムスタンプを取得するためのモジュール


class SimpleMovingAverageService(
    simple_moving_average_pb2_grpc.SimpleMovingAverageServiceServicer
):
    def CalculateSimpleMovingAverage(self, request, context):
        start_time = time.time()  # 開始時刻を記録
        print("gRPCサーバー : calculate_simple_moving_averageサービス リクエスト")

        # 移動平均を計算する関数を呼び出し、結果を moving_average に格納
        moving_average = calculate_simple_moving_average(
            request.stock_data, request.window_size
        )

        # 計算結果をレスポンスとして返す
        response = simple_moving_average_pb2.SimpleMovingAverageResponse(
            moving_average=moving_average
        )

        end_time = time.time()  # 終了時刻を記録
        processing_time = end_time - start_time  # 処理時間を計算
        print(
            f"gRPCサーバー : calculate_simple_moving_averageサービス レスポンス (処理時間: {processing_time:.2f}秒)"
        )

        return response


def serve():
    port = "50005"  # 異なるポート番号を使用
    # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # サーバーに SimpleMovingAverageService を追加
    simple_moving_average_pb2_grpc.add_SimpleMovingAverageServiceServicer_to_server(
        SimpleMovingAverageService(), server
    )
    # サーバーにポート 50053 を追加
    server.add_insecure_port(f"[::]: {port}")
    # サーバーを起動
    server.start()
    print(f"CalculateIndicater_SMA gRPCServer started, listening on port: {port}")
    # サーバー終了まで待機
    server.wait_for_termination()


if __name__ == "__main__":
    serve()

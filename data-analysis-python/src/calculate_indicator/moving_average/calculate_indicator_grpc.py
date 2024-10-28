from concurrent import futures
import grpc
import calculate_indicator_pb2
import calculate_indicator_pb2_grpc
from calculate_indicator_service import calculate_moving_average

class MovingAverageService(calculate_indicator_pb2_grpc.MovingAverageServiceServicer):
    def CalculateMovingAverage(self, request, context):
        # リクエストから株価データを取り出し、float 型に変換
        stock_data = [float(price) for price in request.stock_data]
        # 移動平均を計算する関数を呼び出し、結果を values に格納
        values = calculate_moving_average(stock_data, request.window_size)
        # 計算結果をレスポンスとして返す
        return calculate_indicator_pb2.MovingAverageResponse(values=values)

def serve():
    # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # サーバーに MovingAverageService を追加
    calculate_indicator_pb2_grpc.add_MovingAverageServiceServicer_to_server(MovingAverageService(), server)
    # サーバーにポート 50053 を追加
    server.add_insecure_port('[::]:50053')
    # サーバーを起動
    server.start()
    # サーバー終了まで待機
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
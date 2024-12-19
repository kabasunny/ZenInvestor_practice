# data-analysis-python\src\get_stock_data\get_stock_data_grpc.py
import grpc
from concurrent import futures
import get_stock_data_pb2  # Protocol Buffersコンパイラによって生成されるメッセージの定義を含むPythonモジュール
import get_stock_data_pb2_grpc  # Protocol Buffersコンパイラによって生成されるgRPCサービスに関連するコードを含むPythonモジュール
from get_stock_data_service import get_stock_data  # 株価データ取得関数をインポート
import time  # タイムスタンプを取得するためのモジュール


class GetStockDataService(get_stock_data_pb2_grpc.GetStockDataServiceServicer):
    def GetStockData(self, request, context):
        start_time = time.time()  # 開始時刻を記録
        print("gRPCサーバー : get_stock_dataサービス リクエスト")

        ticker = request.ticker
        period = request.period  # リクエストから期間を取得
        stock_name, stock_data_dict = get_stock_data(
            ticker, period
        )  # 数値データと銘柄名を取得

        # StockDataオブジェクトに変換
        stock_data = {
            date: get_stock_data_pb2.StockData(
                open=values["Open"],
                close=values["Close"],
                high=values["High"],
                low=values["Low"],
                volume=values["Volume"],
            )
            for date, values in stock_data_dict.items()
        }

        end_time = time.time()  # 終了時刻を記録
        processing_time = end_time - start_time  # 処理時間を計算
        print(
            f"gRPCサーバー : get_stock_dataサービス レスポンス (処理時間: {processing_time:.2f}秒)"
        )

        return get_stock_data_pb2.GetStockDataResponse(
            stock_name=stock_name, stock_data=stock_data
        )


def serve():
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10)
    )  # gRPCサーバーインスタンス
    # max_workers=10: 最大10個のワーカースレッドを作成して、リクエストを並行処理
    get_stock_data_pb2_grpc.add_GetStockDataServiceServicer_to_server(
        GetStockDataService(), server
    )  # GetStockServiceクラスとサーバーのインスタンス
    server.add_insecure_port(
        "[::]:50004"
    )  # サーバーがポート50051ですべてのIPアドレス（IPv6を含む）でリスンしリクエストを受け付けるように設定
    server.start()  # サーバー起動
    print("GetStock gRPCServer started, listening on port 50004")
    server.wait_for_termination()  # サーバーが終了するまで待機
    print("Server terminated")


if __name__ == "__main__":
    serve()

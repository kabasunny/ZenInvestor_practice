# data-analysis-python\src\generate_chart_lc_sim\generate_chart_lc_sim_grpc.py
from concurrent import futures
import grpc
import generate_chart_lc_sim_pb2
import generate_chart_lc_sim_pb2_grpc
from generate_chart_lc_sim_service import plot_stop_results
import time
import pandas as pd


class GenerateChartLCService(
    generate_chart_lc_sim_pb2_grpc.GenerateChartLCServiceServicer
):
    def GenerateChart(self, request, context):
        print("gRPCサーバー : generate_chart_lc_sim_service リクエスト")

        # 処理開始時刻の記録
        start_time = time.time()

        # リクエストデータの処理
        dates = pd.to_datetime(
            list(request.dates)
        )  # list に変換してから pd.to_datetime を使用
        close_prices = list(request.close_prices)
        data = pd.DataFrame({"Date": dates, "Close": close_prices})

        purchase_date = pd.to_datetime(request.purchase_date)
        purchase_price = request.purchase_price
        end_date = pd.to_datetime(request.end_date)
        end_price = request.end_price

        # プロット生成サービスの呼び出し
        try:
            chart_data = plot_stop_results(
                data, purchase_date, purchase_price, end_date, end_price
            )
            success = True
            message = "チャート生成に成功しました"
        except Exception as e:
            chart_data = ""
            success = False
            message = f"チャート生成に失敗しました: {e}"

        # レスポンス作成
        response = generate_chart_lc_sim_pb2.GenerateChartLCResponse(
            chart_data=chart_data, success=success, message=message
        )

        # 処理終了時刻の記録
        end_time = time.time()
        elapsed_time = end_time - start_time

        print(
            f"gRPCサーバー : generate_chart_lc_sim_service レスポンス - 処理時間: {elapsed_time:.2f}秒"
        )

        return response


def serve():
    port = "50406"  # 異なるポート番号を使用
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generate_chart_lc_sim_pb2_grpc.add_GenerateChartLCServiceServicer_to_server(
        GenerateChartLCService(), server
    )
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    print(f"GenerateChartLC gRPC Server started, listening on port {port}")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()

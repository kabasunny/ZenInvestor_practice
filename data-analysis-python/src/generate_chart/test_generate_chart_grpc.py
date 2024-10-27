import unittest  # 標準的なテストフレームワークをインポート
import grpc  # gRPC フレームワークをインポート
from concurrent import futures  # 並列実行のための futures モジュールをインポート
import generate_chart_pb2  # 生成されたプロトコルバッファ定義をインポート
import generate_chart_pb2_grpc  # 生成された gRPC サービス定義をインポート
from generate_chart_grpc import ChartService  # ChartService クラスをインポート
import base64  # バイナリデータをBASE64エンコード/デコードするためのライブラリをインポート
import os  # ファイル操作用のモジュールをインポート

class TestGenerateChartGRPC(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))  # 最大10スレッドのスレッドプールを持つ gRPC サーバーを作成
        generate_chart_pb2_grpc.add_ChartServiceServicer_to_server(ChartService(), cls.server)  # ChartService サービスをサーバーに追加
        cls.server.add_insecure_port('[::]:50052')  # サーバーにポート 50052 を追加
        cls.server.start()  # サーバーを起動

    @classmethod
    def tearDownClass(cls):
        cls.server.stop(None)  # サーバーを停止

    def test_generate_chart(self):
        with grpc.insecure_channel('localhost:50052') as channel:  # gRPC チャンネルを作成
            stub = generate_chart_pb2_grpc.ChartServiceStub(channel)  # スタブを作成
            response = stub.GenerateChart(generate_chart_pb2.ChartRequest(ticker="^GSPC"))  # リクエストを送信
            chart_data = base64.b64decode(response.chart_data)  # BASE64デコード
            output_dir = "src/generate_chart/test_output"
            if not os.path.exists(output_dir):  # 出力ディレクトリが存在しない場合は作成
                os.makedirs(output_dir)
            with open(f"{output_dir}/grpc_test_chart.png", "wb") as f:
                f.write(chart_data)  # チャートデータをファイルに書き込む
            print(f"Chart saved as {output_dir}/grpc_test_chart.png")

if __name__ == "__main__":
    unittest.main()  # テストを実行

# サーバーを起動し、クライアントとして gRPC チャンネルを作成してリクエストを送信し、レスポンスとして返された BASE64 エンコードされたチャートデータをデコードしてファイルに保存
# python -m unittest discover -s src/generate_chart  -p 'test*.py'
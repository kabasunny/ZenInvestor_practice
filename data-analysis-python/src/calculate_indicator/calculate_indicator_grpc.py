from concurrent import futures
import grpc
import calculate_indicator_pb2
import calculate_indicator_pb2_grpc
from calculate_indicator_service import calculate_moving_average

class IndicatorService(calculate_indicator_pb2_grpc.IndicatorServiceServicer):
    def CalculateMovingAverage(self, request, context):
        moving_average = calculate_moving_average(request.ticker, request.window_size)
        return calculate_indicator_pb2.IndicatorResponse(moving_average=moving_average)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    calculate_indicator_pb2_grpc.add_IndicatorServiceServicer_to_server(IndicatorService(), server)
    server.add_insecure_port('[::]:50053')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()


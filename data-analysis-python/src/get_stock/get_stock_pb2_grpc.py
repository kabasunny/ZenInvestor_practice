# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import get_stock_pb2 as get__stock__pb2

GRPC_GENERATED_VERSION = '1.67.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in get_stock_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class GetStockServiceStub(object):
    """GetStockServiceというサービスを定義
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetStockData = channel.unary_unary(
                '/GetStockService/GetStockData',
                request_serializer=get__stock__pb2.GetStockRequest.SerializeToString,
                response_deserializer=get__stock__pb2.GetStockResponse.FromString,
                _registered_method=True)


class GetStockServiceServicer(object):
    """GetStockServiceというサービスを定義
    """

    def GetStockData(self, request, context):
        """GetStockDataというRPCメソッドを定義
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GetStockServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetStockData': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStockData,
                    request_deserializer=get__stock__pb2.GetStockRequest.FromString,
                    response_serializer=get__stock__pb2.GetStockResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'GetStockService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('GetStockService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class GetStockService(object):
    """GetStockServiceというサービスを定義
    """

    @staticmethod
    def GetStockData(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/GetStockService/GetStockData',
            get__stock__pb2.GetStockRequest.SerializeToString,
            get__stock__pb2.GetStockResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

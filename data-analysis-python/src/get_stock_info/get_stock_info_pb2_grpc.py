# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import get_stock_info_pb2 as get__stock__info__pb2

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
        + f' but the generated code in get_stock_info_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class GetStockInfoServiceStub(object):
    """株式の全フィールドを取得するためのサービスを定義
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetStockInfo = channel.unary_unary(
                '/GetStockInfoService/GetStockInfo',
                request_serializer=get__stock__info__pb2.GetStockInfoRequest.SerializeToString,
                response_deserializer=get__stock__info__pb2.GetStockInfoResponse.FromString,
                _registered_method=True)


class GetStockInfoServiceServicer(object):
    """株式の全フィールドを取得するためのサービスを定義
    """

    def GetStockInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GetStockInfoServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetStockInfo': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStockInfo,
                    request_deserializer=get__stock__info__pb2.GetStockInfoRequest.FromString,
                    response_serializer=get__stock__info__pb2.GetStockInfoResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'GetStockInfoService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('GetStockInfoService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class GetStockInfoService(object):
    """株式の全フィールドを取得するためのサービスを定義
    """

    @staticmethod
    def GetStockInfo(request,
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
            '/GetStockInfoService/GetStockInfo',
            get__stock__info__pb2.GetStockInfoRequest.SerializeToString,
            get__stock__info__pb2.GetStockInfoResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

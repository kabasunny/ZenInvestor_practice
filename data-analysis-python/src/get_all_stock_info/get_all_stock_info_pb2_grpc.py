# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import get_all_stock_info_pb2 as get__all__stock__info__pb2

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
        + f' but the generated code in get_all_stock_info_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class GetAllStockInfoServiceStub(object):
    """全世界の株式の全フィールドを取得するためのサービスを定義
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetAllStockInfo = channel.unary_unary(
                '/GetAllStockInfoService/GetAllStockInfo',
                request_serializer=get__all__stock__info__pb2.GetAllStockInfoRequest.SerializeToString,
                response_deserializer=get__all__stock__info__pb2.GetAllStockInfoResponse.FromString,
                _registered_method=True)


class GetAllStockInfoServiceServicer(object):
    """全世界の株式の全フィールドを取得するためのサービスを定義
    """

    def GetAllStockInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GetAllStockInfoServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetAllStockInfo': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAllStockInfo,
                    request_deserializer=get__all__stock__info__pb2.GetAllStockInfoRequest.FromString,
                    response_serializer=get__all__stock__info__pb2.GetAllStockInfoResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'GetAllStockInfoService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('GetAllStockInfoService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class GetAllStockInfoService(object):
    """全世界の株式の全フィールドを取得するためのサービスを定義
    """

    @staticmethod
    def GetAllStockInfo(request,
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
            '/GetAllStockInfoService/GetAllStockInfo',
            get__all__stock__info__pb2.GetAllStockInfoRequest.SerializeToString,
            get__all__stock__info__pb2.GetAllStockInfoResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

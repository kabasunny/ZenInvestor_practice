# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: simple_moving_average.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'simple_moving_average.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bsimple_moving_average.proto\"E\n\x1aSimpleMovingAverageRequest\x12\x12\n\nstock_data\x18\x01 \x03(\x02\x12\x13\n\x0bwindow_size\x18\x02 \x01(\x05\"-\n\x1bSimpleMovingAverageResponse\x12\x0e\n\x06values\x18\x01 \x03(\x01\x32y\n\x1aSimpleMovingAverageService\x12[\n\x1c\x43\x61lculateSimpleMovingAverage\x12\x1b.SimpleMovingAverageRequest\x1a\x1c.SimpleMovingAverageResponse\"\x00\x42\x32Z0api-go/src/service/gateway/simple_moving_averageb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'simple_moving_average_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z0api-go/src/service/gateway/simple_moving_average'
  _globals['_SIMPLEMOVINGAVERAGEREQUEST']._serialized_start=31
  _globals['_SIMPLEMOVINGAVERAGEREQUEST']._serialized_end=100
  _globals['_SIMPLEMOVINGAVERAGERESPONSE']._serialized_start=102
  _globals['_SIMPLEMOVINGAVERAGERESPONSE']._serialized_end=147
  _globals['_SIMPLEMOVINGAVERAGESERVICE']._serialized_start=149
  _globals['_SIMPLEMOVINGAVERAGESERVICE']._serialized_end=270
# @@protoc_insertion_point(module_scope)

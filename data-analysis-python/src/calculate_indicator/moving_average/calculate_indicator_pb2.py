# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: calculate_indicator.proto
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
    'calculate_indicator.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x63\x61lculate_indicator.proto\"?\n\x14MovingAverageRequest\x12\x12\n\nstock_data\x18\x01 \x03(\x02\x12\x13\n\x0bwindow_size\x18\x02 \x01(\x05\"\'\n\x15MovingAverageResponse\x12\x0e\n\x06values\x18\x01 \x03(\x01\x32\x61\n\x14MovingAverageService\x12I\n\x16\x43\x61lculateMovingAverage\x12\x15.MovingAverageRequest\x1a\x16.MovingAverageResponse\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'calculate_indicator_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_MOVINGAVERAGEREQUEST']._serialized_start=29
  _globals['_MOVINGAVERAGEREQUEST']._serialized_end=92
  _globals['_MOVINGAVERAGERESPONSE']._serialized_start=94
  _globals['_MOVINGAVERAGERESPONSE']._serialized_end=133
  _globals['_MOVINGAVERAGESERVICE']._serialized_start=135
  _globals['_MOVINGAVERAGESERVICE']._serialized_end=232
# @@protoc_insertion_point(module_scope)
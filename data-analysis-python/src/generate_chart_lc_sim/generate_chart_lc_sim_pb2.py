# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: generate_chart_lc_sim.proto
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
    'generate_chart_lc_sim.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bgenerate_chart_lc_sim.proto\"\x91\x01\n\x16GenerateChartLCRequest\x12\r\n\x05\x64\x61tes\x18\x01 \x03(\t\x12\x14\n\x0c\x63lose_prices\x18\x02 \x03(\x01\x12\x15\n\rpurchase_date\x18\x03 \x01(\t\x12\x16\n\x0epurchase_price\x18\x04 \x01(\x01\x12\x10\n\x08\x65nd_date\x18\x05 \x01(\t\x12\x11\n\tend_price\x18\x06 \x01(\x01\"O\n\x17GenerateChartLCResponse\x12\x12\n\nchart_data\x18\x01 \x01(\t\x12\x0f\n\x07success\x18\x02 \x01(\x08\x12\x0f\n\x07message\x18\x03 \x01(\t2\\\n\x16GenerateChartLCService\x12\x42\n\rGenerateChart\x12\x17.GenerateChartLCRequest\x1a\x18.GenerateChartLCResponseB5Z3api-go/src/service/ms_gateway/generate_chart_lc_simb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'generate_chart_lc_sim_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z3api-go/src/service/ms_gateway/generate_chart_lc_sim'
  _globals['_GENERATECHARTLCREQUEST']._serialized_start=32
  _globals['_GENERATECHARTLCREQUEST']._serialized_end=177
  _globals['_GENERATECHARTLCRESPONSE']._serialized_start=179
  _globals['_GENERATECHARTLCRESPONSE']._serialized_end=258
  _globals['_GENERATECHARTLCSERVICE']._serialized_start=260
  _globals['_GENERATECHARTLCSERVICE']._serialized_end=352
# @@protoc_insertion_point(module_scope)

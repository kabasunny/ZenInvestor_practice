# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: get_oneday_alldata.proto
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
    'get_oneday_alldata.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18get_oneday_alldata.proto\"\'\n\x17GetOneDayAllDataRequest\x12\x0c\n\x04\x64\x61te\x18\x01 \x01(\t\"S\n\tStockData\x12\x0c\n\x04open\x18\x01 \x01(\x01\x12\r\n\x05\x63lose\x18\x02 \x01(\x01\x12\x0c\n\x04high\x18\x03 \x01(\x01\x12\x0b\n\x03low\x18\x04 \x01(\x01\x12\x0e\n\x06volume\x18\x05 \x01(\x01\"\x96\x01\n\x18GetOneDayAllDataResponse\x12<\n\nstock_data\x18\x01 \x03(\x0b\x32(.GetOneDayAllDataResponse.StockDataEntry\x1a<\n\x0eStockDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x19\n\x05value\x18\x02 \x01(\x0b\x32\n.StockData:\x02\x38\x01\x32\x62\n\x17GetOneDayAllDataService\x12G\n\x10GetOneDayAllData\x12\x18.GetOneDayAllDataRequest\x1a\x19.GetOneDayAllDataResponseB2Z0api-go/src/service/ms_gateway/get_oneday_alldatab\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'get_oneday_alldata_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z0api-go/src/service/ms_gateway/get_oneday_alldata'
  _globals['_GETONEDAYALLDATARESPONSE_STOCKDATAENTRY']._loaded_options = None
  _globals['_GETONEDAYALLDATARESPONSE_STOCKDATAENTRY']._serialized_options = b'8\001'
  _globals['_GETONEDAYALLDATAREQUEST']._serialized_start=28
  _globals['_GETONEDAYALLDATAREQUEST']._serialized_end=67
  _globals['_STOCKDATA']._serialized_start=69
  _globals['_STOCKDATA']._serialized_end=152
  _globals['_GETONEDAYALLDATARESPONSE']._serialized_start=155
  _globals['_GETONEDAYALLDATARESPONSE']._serialized_end=305
  _globals['_GETONEDAYALLDATARESPONSE_STOCKDATAENTRY']._serialized_start=245
  _globals['_GETONEDAYALLDATARESPONSE_STOCKDATAENTRY']._serialized_end=305
  _globals['_GETONEDAYALLDATASERVICE']._serialized_start=307
  _globals['_GETONEDAYALLDATASERVICE']._serialized_end=405
# @@protoc_insertion_point(module_scope)

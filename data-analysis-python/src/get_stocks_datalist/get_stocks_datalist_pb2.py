# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: get_stocks_datalist.proto
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
    'get_stocks_datalist.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19get_stocks_datalist.proto\"+\n\x18GetStocksDatalistRequest\x12\x0f\n\x07symbols\x18\x01 \x03(\t\"\x84\x01\n\nStockPrice\x12\x0e\n\x06symbol\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x02 \x01(\t\x12\x0c\n\x04open\x18\x03 \x01(\x01\x12\r\n\x05\x63lose\x18\x04 \x01(\x01\x12\x0c\n\x04high\x18\x05 \x01(\x01\x12\x0b\n\x03low\x18\x06 \x01(\x01\x12\x0e\n\x06volume\x18\x07 \x01(\x03\x12\x10\n\x08turnover\x18\x08 \x01(\x01\">\n\x19GetStocksDatalistResponse\x12!\n\x0cstock_prices\x18\x01 \x03(\x0b\x32\x0b.StockPrice2f\n\x18GetStocksDatalistService\x12J\n\x11GetStocksDatalist\x12\x19.GetStocksDatalistRequest\x1a\x1a.GetStocksDatalistResponseB3Z1api-go/src/service/ms_gateway/get_stocks_datalistb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'get_stocks_datalist_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z1api-go/src/service/ms_gateway/get_stocks_datalist'
  _globals['_GETSTOCKSDATALISTREQUEST']._serialized_start=29
  _globals['_GETSTOCKSDATALISTREQUEST']._serialized_end=72
  _globals['_STOCKPRICE']._serialized_start=75
  _globals['_STOCKPRICE']._serialized_end=207
  _globals['_GETSTOCKSDATALISTRESPONSE']._serialized_start=209
  _globals['_GETSTOCKSDATALISTRESPONSE']._serialized_end=271
  _globals['_GETSTOCKSDATALISTSERVICE']._serialized_start=273
  _globals['_GETSTOCKSDATALISTSERVICE']._serialized_end=375
# @@protoc_insertion_point(module_scope)

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: generate_chart.proto
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
    'generate_chart.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14generate_chart.proto\"\xba\x01\n\x14GenerateChartRequest\x12\x38\n\nstock_data\x18\x01 \x03(\x0b\x32$.GenerateChartRequest.StockDataEntry\x12\"\n\nindicators\x18\x02 \x03(\x0b\x32\x0e.IndicatorData\x1a\x44\n\x0eStockDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12!\n\x05value\x18\x02 \x01(\x0b\x32\x12.StockDataForChart:\x02\x38\x01\"[\n\x11StockDataForChart\x12\x0c\n\x04open\x18\x01 \x01(\x01\x12\r\n\x05\x63lose\x18\x02 \x01(\x01\x12\x0c\n\x04high\x18\x03 \x01(\x01\x12\x0b\n\x03low\x18\x04 \x01(\x01\x12\x0e\n\x06volume\x18\x05 \x01(\x01\"\x8d\x01\n\rIndicatorData\x12\x0c\n\x04type\x18\x01 \x01(\t\x12\x13\n\x0blegend_name\x18\x03 \x01(\t\x12*\n\x06values\x18\x02 \x03(\x0b\x32\x1a.IndicatorData.ValuesEntry\x1a-\n\x0bValuesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x01:\x02\x38\x01\"+\n\x15GenerateChartResponse\x12\x12\n\nchart_data\x18\x01 \x01(\t2V\n\x14GenerateChartService\x12>\n\rGenerateChart\x12\x15.GenerateChartRequest\x1a\x16.GenerateChartResponseB.Z,api-go/src/service/ms_gateway/generate_chartb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'generate_chart_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,api-go/src/service/ms_gateway/generate_chart'
  _globals['_GENERATECHARTREQUEST_STOCKDATAENTRY']._loaded_options = None
  _globals['_GENERATECHARTREQUEST_STOCKDATAENTRY']._serialized_options = b'8\001'
  _globals['_INDICATORDATA_VALUESENTRY']._loaded_options = None
  _globals['_INDICATORDATA_VALUESENTRY']._serialized_options = b'8\001'
  _globals['_GENERATECHARTREQUEST']._serialized_start=25
  _globals['_GENERATECHARTREQUEST']._serialized_end=211
  _globals['_GENERATECHARTREQUEST_STOCKDATAENTRY']._serialized_start=143
  _globals['_GENERATECHARTREQUEST_STOCKDATAENTRY']._serialized_end=211
  _globals['_STOCKDATAFORCHART']._serialized_start=213
  _globals['_STOCKDATAFORCHART']._serialized_end=304
  _globals['_INDICATORDATA']._serialized_start=307
  _globals['_INDICATORDATA']._serialized_end=448
  _globals['_INDICATORDATA_VALUESENTRY']._serialized_start=403
  _globals['_INDICATORDATA_VALUESENTRY']._serialized_end=448
  _globals['_GENERATECHARTRESPONSE']._serialized_start=450
  _globals['_GENERATECHARTRESPONSE']._serialized_end=493
  _globals['_GENERATECHARTSERVICE']._serialized_start=495
  _globals['_GENERATECHARTSERVICE']._serialized_end=581
# @@protoc_insertion_point(module_scope)

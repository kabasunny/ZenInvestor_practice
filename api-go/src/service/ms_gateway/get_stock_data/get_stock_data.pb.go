// get_stock_data.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: get_stock_data.proto

package get_stock_data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// リクエストメッセージを定義
type GetStockDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"` // 銘柄コードを格納するフィールド
	Period string `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"` // 期間を格納するフィールド
}

func (x *GetStockDataRequest) Reset() {
	*x = GetStockDataRequest{}
	mi := &file_get_stock_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStockDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockDataRequest) ProtoMessage() {}

func (x *GetStockDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_stock_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockDataRequest.ProtoReflect.Descriptor instead.
func (*GetStockDataRequest) Descriptor() ([]byte, []int) {
	return file_get_stock_data_proto_rawDescGZIP(), []int{0}
}

func (x *GetStockDataRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *GetStockDataRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

// 株価データを持つメッセージを定義
type StockData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open   float64 `protobuf:"fixed64,1,opt,name=open,proto3" json:"open,omitempty"`
	Close  float64 `protobuf:"fixed64,2,opt,name=close,proto3" json:"close,omitempty"`
	High   float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low    float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Volume float64 `protobuf:"fixed64,5,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *StockData) Reset() {
	*x = StockData{}
	mi := &file_get_stock_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockData) ProtoMessage() {}

func (x *StockData) ProtoReflect() protoreflect.Message {
	mi := &file_get_stock_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockData.ProtoReflect.Descriptor instead.
func (*StockData) Descriptor() ([]byte, []int) {
	return file_get_stock_data_proto_rawDescGZIP(), []int{1}
}

func (x *StockData) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *StockData) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *StockData) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *StockData) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *StockData) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

// レスポンスメッセージを定義
type GetStockDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockData map[string]*StockData `protobuf:"bytes,1,rep,name=stock_data,json=stockData,proto3" json:"stock_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 期間分の株価データを含むフィールド
	StockName string                `protobuf:"bytes,2,opt,name=stock_name,json=stockName,proto3" json:"stock_name,omitempty"`                                                                                         // 銘柄名を含むフィールド（追加）
}

func (x *GetStockDataResponse) Reset() {
	*x = GetStockDataResponse{}
	mi := &file_get_stock_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStockDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockDataResponse) ProtoMessage() {}

func (x *GetStockDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_stock_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockDataResponse.ProtoReflect.Descriptor instead.
func (*GetStockDataResponse) Descriptor() ([]byte, []int) {
	return file_get_stock_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetStockDataResponse) GetStockData() map[string]*StockData {
	if x != nil {
		return x.StockData
	}
	return nil
}

func (x *GetStockDataResponse) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

var File_get_stock_data_proto protoreflect.FileDescriptor

var file_get_stock_data_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x73, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x48, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x52, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a,
	0x2c, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6d, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67,
	0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_stock_data_proto_rawDescOnce sync.Once
	file_get_stock_data_proto_rawDescData = file_get_stock_data_proto_rawDesc
)

func file_get_stock_data_proto_rawDescGZIP() []byte {
	file_get_stock_data_proto_rawDescOnce.Do(func() {
		file_get_stock_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_stock_data_proto_rawDescData)
	})
	return file_get_stock_data_proto_rawDescData
}

var file_get_stock_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_get_stock_data_proto_goTypes = []any{
	(*GetStockDataRequest)(nil),  // 0: GetStockDataRequest
	(*StockData)(nil),            // 1: StockData
	(*GetStockDataResponse)(nil), // 2: GetStockDataResponse
	nil,                          // 3: GetStockDataResponse.StockDataEntry
}
var file_get_stock_data_proto_depIdxs = []int32{
	3, // 0: GetStockDataResponse.stock_data:type_name -> GetStockDataResponse.StockDataEntry
	1, // 1: GetStockDataResponse.StockDataEntry.value:type_name -> StockData
	0, // 2: GetStockDataService.GetStockData:input_type -> GetStockDataRequest
	2, // 3: GetStockDataService.GetStockData:output_type -> GetStockDataResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_get_stock_data_proto_init() }
func file_get_stock_data_proto_init() {
	if File_get_stock_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_stock_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_get_stock_data_proto_goTypes,
		DependencyIndexes: file_get_stock_data_proto_depIdxs,
		MessageInfos:      file_get_stock_data_proto_msgTypes,
	}.Build()
	File_get_stock_data_proto = out.File
	file_get_stock_data_proto_rawDesc = nil
	file_get_stock_data_proto_goTypes = nil
	file_get_stock_data_proto_depIdxs = nil
}

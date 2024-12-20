// data-analysis-python\src\get_stock_data_with_dates\get_stock_data_with_dates.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: get_stock_data_with_dates.proto

package get_stock_data_with_dates

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
type GetStockDataWithDatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker    string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`                        // 銘柄コードを格納するフィールド
	StartDate string `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"` // 期間を格納するフィールド
	EndDate   string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`       // 期間を格納するフィールド
}

func (x *GetStockDataWithDatesRequest) Reset() {
	*x = GetStockDataWithDatesRequest{}
	mi := &file_get_stock_data_with_dates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStockDataWithDatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockDataWithDatesRequest) ProtoMessage() {}

func (x *GetStockDataWithDatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_stock_data_with_dates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockDataWithDatesRequest.ProtoReflect.Descriptor instead.
func (*GetStockDataWithDatesRequest) Descriptor() ([]byte, []int) {
	return file_get_stock_data_with_dates_proto_rawDescGZIP(), []int{0}
}

func (x *GetStockDataWithDatesRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *GetStockDataWithDatesRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetStockDataWithDatesRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// 株価データを持つメッセージを定義
type StockDataWithDates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open   float64 `protobuf:"fixed64,1,opt,name=open,proto3" json:"open,omitempty"`
	Close  float64 `protobuf:"fixed64,2,opt,name=close,proto3" json:"close,omitempty"`
	High   float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low    float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Volume float64 `protobuf:"fixed64,5,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *StockDataWithDates) Reset() {
	*x = StockDataWithDates{}
	mi := &file_get_stock_data_with_dates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockDataWithDates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDataWithDates) ProtoMessage() {}

func (x *StockDataWithDates) ProtoReflect() protoreflect.Message {
	mi := &file_get_stock_data_with_dates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDataWithDates.ProtoReflect.Descriptor instead.
func (*StockDataWithDates) Descriptor() ([]byte, []int) {
	return file_get_stock_data_with_dates_proto_rawDescGZIP(), []int{1}
}

func (x *StockDataWithDates) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *StockDataWithDates) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *StockDataWithDates) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *StockDataWithDates) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *StockDataWithDates) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

// レスポンスメッセージを定義
type GetStockDataWithDatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockData map[string]*StockDataWithDates `protobuf:"bytes,1,rep,name=stock_data,json=stockData,proto3" json:"stock_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 期間分の株価データを含むフィールド  -> map<日付, 株価データ>
	StockName string                         `protobuf:"bytes,2,opt,name=stock_name,json=stockName,proto3" json:"stock_name,omitempty"`                                                                                         // 銘柄名を含むフィールド
}

func (x *GetStockDataWithDatesResponse) Reset() {
	*x = GetStockDataWithDatesResponse{}
	mi := &file_get_stock_data_with_dates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStockDataWithDatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockDataWithDatesResponse) ProtoMessage() {}

func (x *GetStockDataWithDatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_stock_data_with_dates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockDataWithDatesResponse.ProtoReflect.Descriptor instead.
func (*GetStockDataWithDatesResponse) Descriptor() ([]byte, []int) {
	return file_get_stock_data_with_dates_proto_rawDescGZIP(), []int{2}
}

func (x *GetStockDataWithDatesResponse) GetStockData() map[string]*StockDataWithDates {
	if x != nil {
		return x.StockData
	}
	return nil
}

func (x *GetStockDataWithDatesResponse) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

var File_get_stock_data_with_dates_proto protoreflect.FileDescriptor

var file_get_stock_data_with_dates_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x70, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x7c, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x22, 0xdf, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x51, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x6d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_stock_data_with_dates_proto_rawDescOnce sync.Once
	file_get_stock_data_with_dates_proto_rawDescData = file_get_stock_data_with_dates_proto_rawDesc
)

func file_get_stock_data_with_dates_proto_rawDescGZIP() []byte {
	file_get_stock_data_with_dates_proto_rawDescOnce.Do(func() {
		file_get_stock_data_with_dates_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_stock_data_with_dates_proto_rawDescData)
	})
	return file_get_stock_data_with_dates_proto_rawDescData
}

var file_get_stock_data_with_dates_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_get_stock_data_with_dates_proto_goTypes = []any{
	(*GetStockDataWithDatesRequest)(nil),  // 0: GetStockDataWithDatesRequest
	(*StockDataWithDates)(nil),            // 1: StockDataWithDates
	(*GetStockDataWithDatesResponse)(nil), // 2: GetStockDataWithDatesResponse
	nil,                                   // 3: GetStockDataWithDatesResponse.StockDataEntry
}
var file_get_stock_data_with_dates_proto_depIdxs = []int32{
	3, // 0: GetStockDataWithDatesResponse.stock_data:type_name -> GetStockDataWithDatesResponse.StockDataEntry
	1, // 1: GetStockDataWithDatesResponse.StockDataEntry.value:type_name -> StockDataWithDates
	0, // 2: GetStockDataWithDatesService.GetStockData:input_type -> GetStockDataWithDatesRequest
	2, // 3: GetStockDataWithDatesService.GetStockData:output_type -> GetStockDataWithDatesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_get_stock_data_with_dates_proto_init() }
func file_get_stock_data_with_dates_proto_init() {
	if File_get_stock_data_with_dates_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_stock_data_with_dates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_get_stock_data_with_dates_proto_goTypes,
		DependencyIndexes: file_get_stock_data_with_dates_proto_depIdxs,
		MessageInfos:      file_get_stock_data_with_dates_proto_msgTypes,
	}.Build()
	File_get_stock_data_with_dates_proto = out.File
	file_get_stock_data_with_dates_proto_rawDesc = nil
	file_get_stock_data_with_dates_proto_goTypes = nil
	file_get_stock_data_with_dates_proto_depIdxs = nil
}
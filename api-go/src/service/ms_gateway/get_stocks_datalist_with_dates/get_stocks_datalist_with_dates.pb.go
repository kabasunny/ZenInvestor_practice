// data-analysis-python\src\get_stocks_datalist_with_dates\get_stocks_datalist_with_dates.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: get_stocks_datalist_with_dates.proto

package get_stocks_datalist_with_dates

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

// リクエストメッセージ（銘柄コードの一覧と日付を指定）
type GetStocksDatalistWithDatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbols   []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`                      // 銘柄コードのリスト
	StartDate string   `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"` // 取得開始日
	EndDate   string   `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`       // 取得終了日
}

func (x *GetStocksDatalistWithDatesRequest) Reset() {
	*x = GetStocksDatalistWithDatesRequest{}
	mi := &file_get_stocks_datalist_with_dates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStocksDatalistWithDatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksDatalistWithDatesRequest) ProtoMessage() {}

func (x *GetStocksDatalistWithDatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_stocks_datalist_with_dates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksDatalistWithDatesRequest.ProtoReflect.Descriptor instead.
func (*GetStocksDatalistWithDatesRequest) Descriptor() ([]byte, []int) {
	return file_get_stocks_datalist_with_dates_proto_rawDescGZIP(), []int{0}
}

func (x *GetStocksDatalistWithDatesRequest) GetSymbols() []string {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *GetStocksDatalistWithDatesRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetStocksDatalistWithDatesRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// 株価情報を格納するメッセージ
type StockPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`       // 銘柄コード
	Date     string  `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`           // 日付
	Open     float64 `protobuf:"fixed64,3,opt,name=open,proto3" json:"open,omitempty"`         // 始値
	Close    float64 `protobuf:"fixed64,4,opt,name=close,proto3" json:"close,omitempty"`       // 終値
	High     float64 `protobuf:"fixed64,5,opt,name=high,proto3" json:"high,omitempty"`         // 高値
	Low      float64 `protobuf:"fixed64,6,opt,name=low,proto3" json:"low,omitempty"`           // 安値
	Volume   int64   `protobuf:"varint,7,opt,name=volume,proto3" json:"volume,omitempty"`      // 出来高
	Turnover float64 `protobuf:"fixed64,8,opt,name=turnover,proto3" json:"turnover,omitempty"` // 売買代金（取引金額）
}

func (x *StockPrice) Reset() {
	*x = StockPrice{}
	mi := &file_get_stocks_datalist_with_dates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockPrice) ProtoMessage() {}

func (x *StockPrice) ProtoReflect() protoreflect.Message {
	mi := &file_get_stocks_datalist_with_dates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockPrice.ProtoReflect.Descriptor instead.
func (*StockPrice) Descriptor() ([]byte, []int) {
	return file_get_stocks_datalist_with_dates_proto_rawDescGZIP(), []int{1}
}

func (x *StockPrice) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *StockPrice) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *StockPrice) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *StockPrice) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *StockPrice) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *StockPrice) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *StockPrice) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *StockPrice) GetTurnover() float64 {
	if x != nil {
		return x.Turnover
	}
	return 0
}

// 株価情報を格納するレスポンスメッセージ
type GetStocksDatalistWithDatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockPrices []*StockPrice `protobuf:"bytes,1,rep,name=stock_prices,json=stockPrices,proto3" json:"stock_prices,omitempty"` // 複数の株価情報
}

func (x *GetStocksDatalistWithDatesResponse) Reset() {
	*x = GetStocksDatalistWithDatesResponse{}
	mi := &file_get_stocks_datalist_with_dates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStocksDatalistWithDatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksDatalistWithDatesResponse) ProtoMessage() {}

func (x *GetStocksDatalistWithDatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_stocks_datalist_with_dates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksDatalistWithDatesResponse.ProtoReflect.Descriptor instead.
func (*GetStocksDatalistWithDatesResponse) Descriptor() ([]byte, []int) {
	return file_get_stocks_datalist_with_dates_proto_rawDescGZIP(), []int{2}
}

func (x *GetStocksDatalistWithDatesResponse) GetStockPrices() []*StockPrice {
	if x != nil {
		return x.StockPrices
	}
	return nil
}

var File_get_stocks_datalist_with_dates_proto protoreflect.FileDescriptor

var file_get_stocks_datalist_with_dates_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22,
	0xbc, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x54,
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x44, 0x61, 0x74, 0x61, 0x6c,
	0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x32, 0x81, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x22, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x44, 0x61, 0x74, 0x61, 0x6c,
	0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d,
	0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_stocks_datalist_with_dates_proto_rawDescOnce sync.Once
	file_get_stocks_datalist_with_dates_proto_rawDescData = file_get_stocks_datalist_with_dates_proto_rawDesc
)

func file_get_stocks_datalist_with_dates_proto_rawDescGZIP() []byte {
	file_get_stocks_datalist_with_dates_proto_rawDescOnce.Do(func() {
		file_get_stocks_datalist_with_dates_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_stocks_datalist_with_dates_proto_rawDescData)
	})
	return file_get_stocks_datalist_with_dates_proto_rawDescData
}

var file_get_stocks_datalist_with_dates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_get_stocks_datalist_with_dates_proto_goTypes = []any{
	(*GetStocksDatalistWithDatesRequest)(nil),  // 0: GetStocksDatalistWithDatesRequest
	(*StockPrice)(nil),                         // 1: StockPrice
	(*GetStocksDatalistWithDatesResponse)(nil), // 2: GetStocksDatalistWithDatesResponse
}
var file_get_stocks_datalist_with_dates_proto_depIdxs = []int32{
	1, // 0: GetStocksDatalistWithDatesResponse.stock_prices:type_name -> StockPrice
	0, // 1: GetStocksDatalistWithDatesService.GetStocksDatalist:input_type -> GetStocksDatalistWithDatesRequest
	2, // 2: GetStocksDatalistWithDatesService.GetStocksDatalist:output_type -> GetStocksDatalistWithDatesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_get_stocks_datalist_with_dates_proto_init() }
func file_get_stocks_datalist_with_dates_proto_init() {
	if File_get_stocks_datalist_with_dates_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_stocks_datalist_with_dates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_get_stocks_datalist_with_dates_proto_goTypes,
		DependencyIndexes: file_get_stocks_datalist_with_dates_proto_depIdxs,
		MessageInfos:      file_get_stocks_datalist_with_dates_proto_msgTypes,
	}.Build()
	File_get_stocks_datalist_with_dates_proto = out.File
	file_get_stocks_datalist_with_dates_proto_rawDesc = nil
	file_get_stocks_datalist_with_dates_proto_goTypes = nil
	file_get_stocks_datalist_with_dates_proto_depIdxs = nil
}
// generate_chart.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: generate_chart.proto

package generate_chart

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
type GenerateChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockData     map[string]*StockDataForChart `protobuf:"bytes,1,rep,name=stock_data,json=stockData,proto3" json:"stock_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 株価データを含むフィールド
	Indicators    []*IndicatorData              `protobuf:"bytes,2,rep,name=indicators,proto3" json:"indicators,omitempty"`                                                                                                        // 複数の指標データを含むフィールド
	IncludeVolume bool                          `protobuf:"varint,3,opt,name=include_volume,json=includeVolume,proto3" json:"include_volume,omitempty"`                                                                            // 出来高を含むかどうかを示すフィールドを追加
}

func (x *GenerateChartRequest) Reset() {
	*x = GenerateChartRequest{}
	mi := &file_generate_chart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChartRequest) ProtoMessage() {}

func (x *GenerateChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generate_chart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChartRequest.ProtoReflect.Descriptor instead.
func (*GenerateChartRequest) Descriptor() ([]byte, []int) {
	return file_generate_chart_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateChartRequest) GetStockData() map[string]*StockDataForChart {
	if x != nil {
		return x.StockData
	}
	return nil
}

func (x *GenerateChartRequest) GetIndicators() []*IndicatorData {
	if x != nil {
		return x.Indicators
	}
	return nil
}

func (x *GenerateChartRequest) GetIncludeVolume() bool {
	if x != nil {
		return x.IncludeVolume
	}
	return false
}

// 株価データを持つメッセージを定義
type StockDataForChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open   float64 `protobuf:"fixed64,1,opt,name=open,proto3" json:"open,omitempty"`
	Close  float64 `protobuf:"fixed64,2,opt,name=close,proto3" json:"close,omitempty"`
	High   float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low    float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Volume float64 `protobuf:"fixed64,5,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *StockDataForChart) Reset() {
	*x = StockDataForChart{}
	mi := &file_generate_chart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockDataForChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDataForChart) ProtoMessage() {}

func (x *StockDataForChart) ProtoReflect() protoreflect.Message {
	mi := &file_generate_chart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDataForChart.ProtoReflect.Descriptor instead.
func (*StockDataForChart) Descriptor() ([]byte, []int) {
	return file_generate_chart_proto_rawDescGZIP(), []int{1}
}

func (x *StockDataForChart) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *StockDataForChart) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *StockDataForChart) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *StockDataForChart) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *StockDataForChart) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

// 指標データを持つメッセージを定義
type IndicatorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string             `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`                                                                                               // 指標の種類 (SMA, EMA, RSIなど)
	LegendName string             `protobuf:"bytes,3,opt,name=legend_name,json=legendName,proto3" json:"legend_name,omitempty"`                                                                 // 凡例名 (SMA20, SMA30など) // 新しいフィールド
	Values     map[string]float64 `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"` // 指標の値を含むフィールド
}

func (x *IndicatorData) Reset() {
	*x = IndicatorData{}
	mi := &file_generate_chart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndicatorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicatorData) ProtoMessage() {}

func (x *IndicatorData) ProtoReflect() protoreflect.Message {
	mi := &file_generate_chart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicatorData.ProtoReflect.Descriptor instead.
func (*IndicatorData) Descriptor() ([]byte, []int) {
	return file_generate_chart_proto_rawDescGZIP(), []int{2}
}

func (x *IndicatorData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IndicatorData) GetLegendName() string {
	if x != nil {
		return x.LegendName
	}
	return ""
}

func (x *IndicatorData) GetValues() map[string]float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

// レスポンスメッセージを定義
type GenerateChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChartData string `protobuf:"bytes,1,opt,name=chart_data,json=chartData,proto3" json:"chart_data,omitempty"` // チャート可視化データ (Base64エンコード) を含むフィールド
}

func (x *GenerateChartResponse) Reset() {
	*x = GenerateChartResponse{}
	mi := &file_generate_chart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChartResponse) ProtoMessage() {}

func (x *GenerateChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generate_chart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChartResponse.ProtoReflect.Descriptor instead.
func (*GenerateChartResponse) Descriptor() ([]byte, []int) {
	return file_generate_chart_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateChartResponse) GetChartData() string {
	if x != nil {
		return x.ChartData
	}
	return ""
}

var File_generate_chart_proto protoreflect.FileDescriptor

var file_generate_chart_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a, 0x50, 0x0a, 0x0e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7b, 0x0a,
	0x11, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x0d, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x32, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2e, 0x5a, 0x2c, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generate_chart_proto_rawDescOnce sync.Once
	file_generate_chart_proto_rawDescData = file_generate_chart_proto_rawDesc
)

func file_generate_chart_proto_rawDescGZIP() []byte {
	file_generate_chart_proto_rawDescOnce.Do(func() {
		file_generate_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_generate_chart_proto_rawDescData)
	})
	return file_generate_chart_proto_rawDescData
}

var file_generate_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_generate_chart_proto_goTypes = []any{
	(*GenerateChartRequest)(nil),  // 0: GenerateChartRequest
	(*StockDataForChart)(nil),     // 1: StockDataForChart
	(*IndicatorData)(nil),         // 2: IndicatorData
	(*GenerateChartResponse)(nil), // 3: GenerateChartResponse
	nil,                           // 4: GenerateChartRequest.StockDataEntry
	nil,                           // 5: IndicatorData.ValuesEntry
}
var file_generate_chart_proto_depIdxs = []int32{
	4, // 0: GenerateChartRequest.stock_data:type_name -> GenerateChartRequest.StockDataEntry
	2, // 1: GenerateChartRequest.indicators:type_name -> IndicatorData
	5, // 2: IndicatorData.values:type_name -> IndicatorData.ValuesEntry
	1, // 3: GenerateChartRequest.StockDataEntry.value:type_name -> StockDataForChart
	0, // 4: GenerateChartService.GenerateChart:input_type -> GenerateChartRequest
	3, // 5: GenerateChartService.GenerateChart:output_type -> GenerateChartResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_generate_chart_proto_init() }
func file_generate_chart_proto_init() {
	if File_generate_chart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_generate_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generate_chart_proto_goTypes,
		DependencyIndexes: file_generate_chart_proto_depIdxs,
		MessageInfos:      file_generate_chart_proto_msgTypes,
	}.Build()
	File_generate_chart_proto = out.File
	file_generate_chart_proto_rawDesc = nil
	file_generate_chart_proto_goTypes = nil
	file_generate_chart_proto_depIdxs = nil
}

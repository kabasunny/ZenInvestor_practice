// data-analysis-python\src\get_holiday_info\get_trading_calendar_jq.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc3
// source: get_trading_calendar_jq.proto

package get_trading_calendar_jq_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GetTradingCalendarJqService_GetTradingCalendarJq_FullMethodName = "/GetTradingCalendarJqService/GetTradingCalendarJq"
)

// GetTradingCalendarJqServiceClient is the client API for GetTradingCalendarJqService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GetTradingCalendarJqServiceというサービスを定義
type GetTradingCalendarJqServiceClient interface {
	GetTradingCalendarJq(ctx context.Context, in *GetTradingCalendarJqRequest, opts ...grpc.CallOption) (*GetTradingCalendarJqResponse, error)
}

type getTradingCalendarJqServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetTradingCalendarJqServiceClient(cc grpc.ClientConnInterface) GetTradingCalendarJqServiceClient {
	return &getTradingCalendarJqServiceClient{cc}
}

func (c *getTradingCalendarJqServiceClient) GetTradingCalendarJq(ctx context.Context, in *GetTradingCalendarJqRequest, opts ...grpc.CallOption) (*GetTradingCalendarJqResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTradingCalendarJqResponse)
	err := c.cc.Invoke(ctx, GetTradingCalendarJqService_GetTradingCalendarJq_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTradingCalendarJqServiceServer is the server API for GetTradingCalendarJqService service.
// All implementations must embed UnimplementedGetTradingCalendarJqServiceServer
// for forward compatibility.
//
// GetTradingCalendarJqServiceというサービスを定義
type GetTradingCalendarJqServiceServer interface {
	GetTradingCalendarJq(context.Context, *GetTradingCalendarJqRequest) (*GetTradingCalendarJqResponse, error)
	mustEmbedUnimplementedGetTradingCalendarJqServiceServer()
}

// UnimplementedGetTradingCalendarJqServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetTradingCalendarJqServiceServer struct{}

func (UnimplementedGetTradingCalendarJqServiceServer) GetTradingCalendarJq(context.Context, *GetTradingCalendarJqRequest) (*GetTradingCalendarJqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradingCalendarJq not implemented")
}
func (UnimplementedGetTradingCalendarJqServiceServer) mustEmbedUnimplementedGetTradingCalendarJqServiceServer() {
}
func (UnimplementedGetTradingCalendarJqServiceServer) testEmbeddedByValue() {}

// UnsafeGetTradingCalendarJqServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetTradingCalendarJqServiceServer will
// result in compilation errors.
type UnsafeGetTradingCalendarJqServiceServer interface {
	mustEmbedUnimplementedGetTradingCalendarJqServiceServer()
}

func RegisterGetTradingCalendarJqServiceServer(s grpc.ServiceRegistrar, srv GetTradingCalendarJqServiceServer) {
	// If the following call pancis, it indicates UnimplementedGetTradingCalendarJqServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetTradingCalendarJqService_ServiceDesc, srv)
}

func _GetTradingCalendarJqService_GetTradingCalendarJq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradingCalendarJqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTradingCalendarJqServiceServer).GetTradingCalendarJq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetTradingCalendarJqService_GetTradingCalendarJq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTradingCalendarJqServiceServer).GetTradingCalendarJq(ctx, req.(*GetTradingCalendarJqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetTradingCalendarJqService_ServiceDesc is the grpc.ServiceDesc for GetTradingCalendarJqService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetTradingCalendarJqService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetTradingCalendarJqService",
	HandlerType: (*GetTradingCalendarJqServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTradingCalendarJq",
			Handler:    _GetTradingCalendarJqService_GetTradingCalendarJq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "get_trading_calendar_jq.proto",
}

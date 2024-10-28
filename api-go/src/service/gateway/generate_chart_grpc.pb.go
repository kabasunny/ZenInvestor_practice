// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: generate_chart.proto

package gateway

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
	ChartService_GenerateChart_FullMethodName = "/ChartService/GenerateChart"
)

// ChartServiceClient is the client API for ChartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChartServiceClient interface {
	GenerateChart(ctx context.Context, in *ChartRequest, opts ...grpc.CallOption) (*ChartResponse, error)
}

type chartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChartServiceClient(cc grpc.ClientConnInterface) ChartServiceClient {
	return &chartServiceClient{cc}
}

func (c *chartServiceClient) GenerateChart(ctx context.Context, in *ChartRequest, opts ...grpc.CallOption) (*ChartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChartResponse)
	err := c.cc.Invoke(ctx, ChartService_GenerateChart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChartServiceServer is the server API for ChartService service.
// All implementations must embed UnimplementedChartServiceServer
// for forward compatibility.
type ChartServiceServer interface {
	GenerateChart(context.Context, *ChartRequest) (*ChartResponse, error)
	mustEmbedUnimplementedChartServiceServer()
}

// UnimplementedChartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChartServiceServer struct{}

func (UnimplementedChartServiceServer) GenerateChart(context.Context, *ChartRequest) (*ChartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateChart not implemented")
}
func (UnimplementedChartServiceServer) mustEmbedUnimplementedChartServiceServer() {}
func (UnimplementedChartServiceServer) testEmbeddedByValue()                      {}

// UnsafeChartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChartServiceServer will
// result in compilation errors.
type UnsafeChartServiceServer interface {
	mustEmbedUnimplementedChartServiceServer()
}

func RegisterChartServiceServer(s grpc.ServiceRegistrar, srv ChartServiceServer) {
	// If the following call pancis, it indicates UnimplementedChartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChartService_ServiceDesc, srv)
}

func _ChartService_GenerateChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartServiceServer).GenerateChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChartService_GenerateChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartServiceServer).GenerateChart(ctx, req.(*ChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChartService_ServiceDesc is the grpc.ServiceDesc for ChartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChartService",
	HandlerType: (*ChartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateChart",
			Handler:    _ChartService_GenerateChart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate_chart.proto",
}

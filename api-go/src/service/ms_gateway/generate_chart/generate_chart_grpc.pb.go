// generate_chart.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: generate_chart.proto

package generate_chart

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
	GenerateChartService_GenerateChart_FullMethodName = "/GenerateChartService/GenerateChart"
)

// GenerateChartServiceClient is the client API for GenerateChartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GenerateChartServiceというサービスを定義
type GenerateChartServiceClient interface {
	GenerateChart(ctx context.Context, in *GenerateChartRequest, opts ...grpc.CallOption) (*GenerateChartResponse, error)
}

type generateChartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenerateChartServiceClient(cc grpc.ClientConnInterface) GenerateChartServiceClient {
	return &generateChartServiceClient{cc}
}

func (c *generateChartServiceClient) GenerateChart(ctx context.Context, in *GenerateChartRequest, opts ...grpc.CallOption) (*GenerateChartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateChartResponse)
	err := c.cc.Invoke(ctx, GenerateChartService_GenerateChart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateChartServiceServer is the server API for GenerateChartService service.
// All implementations must embed UnimplementedGenerateChartServiceServer
// for forward compatibility.
//
// GenerateChartServiceというサービスを定義
type GenerateChartServiceServer interface {
	GenerateChart(context.Context, *GenerateChartRequest) (*GenerateChartResponse, error)
	mustEmbedUnimplementedGenerateChartServiceServer()
}

// UnimplementedGenerateChartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGenerateChartServiceServer struct{}

func (UnimplementedGenerateChartServiceServer) GenerateChart(context.Context, *GenerateChartRequest) (*GenerateChartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateChart not implemented")
}
func (UnimplementedGenerateChartServiceServer) mustEmbedUnimplementedGenerateChartServiceServer() {}
func (UnimplementedGenerateChartServiceServer) testEmbeddedByValue()                              {}

// UnsafeGenerateChartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenerateChartServiceServer will
// result in compilation errors.
type UnsafeGenerateChartServiceServer interface {
	mustEmbedUnimplementedGenerateChartServiceServer()
}

func RegisterGenerateChartServiceServer(s grpc.ServiceRegistrar, srv GenerateChartServiceServer) {
	// If the following call pancis, it indicates UnimplementedGenerateChartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GenerateChartService_ServiceDesc, srv)
}

func _GenerateChartService_GenerateChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateChartServiceServer).GenerateChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerateChartService_GenerateChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateChartServiceServer).GenerateChart(ctx, req.(*GenerateChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenerateChartService_ServiceDesc is the grpc.ServiceDesc for GenerateChartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenerateChartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GenerateChartService",
	HandlerType: (*GenerateChartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateChart",
			Handler:    _GenerateChartService_GenerateChart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate_chart.proto",
}

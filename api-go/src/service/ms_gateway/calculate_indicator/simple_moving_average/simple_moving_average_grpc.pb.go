// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: simple_moving_average.proto

package simple_moving_average

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
	SimpleMovingAverageService_CalculateSimpleMovingAverage_FullMethodName = "/SimpleMovingAverageService/CalculateSimpleMovingAverage"
)

// SimpleMovingAverageServiceClient is the client API for SimpleMovingAverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleMovingAverageServiceClient interface {
	CalculateSimpleMovingAverage(ctx context.Context, in *SimpleMovingAverageRequest, opts ...grpc.CallOption) (*SimpleMovingAverageResponse, error)
}

type simpleMovingAverageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleMovingAverageServiceClient(cc grpc.ClientConnInterface) SimpleMovingAverageServiceClient {
	return &simpleMovingAverageServiceClient{cc}
}

func (c *simpleMovingAverageServiceClient) CalculateSimpleMovingAverage(ctx context.Context, in *SimpleMovingAverageRequest, opts ...grpc.CallOption) (*SimpleMovingAverageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SimpleMovingAverageResponse)
	err := c.cc.Invoke(ctx, SimpleMovingAverageService_CalculateSimpleMovingAverage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleMovingAverageServiceServer is the server API for SimpleMovingAverageService service.
// All implementations must embed UnimplementedSimpleMovingAverageServiceServer
// for forward compatibility.
type SimpleMovingAverageServiceServer interface {
	CalculateSimpleMovingAverage(context.Context, *SimpleMovingAverageRequest) (*SimpleMovingAverageResponse, error)
	mustEmbedUnimplementedSimpleMovingAverageServiceServer()
}

// UnimplementedSimpleMovingAverageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSimpleMovingAverageServiceServer struct{}

func (UnimplementedSimpleMovingAverageServiceServer) CalculateSimpleMovingAverage(context.Context, *SimpleMovingAverageRequest) (*SimpleMovingAverageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateSimpleMovingAverage not implemented")
}
func (UnimplementedSimpleMovingAverageServiceServer) mustEmbedUnimplementedSimpleMovingAverageServiceServer() {
}
func (UnimplementedSimpleMovingAverageServiceServer) testEmbeddedByValue() {}

// UnsafeSimpleMovingAverageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleMovingAverageServiceServer will
// result in compilation errors.
type UnsafeSimpleMovingAverageServiceServer interface {
	mustEmbedUnimplementedSimpleMovingAverageServiceServer()
}

func RegisterSimpleMovingAverageServiceServer(s grpc.ServiceRegistrar, srv SimpleMovingAverageServiceServer) {
	// If the following call pancis, it indicates UnimplementedSimpleMovingAverageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SimpleMovingAverageService_ServiceDesc, srv)
}

func _SimpleMovingAverageService_CalculateSimpleMovingAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMovingAverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleMovingAverageServiceServer).CalculateSimpleMovingAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleMovingAverageService_CalculateSimpleMovingAverage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleMovingAverageServiceServer).CalculateSimpleMovingAverage(ctx, req.(*SimpleMovingAverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleMovingAverageService_ServiceDesc is the grpc.ServiceDesc for SimpleMovingAverageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleMovingAverageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SimpleMovingAverageService",
	HandlerType: (*SimpleMovingAverageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateSimpleMovingAverage",
			Handler:    _SimpleMovingAverageService_CalculateSimpleMovingAverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple_moving_average.proto",
}
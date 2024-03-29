// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gateway/grpc_gateway.proto

package gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GateWayService_Echo_FullMethodName = "/proto.grpc_gateway.GateWayService/Echo"
)

// GateWayServiceClient is the client API for GateWayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GateWayServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type gateWayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGateWayServiceClient(cc grpc.ClientConnInterface) GateWayServiceClient {
	return &gateWayServiceClient{cc}
}

func (c *gateWayServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, GateWayService_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateWayServiceServer is the server API for GateWayService service.
// All implementations must embed UnimplementedGateWayServiceServer
// for forward compatibility
type GateWayServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedGateWayServiceServer()
}

// UnimplementedGateWayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGateWayServiceServer struct {
}

func (UnimplementedGateWayServiceServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedGateWayServiceServer) mustEmbedUnimplementedGateWayServiceServer() {}

// UnsafeGateWayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GateWayServiceServer will
// result in compilation errors.
type UnsafeGateWayServiceServer interface {
	mustEmbedUnimplementedGateWayServiceServer()
}

func RegisterGateWayServiceServer(s grpc.ServiceRegistrar, srv GateWayServiceServer) {
	s.RegisterService(&GateWayService_ServiceDesc, srv)
}

func _GateWayService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateWayServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GateWayService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateWayServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GateWayService_ServiceDesc is the grpc.ServiceDesc for GateWayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GateWayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.grpc_gateway.GateWayService",
	HandlerType: (*GateWayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _GateWayService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/grpc_gateway.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: internal/proto/EthereumService.proto

package proto

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
	EthereumService_GetAccount_FullMethodName  = "/proto.EthereumService/GetAccount"
	EthereumService_GetAccounts_FullMethodName = "/proto.EthereumService/GetAccounts"
)

// EthereumServiceClient is the client API for EthereumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EthereumServiceClient interface {
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	GetAccounts(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GetAccountsRequest, GetAccountsResponse], error)
}

type ethereumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEthereumServiceClient(cc grpc.ClientConnInterface) EthereumServiceClient {
	return &ethereumServiceClient{cc}
}

func (c *ethereumServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, EthereumService_GetAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethereumServiceClient) GetAccounts(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GetAccountsRequest, GetAccountsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EthereumService_ServiceDesc.Streams[0], EthereumService_GetAccounts_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetAccountsRequest, GetAccountsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EthereumService_GetAccountsClient = grpc.BidiStreamingClient[GetAccountsRequest, GetAccountsResponse]

// EthereumServiceServer is the server API for EthereumService service.
// All implementations must embed UnimplementedEthereumServiceServer
// for forward compatibility.
type EthereumServiceServer interface {
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	GetAccounts(grpc.BidiStreamingServer[GetAccountsRequest, GetAccountsResponse]) error
	mustEmbedUnimplementedEthereumServiceServer()
}

// UnimplementedEthereumServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEthereumServiceServer struct{}

func (UnimplementedEthereumServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedEthereumServiceServer) GetAccounts(grpc.BidiStreamingServer[GetAccountsRequest, GetAccountsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedEthereumServiceServer) mustEmbedUnimplementedEthereumServiceServer() {}
func (UnimplementedEthereumServiceServer) testEmbeddedByValue()                         {}

// UnsafeEthereumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EthereumServiceServer will
// result in compilation errors.
type UnsafeEthereumServiceServer interface {
	mustEmbedUnimplementedEthereumServiceServer()
}

func RegisterEthereumServiceServer(s grpc.ServiceRegistrar, srv EthereumServiceServer) {
	// If the following call pancis, it indicates UnimplementedEthereumServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EthereumService_ServiceDesc, srv)
}

func _EthereumService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EthereumService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthereumService_GetAccounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EthereumServiceServer).GetAccounts(&grpc.GenericServerStream[GetAccountsRequest, GetAccountsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EthereumService_GetAccountsServer = grpc.BidiStreamingServer[GetAccountsRequest, GetAccountsResponse]

// EthereumService_ServiceDesc is the grpc.ServiceDesc for EthereumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EthereumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EthereumService",
	HandlerType: (*EthereumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _EthereumService_GetAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAccounts",
			Handler:       _EthereumService_GetAccounts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/EthereumService.proto",
}
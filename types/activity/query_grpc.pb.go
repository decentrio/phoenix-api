// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: activity/query.proto

package activvity

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
	ActivityQuery_ActivityAtNft_FullMethodName = "/activity.ActivityQuery/ActivityAtNft"
)

// ActivityQueryClient is the client API for ActivityQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityQueryClient interface {
	// NftsPopular queries list popular activity
	ActivityAtNft(ctx context.Context, in *ActivityAtNftRequest, opts ...grpc.CallOption) (*ActivityAtNftResponse, error)
}

type activityQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityQueryClient(cc grpc.ClientConnInterface) ActivityQueryClient {
	return &activityQueryClient{cc}
}

func (c *activityQueryClient) ActivityAtNft(ctx context.Context, in *ActivityAtNftRequest, opts ...grpc.CallOption) (*ActivityAtNftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityAtNftResponse)
	err := c.cc.Invoke(ctx, ActivityQuery_ActivityAtNft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityQueryServer is the server API for ActivityQuery service.
// All implementations must embed UnimplementedActivityQueryServer
// for forward compatibility.
type ActivityQueryServer interface {
	// NftsPopular queries list popular activity
	ActivityAtNft(context.Context, *ActivityAtNftRequest) (*ActivityAtNftResponse, error)
	mustEmbedUnimplementedActivityQueryServer()
}

// UnimplementedActivityQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityQueryServer struct{}

func (UnimplementedActivityQueryServer) ActivityAtNft(context.Context, *ActivityAtNftRequest) (*ActivityAtNftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityAtNft not implemented")
}
func (UnimplementedActivityQueryServer) mustEmbedUnimplementedActivityQueryServer() {}
func (UnimplementedActivityQueryServer) testEmbeddedByValue()                       {}

// UnsafeActivityQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityQueryServer will
// result in compilation errors.
type UnsafeActivityQueryServer interface {
	mustEmbedUnimplementedActivityQueryServer()
}

func RegisterActivityQueryServer(s grpc.ServiceRegistrar, srv ActivityQueryServer) {
	// If the following call pancis, it indicates UnimplementedActivityQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActivityQuery_ServiceDesc, srv)
}

func _ActivityQuery_ActivityAtNft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityAtNftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityQueryServer).ActivityAtNft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityQuery_ActivityAtNft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityQueryServer).ActivityAtNft(ctx, req.(*ActivityAtNftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityQuery_ServiceDesc is the grpc.ServiceDesc for ActivityQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.ActivityQuery",
	HandlerType: (*ActivityQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivityAtNft",
			Handler:    _ActivityQuery_ActivityAtNft_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity/query.proto",
}

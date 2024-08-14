// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/shortener/v1/shortener.proto

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Shortener_Short_FullMethodName         = "/proto.shortener.v1.Shortener/Short"
	Shortener_Expand_FullMethodName        = "/proto.shortener.v1.Shortener/Expand"
	Shortener_Delete_FullMethodName        = "/proto.shortener.v1.Shortener/Delete"
	Shortener_GetUserBucket_FullMethodName = "/proto.shortener.v1.Shortener/GetUserBucket"
	Shortener_ShortBatch_FullMethodName    = "/proto.shortener.v1.Shortener/ShortBatch"
	Shortener_Stats_FullMethodName         = "/proto.shortener.v1.Shortener/Stats"
)

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	Short(ctx context.Context, in *ShortRequest, opts ...grpc.CallOption) (*ShortResponse, error)
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserBucket(ctx context.Context, in *GetUserBucketRequest, opts ...grpc.CallOption) (*GetUserBucketResponse, error)
	ShortBatch(ctx context.Context, in *ShortBatchRequest, opts ...grpc.CallOption) (*ShortBatchResponse, error)
	Stats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatsResponse, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) Short(ctx context.Context, in *ShortRequest, opts ...grpc.CallOption) (*ShortResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShortResponse)
	err := c.cc.Invoke(ctx, Shortener_Short_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpandResponse)
	err := c.cc.Invoke(ctx, Shortener_Expand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Shortener_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetUserBucket(ctx context.Context, in *GetUserBucketRequest, opts ...grpc.CallOption) (*GetUserBucketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserBucketResponse)
	err := c.cc.Invoke(ctx, Shortener_GetUserBucket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) ShortBatch(ctx context.Context, in *ShortBatchRequest, opts ...grpc.CallOption) (*ShortBatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShortBatchResponse)
	err := c.cc.Invoke(ctx, Shortener_ShortBatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) Stats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, Shortener_Stats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility.
type ShortenerServer interface {
	Short(context.Context, *ShortRequest) (*ShortResponse, error)
	Expand(context.Context, *ExpandRequest) (*ExpandResponse, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	GetUserBucket(context.Context, *GetUserBucketRequest) (*GetUserBucketResponse, error)
	ShortBatch(context.Context, *ShortBatchRequest) (*ShortBatchResponse, error)
	Stats(context.Context, *emptypb.Empty) (*StatsResponse, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShortenerServer struct{}

func (UnimplementedShortenerServer) Short(context.Context, *ShortRequest) (*ShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Short not implemented")
}
func (UnimplementedShortenerServer) Expand(context.Context, *ExpandRequest) (*ExpandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedShortenerServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedShortenerServer) GetUserBucket(context.Context, *GetUserBucketRequest) (*GetUserBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBucket not implemented")
}
func (UnimplementedShortenerServer) ShortBatch(context.Context, *ShortBatchRequest) (*ShortBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortBatch not implemented")
}
func (UnimplementedShortenerServer) Stats(context.Context, *emptypb.Empty) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}
func (UnimplementedShortenerServer) testEmbeddedByValue()                   {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	// If the following call pancis, it indicates UnimplementedShortenerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_Short_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Short(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_Short_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Short(ctx, req.(*ShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_Expand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Expand(ctx, req.(*ExpandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetUserBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetUserBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_GetUserBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetUserBucket(ctx, req.(*GetUserBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_ShortBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).ShortBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_ShortBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).ShortBatch(ctx, req.(*ShortBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_Stats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Stats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.shortener.v1.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Short",
			Handler:    _Shortener_Short_Handler,
		},
		{
			MethodName: "Expand",
			Handler:    _Shortener_Expand_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Shortener_Delete_Handler,
		},
		{
			MethodName: "GetUserBucket",
			Handler:    _Shortener_GetUserBucket_Handler,
		},
		{
			MethodName: "ShortBatch",
			Handler:    _Shortener_ShortBatch_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Shortener_Stats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shortener/v1/shortener.proto",
}
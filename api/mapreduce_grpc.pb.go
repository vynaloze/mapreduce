// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// MapReduceRegistryClient is the client API for MapReduceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapReduceRegistryClient interface {
	Register(ctx context.Context, in *MapReduceExecutable, opts ...grpc.CallOption) (*Empty, error)
}

type mapReduceRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewMapReduceRegistryClient(cc grpc.ClientConnInterface) MapReduceRegistryClient {
	return &mapReduceRegistryClient{cc}
}

func (c *mapReduceRegistryClient) Register(ctx context.Context, in *MapReduceExecutable, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/MapReduceRegistry/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapReduceRegistryServer is the server API for MapReduceRegistry service.
// All implementations must embed UnimplementedMapReduceRegistryServer
// for forward compatibility
type MapReduceRegistryServer interface {
	Register(context.Context, *MapReduceExecutable) (*Empty, error)
	mustEmbedUnimplementedMapReduceRegistryServer()
}

// UnimplementedMapReduceRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedMapReduceRegistryServer struct {
}

func (UnimplementedMapReduceRegistryServer) Register(context.Context, *MapReduceExecutable) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMapReduceRegistryServer) mustEmbedUnimplementedMapReduceRegistryServer() {}

// UnsafeMapReduceRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapReduceRegistryServer will
// result in compilation errors.
type UnsafeMapReduceRegistryServer interface {
	mustEmbedUnimplementedMapReduceRegistryServer()
}

func RegisterMapReduceRegistryServer(s grpc.ServiceRegistrar, srv MapReduceRegistryServer) {
	s.RegisterService(&MapReduceRegistry_ServiceDesc, srv)
}

func _MapReduceRegistry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapReduceExecutable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapReduceRegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MapReduceRegistry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapReduceRegistryServer).Register(ctx, req.(*MapReduceExecutable))
	}
	return interceptor(ctx, in, info, handler)
}

// MapReduceRegistry_ServiceDesc is the grpc.ServiceDesc for MapReduceRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapReduceRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MapReduceRegistry",
	HandlerType: (*MapReduceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _MapReduceRegistry_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mapreduce.proto",
}

// MapReduceClient is the client API for MapReduce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapReduceClient interface {
	Map(ctx context.Context, in *Pair, opts ...grpc.CallOption) (MapReduce_MapClient, error)
	Reduce(ctx context.Context, opts ...grpc.CallOption) (MapReduce_ReduceClient, error)
}

type mapReduceClient struct {
	cc grpc.ClientConnInterface
}

func NewMapReduceClient(cc grpc.ClientConnInterface) MapReduceClient {
	return &mapReduceClient{cc}
}

func (c *mapReduceClient) Map(ctx context.Context, in *Pair, opts ...grpc.CallOption) (MapReduce_MapClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapReduce_ServiceDesc.Streams[0], "/MapReduce/Map", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapReduceMapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapReduce_MapClient interface {
	Recv() (*Pair, error)
	grpc.ClientStream
}

type mapReduceMapClient struct {
	grpc.ClientStream
}

func (x *mapReduceMapClient) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapReduceClient) Reduce(ctx context.Context, opts ...grpc.CallOption) (MapReduce_ReduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapReduce_ServiceDesc.Streams[1], "/MapReduce/Reduce", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapReduceReduceClient{stream}
	return x, nil
}

type MapReduce_ReduceClient interface {
	Send(*Pair) error
	Recv() (*Pair, error)
	grpc.ClientStream
}

type mapReduceReduceClient struct {
	grpc.ClientStream
}

func (x *mapReduceReduceClient) Send(m *Pair) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mapReduceReduceClient) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MapReduceServer is the server API for MapReduce service.
// All implementations must embed UnimplementedMapReduceServer
// for forward compatibility
type MapReduceServer interface {
	Map(*Pair, MapReduce_MapServer) error
	Reduce(MapReduce_ReduceServer) error
	mustEmbedUnimplementedMapReduceServer()
}

// UnimplementedMapReduceServer must be embedded to have forward compatible implementations.
type UnimplementedMapReduceServer struct {
}

func (UnimplementedMapReduceServer) Map(*Pair, MapReduce_MapServer) error {
	return status.Errorf(codes.Unimplemented, "method Map not implemented")
}
func (UnimplementedMapReduceServer) Reduce(MapReduce_ReduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Reduce not implemented")
}
func (UnimplementedMapReduceServer) mustEmbedUnimplementedMapReduceServer() {}

// UnsafeMapReduceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapReduceServer will
// result in compilation errors.
type UnsafeMapReduceServer interface {
	mustEmbedUnimplementedMapReduceServer()
}

func RegisterMapReduceServer(s grpc.ServiceRegistrar, srv MapReduceServer) {
	s.RegisterService(&MapReduce_ServiceDesc, srv)
}

func _MapReduce_Map_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pair)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapReduceServer).Map(m, &mapReduceMapServer{stream})
}

type MapReduce_MapServer interface {
	Send(*Pair) error
	grpc.ServerStream
}

type mapReduceMapServer struct {
	grpc.ServerStream
}

func (x *mapReduceMapServer) Send(m *Pair) error {
	return x.ServerStream.SendMsg(m)
}

func _MapReduce_Reduce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MapReduceServer).Reduce(&mapReduceReduceServer{stream})
}

type MapReduce_ReduceServer interface {
	Send(*Pair) error
	Recv() (*Pair, error)
	grpc.ServerStream
}

type mapReduceReduceServer struct {
	grpc.ServerStream
}

func (x *mapReduceReduceServer) Send(m *Pair) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mapReduceReduceServer) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MapReduce_ServiceDesc is the grpc.ServiceDesc for MapReduce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapReduce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MapReduce",
	HandlerType: (*MapReduceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Map",
			Handler:       _MapReduce_Map_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Reduce",
			Handler:       _MapReduce_Reduce_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/mapreduce.proto",
}

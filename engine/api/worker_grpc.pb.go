// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	api "github.com/vynaloze/mapreduce/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MapWorkerClient is the client API for MapWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapWorkerClient interface {
	Map(ctx context.Context, in *MapTask, opts ...grpc.CallOption) (MapWorker_MapClient, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (MapWorker_GetClient, error)
}

type mapWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewMapWorkerClient(cc grpc.ClientConnInterface) MapWorkerClient {
	return &mapWorkerClient{cc}
}

func (c *mapWorkerClient) Map(ctx context.Context, in *MapTask, opts ...grpc.CallOption) (MapWorker_MapClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapWorker_ServiceDesc.Streams[0], "/MapWorker/Map", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapWorkerMapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapWorker_MapClient interface {
	Recv() (*Key, error)
	grpc.ClientStream
}

type mapWorkerMapClient struct {
	grpc.ClientStream
}

func (x *mapWorkerMapClient) Recv() (*Key, error) {
	m := new(Key)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapWorkerClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (MapWorker_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapWorker_ServiceDesc.Streams[1], "/MapWorker/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapWorkerGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapWorker_GetClient interface {
	Recv() (*Value, error)
	grpc.ClientStream
}

type mapWorkerGetClient struct {
	grpc.ClientStream
}

func (x *mapWorkerGetClient) Recv() (*Value, error) {
	m := new(Value)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MapWorkerServer is the server API for MapWorker service.
// All implementations must embed UnimplementedMapWorkerServer
// for forward compatibility
type MapWorkerServer interface {
	Map(*MapTask, MapWorker_MapServer) error
	Get(*Key, MapWorker_GetServer) error
	mustEmbedUnimplementedMapWorkerServer()
}

// UnimplementedMapWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedMapWorkerServer struct {
}

func (UnimplementedMapWorkerServer) Map(*MapTask, MapWorker_MapServer) error {
	return status.Errorf(codes.Unimplemented, "method Map not implemented")
}
func (UnimplementedMapWorkerServer) Get(*Key, MapWorker_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMapWorkerServer) mustEmbedUnimplementedMapWorkerServer() {}

// UnsafeMapWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapWorkerServer will
// result in compilation errors.
type UnsafeMapWorkerServer interface {
	mustEmbedUnimplementedMapWorkerServer()
}

func RegisterMapWorkerServer(s grpc.ServiceRegistrar, srv MapWorkerServer) {
	s.RegisterService(&MapWorker_ServiceDesc, srv)
}

func _MapWorker_Map_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MapTask)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapWorkerServer).Map(m, &mapWorkerMapServer{stream})
}

type MapWorker_MapServer interface {
	Send(*Key) error
	grpc.ServerStream
}

type mapWorkerMapServer struct {
	grpc.ServerStream
}

func (x *mapWorkerMapServer) Send(m *Key) error {
	return x.ServerStream.SendMsg(m)
}

func _MapWorker_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Key)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapWorkerServer).Get(m, &mapWorkerGetServer{stream})
}

type MapWorker_GetServer interface {
	Send(*Value) error
	grpc.ServerStream
}

type mapWorkerGetServer struct {
	grpc.ServerStream
}

func (x *mapWorkerGetServer) Send(m *Value) error {
	return x.ServerStream.SendMsg(m)
}

// MapWorker_ServiceDesc is the grpc.ServiceDesc for MapWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MapWorker",
	HandlerType: (*MapWorkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Map",
			Handler:       _MapWorker_Map_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Get",
			Handler:       _MapWorker_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "engine/api/worker.proto",
}

// ReduceWorkerClient is the client API for ReduceWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReduceWorkerClient interface {
	Read(ctx context.Context, opts ...grpc.CallOption) (ReduceWorker_ReadClient, error)
	Reduce(ctx context.Context, in *ReduceTask, opts ...grpc.CallOption) (*api.DFSFile, error)
}

type reduceWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewReduceWorkerClient(cc grpc.ClientConnInterface) ReduceWorkerClient {
	return &reduceWorkerClient{cc}
}

func (c *reduceWorkerClient) Read(ctx context.Context, opts ...grpc.CallOption) (ReduceWorker_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReduceWorker_ServiceDesc.Streams[0], "/ReduceWorker/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &reduceWorkerReadClient{stream}
	return x, nil
}

type ReduceWorker_ReadClient interface {
	Send(*RemoteKey) error
	CloseAndRecv() (*Done, error)
	grpc.ClientStream
}

type reduceWorkerReadClient struct {
	grpc.ClientStream
}

func (x *reduceWorkerReadClient) Send(m *RemoteKey) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reduceWorkerReadClient) CloseAndRecv() (*Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reduceWorkerClient) Reduce(ctx context.Context, in *ReduceTask, opts ...grpc.CallOption) (*api.DFSFile, error) {
	out := new(api.DFSFile)
	err := c.cc.Invoke(ctx, "/ReduceWorker/Reduce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReduceWorkerServer is the server API for ReduceWorker service.
// All implementations must embed UnimplementedReduceWorkerServer
// for forward compatibility
type ReduceWorkerServer interface {
	Read(ReduceWorker_ReadServer) error
	Reduce(context.Context, *ReduceTask) (*api.DFSFile, error)
	mustEmbedUnimplementedReduceWorkerServer()
}

// UnimplementedReduceWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedReduceWorkerServer struct {
}

func (UnimplementedReduceWorkerServer) Read(ReduceWorker_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedReduceWorkerServer) Reduce(context.Context, *ReduceTask) (*api.DFSFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reduce not implemented")
}
func (UnimplementedReduceWorkerServer) mustEmbedUnimplementedReduceWorkerServer() {}

// UnsafeReduceWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReduceWorkerServer will
// result in compilation errors.
type UnsafeReduceWorkerServer interface {
	mustEmbedUnimplementedReduceWorkerServer()
}

func RegisterReduceWorkerServer(s grpc.ServiceRegistrar, srv ReduceWorkerServer) {
	s.RegisterService(&ReduceWorker_ServiceDesc, srv)
}

func _ReduceWorker_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReduceWorkerServer).Read(&reduceWorkerReadServer{stream})
}

type ReduceWorker_ReadServer interface {
	SendAndClose(*Done) error
	Recv() (*RemoteKey, error)
	grpc.ServerStream
}

type reduceWorkerReadServer struct {
	grpc.ServerStream
}

func (x *reduceWorkerReadServer) SendAndClose(m *Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reduceWorkerReadServer) Recv() (*RemoteKey, error) {
	m := new(RemoteKey)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ReduceWorker_Reduce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReduceTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReduceWorkerServer).Reduce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReduceWorker/Reduce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReduceWorkerServer).Reduce(ctx, req.(*ReduceTask))
	}
	return interceptor(ctx, in, info, handler)
}

// ReduceWorker_ServiceDesc is the grpc.ServiceDesc for ReduceWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReduceWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReduceWorker",
	HandlerType: (*ReduceWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reduce",
			Handler:    _ReduceWorker_Reduce_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _ReduceWorker_Read_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "engine/api/worker.proto",
}
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

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterClient interface {
	Submit(ctx context.Context, in *Job, opts ...grpc.CallOption) (Master_SubmitClient, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Submit(ctx context.Context, in *Job, opts ...grpc.CallOption) (Master_SubmitClient, error) {
	stream, err := c.cc.NewStream(ctx, &Master_ServiceDesc.Streams[0], "/Master/Submit", opts...)
	if err != nil {
		return nil, err
	}
	x := &masterSubmitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Master_SubmitClient interface {
	Recv() (*JobStatus, error)
	grpc.ClientStream
}

type masterSubmitClient struct {
	grpc.ClientStream
}

func (x *masterSubmitClient) Recv() (*JobStatus, error) {
	m := new(JobStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MasterServer is the server API for Master service.
// All implementations must embed UnimplementedMasterServer
// for forward compatibility
type MasterServer interface {
	Submit(*Job, Master_SubmitServer) error
	mustEmbedUnimplementedMasterServer()
}

// UnimplementedMasterServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (UnimplementedMasterServer) Submit(*Job, Master_SubmitServer) error {
	return status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedMasterServer) mustEmbedUnimplementedMasterServer() {}

// UnsafeMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServer will
// result in compilation errors.
type UnsafeMasterServer interface {
	mustEmbedUnimplementedMasterServer()
}

func RegisterMasterServer(s grpc.ServiceRegistrar, srv MasterServer) {
	s.RegisterService(&Master_ServiceDesc, srv)
}

func _Master_Submit_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Job)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MasterServer).Submit(m, &masterSubmitServer{stream})
}

type Master_SubmitServer interface {
	Send(*JobStatus) error
	grpc.ServerStream
}

type masterSubmitServer struct {
	grpc.ServerStream
}

func (x *masterSubmitServer) Send(m *JobStatus) error {
	return x.ServerStream.SendMsg(m)
}

// Master_ServiceDesc is the grpc.ServiceDesc for Master service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Master_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Master",
	HandlerType: (*MasterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Submit",
			Handler:       _Master_Submit_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/submit.proto",
}

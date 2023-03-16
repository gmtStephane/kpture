// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: kpture/kpture.proto

package kapture

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

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	AddPacket(ctx context.Context, opts ...grpc.CallOption) (AgentService_AddPacketClient, error)
	Ready(ctx context.Context, in *Pod, opts ...grpc.CallOption) (*Empty, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) AddPacket(ctx context.Context, opts ...grpc.CallOption) (AgentService_AddPacketClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], "/service.AgentService/AddPacket", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceAddPacketClient{stream}
	return x, nil
}

type AgentService_AddPacketClient interface {
	Send(*PacketDescriptor) error
	Recv() (*Empty, error)
	grpc.ClientStream
}

type agentServiceAddPacketClient struct {
	grpc.ClientStream
}

func (x *agentServiceAddPacketClient) Send(m *PacketDescriptor) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServiceAddPacketClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) Ready(ctx context.Context, in *Pod, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.AgentService/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	AddPacket(AgentService_AddPacketServer) error
	Ready(context.Context, *Pod) (*Empty, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) AddPacket(AgentService_AddPacketServer) error {
	return status.Errorf(codes.Unimplemented, "method AddPacket not implemented")
}
func (UnimplementedAgentServiceServer) Ready(context.Context, *Pod) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_AddPacket_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).AddPacket(&agentServiceAddPacketServer{stream})
}

type AgentService_AddPacketServer interface {
	Send(*Empty) error
	Recv() (*PacketDescriptor, error)
	grpc.ServerStream
}

type agentServiceAddPacketServer struct {
	grpc.ServerStream
}

func (x *agentServiceAddPacketServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServiceAddPacketServer) Recv() (*PacketDescriptor, error) {
	m := new(PacketDescriptor)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AgentService_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AgentService/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Ready(ctx, req.(*Pod))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ready",
			Handler:    _AgentService_Ready_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddPacket",
			Handler:       _AgentService_AddPacket_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "kpture/kpture.proto",
}

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	GetPackets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_GetPacketsClient, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) GetPackets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_GetPacketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], "/service.ClientService/GetPackets", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceGetPacketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_GetPacketsClient interface {
	Recv() (*PacketDescriptor, error)
	grpc.ClientStream
}

type clientServiceGetPacketsClient struct {
	grpc.ClientStream
}

func (x *clientServiceGetPacketsClient) Recv() (*PacketDescriptor, error) {
	m := new(PacketDescriptor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	GetPackets(*Empty, ClientService_GetPacketsServer) error
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) GetPackets(*Empty, ClientService_GetPacketsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPackets not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_GetPackets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).GetPackets(m, &clientServiceGetPacketsServer{stream})
}

type ClientService_GetPacketsServer interface {
	Send(*PacketDescriptor) error
	grpc.ServerStream
}

type clientServiceGetPacketsServer struct {
	grpc.ServerStream
}

func (x *clientServiceGetPacketsServer) Send(m *PacketDescriptor) error {
	return x.ServerStream.SendMsg(m)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPackets",
			Handler:       _ClientService_GetPackets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kpture/kpture.proto",
}

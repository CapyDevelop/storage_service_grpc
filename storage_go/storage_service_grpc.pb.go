// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: storage_service.proto

package storage_go

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
	StorageService_Put_FullMethodName = "/storage_service.StorageService/Put"
)

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	// rpc Get(GetRequest) returns (GetResponse) {}
	Put(ctx context.Context, opts ...grpc.CallOption) (StorageService_PutClient, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) Put(ctx context.Context, opts ...grpc.CallOption) (StorageService_PutClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[0], StorageService_Put_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServicePutClient{stream}
	return x, nil
}

type StorageService_PutClient interface {
	Send(*PutRequest) error
	CloseAndRecv() (*PutResponse, error)
	grpc.ClientStream
}

type storageServicePutClient struct {
	grpc.ClientStream
}

func (x *storageServicePutClient) Send(m *PutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageServicePutClient) CloseAndRecv() (*PutResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	// rpc Get(GetRequest) returns (GetResponse) {}
	Put(StorageService_PutServer) error
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) Put(StorageService_PutServer) error {
	return status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_Put_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceServer).Put(&storageServicePutServer{stream})
}

type StorageService_PutServer interface {
	SendAndClose(*PutResponse) error
	Recv() (*PutRequest, error)
	grpc.ServerStream
}

type storageServicePutServer struct {
	grpc.ServerStream
}

func (x *storageServicePutServer) SendAndClose(m *PutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageServicePutServer) Recv() (*PutRequest, error) {
	m := new(PutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage_service.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Put",
			Handler:       _StorageService_Put_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage_service.proto",
}

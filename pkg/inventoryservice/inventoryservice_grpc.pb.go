//
//
// Copyright 2024 kinops authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.11
// source: inventoryservice/inventoryservice.proto

package inventoryservice

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
	InventoryManager_UnaryInventoryRequest_FullMethodName    = "/inventoryservice.InventoryManager/UnaryInventoryRequest"
	InventoryManager_ServerStreamingInventory_FullMethodName = "/inventoryservice.InventoryManager/ServerStreamingInventory"
)

// InventoryManagerClient is the client API for InventoryManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// InventoryManager is the deployment service.
type InventoryManagerClient interface {
	// UnaryInventoryRequest is unary request for deployment.
	UnaryInventoryRequest(ctx context.Context, in *InventoryRequest, opts ...grpc.CallOption) (*InventoryResponse, error)
	// ServerStreamingInventory is server side streaming of inventory updates.
	ServerStreamingInventory(ctx context.Context, in *InventoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InventoryResponse], error)
}

type inventoryManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryManagerClient(cc grpc.ClientConnInterface) InventoryManagerClient {
	return &inventoryManagerClient{cc}
}

func (c *inventoryManagerClient) UnaryInventoryRequest(ctx context.Context, in *InventoryRequest, opts ...grpc.CallOption) (*InventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InventoryResponse)
	err := c.cc.Invoke(ctx, InventoryManager_UnaryInventoryRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryManagerClient) ServerStreamingInventory(ctx context.Context, in *InventoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InventoryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &InventoryManager_ServiceDesc.Streams[0], InventoryManager_ServerStreamingInventory_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[InventoryRequest, InventoryResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type InventoryManager_ServerStreamingInventoryClient = grpc.ServerStreamingClient[InventoryResponse]

// InventoryManagerServer is the server API for InventoryManager service.
// All implementations must embed UnimplementedInventoryManagerServer
// for forward compatibility.
//
// InventoryManager is the deployment service.
type InventoryManagerServer interface {
	// UnaryInventoryRequest is unary request for deployment.
	UnaryInventoryRequest(context.Context, *InventoryRequest) (*InventoryResponse, error)
	// ServerStreamingInventory is server side streaming of inventory updates.
	ServerStreamingInventory(*InventoryRequest, grpc.ServerStreamingServer[InventoryResponse]) error
	mustEmbedUnimplementedInventoryManagerServer()
}

// UnimplementedInventoryManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInventoryManagerServer struct{}

func (UnimplementedInventoryManagerServer) UnaryInventoryRequest(context.Context, *InventoryRequest) (*InventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryInventoryRequest not implemented")
}
func (UnimplementedInventoryManagerServer) ServerStreamingInventory(*InventoryRequest, grpc.ServerStreamingServer[InventoryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingInventory not implemented")
}
func (UnimplementedInventoryManagerServer) mustEmbedUnimplementedInventoryManagerServer() {}
func (UnimplementedInventoryManagerServer) testEmbeddedByValue()                          {}

// UnsafeInventoryManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryManagerServer will
// result in compilation errors.
type UnsafeInventoryManagerServer interface {
	mustEmbedUnimplementedInventoryManagerServer()
}

func RegisterInventoryManagerServer(s grpc.ServiceRegistrar, srv InventoryManagerServer) {
	// If the following call pancis, it indicates UnimplementedInventoryManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryManager_ServiceDesc, srv)
}

func _InventoryManager_UnaryInventoryRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryManagerServer).UnaryInventoryRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryManager_UnaryInventoryRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryManagerServer).UnaryInventoryRequest(ctx, req.(*InventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryManager_ServerStreamingInventory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InventoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InventoryManagerServer).ServerStreamingInventory(m, &grpc.GenericServerStream[InventoryRequest, InventoryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type InventoryManager_ServerStreamingInventoryServer = grpc.ServerStreamingServer[InventoryResponse]

// InventoryManager_ServiceDesc is the grpc.ServiceDesc for InventoryManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventoryservice.InventoryManager",
	HandlerType: (*InventoryManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryInventoryRequest",
			Handler:    _InventoryManager_UnaryInventoryRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingInventory",
			Handler:       _InventoryManager_ServerStreamingInventory_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inventoryservice/inventoryservice.proto",
}

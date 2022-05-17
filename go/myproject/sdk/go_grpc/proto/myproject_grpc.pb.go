// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TodoRpcClient is the client API for TodoRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoRpcClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Create(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error)
	Update(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error)
	Remove(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*Nothing, error)
}

type todoRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoRpcClient(cc grpc.ClientConnInterface) TodoRpcClient {
	return &todoRpcClient{cc}
}

func (c *todoRpcClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/myproject.TodoRpc/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRpcClient) Create(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/myproject.TodoRpc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRpcClient) Update(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/myproject.TodoRpc/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRpcClient) Remove(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/myproject.TodoRpc/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoRpcServer is the server API for TodoRpc service.
// All implementations must embed UnimplementedTodoRpcServer
// for forward compatibility
type TodoRpcServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Create(context.Context, *TodoItem) (*TodoItem, error)
	Update(context.Context, *TodoItem) (*TodoItem, error)
	Remove(context.Context, *TodoItem) (*Nothing, error)
	mustEmbedUnimplementedTodoRpcServer()
}

// UnimplementedTodoRpcServer must be embedded to have forward compatible implementations.
type UnimplementedTodoRpcServer struct {
}

func (UnimplementedTodoRpcServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedTodoRpcServer) Create(context.Context, *TodoItem) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTodoRpcServer) Update(context.Context, *TodoItem) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTodoRpcServer) Remove(context.Context, *TodoItem) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedTodoRpcServer) mustEmbedUnimplementedTodoRpcServer() {}

// UnsafeTodoRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoRpcServer will
// result in compilation errors.
type UnsafeTodoRpcServer interface {
	mustEmbedUnimplementedTodoRpcServer()
}

func RegisterTodoRpcServer(s grpc.ServiceRegistrar, srv TodoRpcServer) {
	s.RegisterService(&TodoRpc_ServiceDesc, srv)
}

func _TodoRpc_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRpcServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproject.TodoRpc/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRpcServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRpc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRpcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproject.TodoRpc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRpcServer).Create(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRpc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRpcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproject.TodoRpc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRpcServer).Update(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRpc_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRpcServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproject.TodoRpc/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRpcServer).Remove(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoRpc_ServiceDesc is the grpc.ServiceDesc for TodoRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myproject.TodoRpc",
	HandlerType: (*TodoRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _TodoRpc_Search_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TodoRpc_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoRpc_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _TodoRpc_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproject.proto",
}

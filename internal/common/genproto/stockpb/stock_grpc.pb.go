// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: stockpb/stock.proto

package stockpb

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

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error)
	CheckIfItemsInStock(ctx context.Context, in *CheckIfItemsInStockRequest, opts ...grpc.CallOption) (*CheckIfItemsInStockResponse, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, "/stockpb.StockService/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) CheckIfItemsInStock(ctx context.Context, in *CheckIfItemsInStockRequest, opts ...grpc.CallOption) (*CheckIfItemsInStockResponse, error) {
	out := new(CheckIfItemsInStockResponse)
	err := c.cc.Invoke(ctx, "/stockpb.StockService/CheckIfItemsInStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations should embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error)
	CheckIfItemsInStock(context.Context, *CheckIfItemsInStockRequest) (*CheckIfItemsInStockResponse, error)
}

// UnimplementedStockServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedStockServiceServer) CheckIfItemsInStock(context.Context, *CheckIfItemsInStockRequest) (*CheckIfItemsInStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfItemsInStock not implemented")
}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stockpb.StockService/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).GetItems(ctx, req.(*GetItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_CheckIfItemsInStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfItemsInStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CheckIfItemsInStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stockpb.StockService/CheckIfItemsInStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CheckIfItemsInStock(ctx, req.(*CheckIfItemsInStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stockpb.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItems",
			Handler:    _StockService_GetItems_Handler,
		},
		{
			MethodName: "CheckIfItemsInStock",
			Handler:    _StockService_CheckIfItemsInStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stockpb/stock.proto",
}

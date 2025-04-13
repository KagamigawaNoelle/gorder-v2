package ports

import (
	"context"
	"github.com/KagamigawaNoelle/gorder-v2/common/genproto/orderpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) CreateOrder(context.Context, *orderpb.CreateOrderRequest) (*emptypb.Empty, error) {

	panic("implement me")
}

func (G GRPCServer) GetOrder(context.Context, *orderpb.GetOrderRequest) (*orderpb.Order, error) {

	panic("implement me")
}

func (G GRPCServer) UpdateOrder(context.Context, *orderpb.Order) (*emptypb.Empty, error) {

	panic(`implement me`)
}

package ports

import (
	"context"
	"github.com/KagamigawaNoelle/gorder-v2/common/genproto/stockpb"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) GetItems(context.Context, *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	panic("implement me")
}

func (G GRPCServer) CheckIfItemsInStock(context.Context, *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	panic("implement me")
}

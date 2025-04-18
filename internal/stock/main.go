package main

import (
	"context"
	"github.com/KagamigawaNoelle/gorder-v2/common/config"
	"github.com/KagamigawaNoelle/gorder-v2/common/genproto/stockpb"
	"github.com/KagamigawaNoelle/gorder-v2/common/server"
	"github.com/KagamigawaNoelle/gorder-v2/stock/ports"
	"github.com/KagamigawaNoelle/gorder-v2/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	logrus.Info(serverType)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// 暂时不用
	default:
		panic("unexpected server type")
	}
}

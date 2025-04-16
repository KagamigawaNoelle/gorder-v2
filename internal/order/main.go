package main

import (
	"context"
	"github.com/KagamigawaNoelle/gorder-v2/common/genproto/orderpb"
	"github.com/KagamigawaNoelle/gorder-v2/common/server"
	"github.com/KagamigawaNoelle/gorder-v2/order/ports"
	"github.com/KagamigawaNoelle/gorder-v2/order/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"

	"github.com/KagamigawaNoelle/gorder-v2/common/config"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serverName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := service.NewApplication(ctx)

	go server.RunGRPCServer(serverName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(app)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serverName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{app}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}

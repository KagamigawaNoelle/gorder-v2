package service

import (
	"context"
	"github.com/KagamigawaNoelle/gorder-v2/common/metrics"
	"github.com/KagamigawaNoelle/gorder-v2/order/app"
	"github.com/KagamigawaNoelle/gorder-v2/order/app/query"
	"github.com/KagamigawaNoelle/gorder-v2/order/adapters"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			query.NewGetCustomerOrderHandler(orderRepo, logger, metricClient)
		},
	}
}

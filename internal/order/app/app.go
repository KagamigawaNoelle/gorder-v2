package app

import (
	"github.com/KagamigawaNoelle/gorder-v2/order/app/command"
	"github.com/KagamigawaNoelle/gorder-v2/order/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	UpdateOrder command.UpdateOrderHandler
	CreateOrder command.CreateOrderHandler
}

type Queries struct {
	GetCustomerOrder query.GetCustomerOrderHandler
}

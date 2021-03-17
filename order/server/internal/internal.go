package internal

import (
	"context"

	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/order/ent"
)

const (
	Endpoint = "v1/orders"
)

func NewEmpty(id string) *resources.Empty {
	return &resources.Empty{
		Links: &resources.SelfLinks{
			Self: Endpoint + "/" + id,
		},
	}
}

func NewGetOrderResponse(ctx context.Context, arg *ent.Order) *services.GetOrderResponse {
	orderItems := arg.QueryItems().AllX(ctx)

	items := make([]*resources.OrderItem, len(orderItems))
	for i := range orderItems {
		items[i] = &resources.OrderItem{
			ProductId: int64(orderItems[i].ProductID),
			Qty:       int32(orderItems[i].Qty),
		}
	}
	return &services.GetOrderResponse{
		Order: &resources.Order{
			Id:     arg.ID,
			UserId: arg.UserID,
			Items:  items,
		},
		Links: &resources.SelfLinks{
			Self: Endpoint + "/" + arg.ID,
		},
	}
}

func NewListOrdersResponse(ctx context.Context, args []*ent.Order) *services.ListOrdersResponse {
	resp := make([]*services.GetOrderResponse, len(args))
	for i := range args {
		resp[i] = NewGetOrderResponse(ctx, args[i])
	}
	return &services.ListOrdersResponse{
		Orders: resp,
		Links: &resources.SelfLinks{
			Self: Endpoint,
		},
	}
}

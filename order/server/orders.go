package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/order"
	"github.com/sei-ri/microservice.io/order/ent"
	entorder "github.com/sei-ri/microservice.io/order/ent/order"
	"github.com/sei-ri/microservice.io/order/server/internal"
)

func (s *Service) CreateOrder(ctx context.Context, req *services.CreateOrderRequest) (*resources.Empty, error) {
	id := uuid.New().String()

	// TODO: SAGA
	// create order
	// checking user
	// deducting product qty
	// deducting amount
	// change order status

	items := make([]*ent.Item, len(req.Items))
	for i := range req.Items {
		item, err := s.Store.Item.Create().
			SetProductID(int(req.Items[i].ProductId)).
			SetQty(int(req.Items[i].Qty)).Save(ctx)
		if err != nil {
			return nil, err
		}
		items[i] = item
	}

	if _, err := s.Store.Order.Create().
		SetID(id).
		SetUserID(req.UserId).
		AddItems(items...).
		Save(ctx); err != nil {
		return nil, err
	}

	return internal.NewEmpty(id), nil
}

func (s *Service) GetOrder(ctx context.Context, req *services.GetOrderRequest) (*services.GetOrderResponse, error) {
	item, err := s.Store.Order.Get(ctx, req.Id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, order.ErrOrderNotFound
		}
	}
	return internal.NewGetOrderResponse(ctx, item), nil
}

func (s *Service) ListOrders(ctx context.Context, req *services.ListOrdersRequest) (*services.ListOrdersResponse, error) {
	items, err := s.Store.Order.Query().Where(entorder.UserIDEQ(req.UserId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.NewListOrdersResponse(ctx, items), nil
}

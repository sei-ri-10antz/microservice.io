package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/xid"
	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/product"
	"github.com/sei-ri/microservice.io/product/ent"
	entproduct "github.com/sei-ri/microservice.io/product/ent/product"
	"github.com/sei-ri/microservice.io/product/server/internal"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *Service) CreateProduct(ctx context.Context, req *services.CreateProductRequest) (*resources.Empty, error) {
	if req.Id == nil {
		req.Id = &wrapperspb.StringValue{Value: xid.New().String()}
	}
	return internal.NewEmptyWithID(req.Id.Value), s.EventSourcing.Dispatch(ctx, req)
}

func (s *Service) UpdateProduct(ctx context.Context, req *services.UpdateProductRequest) (*resources.Empty, error) {
	if !s.Store.Product.Query().Where(entproduct.IDEQ(req.Id)).ExistX(ctx) {
		return nil, product.ErrProductNotFound
	}

	update := s.Store.Product.UpdateOneID(req.Id)
	if v := req.Name; v != nil {
		update.SetName(v.Value)
	}
	if v := req.ImageUrl; v != nil {
		update.SetImageURL(v.Value)
	}
	if v := req.Price; v != nil {
		update.AddPrice(int(v.Value))
	}
	if v := req.Qty; v != nil {
		update.AddQty(int(v.Value))
	}

	// not es
	if _, err := update.Save(ctx); err != nil {
		return nil, err
	}
	return internal.NewEmptyWithID(req.Id), nil
}

func (s *Service) DeductProductQty(ctx context.Context, req *services.DeductProductQtyRequest) (*empty.Empty, error) {
	item, err := s.Store.Product.Get(ctx, req.Id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, product.ErrProductNotFound
		}
	}

	if item.Qty-int(req.Qty) < 0 {
		return nil, product.ErrProductQtyBalanceOut
	}

	if err := s.EventSourcing.Dispatch(ctx, req); err != nil {
		return nil, err
	}

	return internal.NewEmpty(), nil
}

func (s *Service) GetProduct(ctx context.Context, req *services.GetProductRequest) (*services.GetProductResponse, error) {
	item, err := s.Store.Product.Get(ctx, req.Id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, product.ErrProductNotFound
		}
		return nil, err
	}
	return internal.NewGetProductResponse(item), nil
}

func (s *Service) ListProducts(ctx context.Context, req *services.ListProductsRequest) (*services.ListProductsResponse, error) {
	var limit, offset int

	if limit = int(req.Limit); limit <= 0 {
		limit = 20
	}
	if offset = int(req.Offset); offset <= 0 {
		offset = 0
	}

	items, err := s.Store.Product.Query().Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.NewListProductsResponse(items), nil
}

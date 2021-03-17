package internal

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/product/ent"
)

const (
	Endpoint = "v1/products"
)

func NewEmpty() *empty.Empty {
	return &empty.Empty{}
}

func NewEmptyWithID(id interface{}) *resources.Empty {
	return &resources.Empty{
		Links: &resources.SelfLinks{
			Self: fmt.Sprintf("%s/%v", Endpoint, id),
		},
	}
}

func NewGetProductResponse(arg *ent.Product) *services.GetProductResponse {
	return &services.GetProductResponse{
		Product: &resources.Product{
			Id:        int64(arg.ID),
			Name:      arg.Name,
			ImageUrl:  arg.ImageURL,
			Price:     int32(arg.Price),
			Qty:       int32(arg.Qty),
			CreatedAt: arg.CreatedAt.Format(time.RFC3339),
			UpdatedAt: arg.UpdatedAt.Format(time.RFC3339),
		},
		Links: &resources.SelfLinks{
			Self: fmt.Sprintf("%s/%d", Endpoint, arg.ID),
		},
	}
}

func NewListProductsResponse(args []*ent.Product) *services.ListProductsResponse {
	resp := make([]*services.GetProductResponse, len(args))
	for i := range args {
		resp[i] = NewGetProductResponse(args[i])
	}
	return &services.ListProductsResponse{
		Products: resp,
		Links: &resources.SelfLinks{
			Self: Endpoint,
		},
	}
}

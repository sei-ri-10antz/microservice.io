package internal

import (
	"time"

	"github.com/sei-ri/microservice.io/account/ent"
	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
)

const (
	Endpoint = "v1/accounts"
)

func NewEmpty(id string) *resources.Empty {
	return &resources.Empty{
		Links: &resources.SelfLinks{
			Self: Endpoint + "/" + id,
		},
	}
}

func NewGetAccountResponse(arg *ent.Account) *services.GetAccountResponse {
	return &services.GetAccountResponse{
		Account: &resources.Account{
			Id:        arg.ID,
			Email:     arg.Email,
			Password:  arg.Password,
			CreatedAt: arg.CreatedAt.Format(time.RFC3339),
			UpdatedAt: arg.UpdatedAt.Format(time.RFC3339),
		},
		Links: &resources.SelfLinks{
			Self: Endpoint + "/" + arg.ID,
		},
	}
}

func NewListAccountsResponse(args []*ent.Account) *services.ListAccountsResponse {
	resp := make([]*services.GetAccountResponse, len(args))
	for i := range args {
		resp[i] = NewGetAccountResponse(args[i])
	}
	return &services.ListAccountsResponse{
		Accounts: resp,
		Links: &resources.SelfLinks{
			Self: Endpoint,
		},
	}
}

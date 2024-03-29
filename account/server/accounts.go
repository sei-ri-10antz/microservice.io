package server

import (
	"context"

	"github.com/rs/xid"
	"github.com/sei-ri/microservice.io/account"
	"github.com/sei-ri/microservice.io/account/ent"
	entaccount "github.com/sei-ri/microservice.io/account/ent/account"
	"github.com/sei-ri/microservice.io/account/server/internal"
	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *Service) CreateAccount(ctx context.Context, req *services.CreateAccountRequest) (*resources.Empty, error) {
	if s.Store.Account.Query().Where(entaccount.EmailEQ(req.Email)).ExistX(ctx) {
		return nil, account.ErrEmailAlreadyExists
	}

	if req.Id == nil {
		req.Id = &wrapperspb.StringValue{Value: xid.New().String()}
	}

	if err := s.EventSourcing.Dispatch(ctx, req); err != nil {
		return nil, err
	}

	return internal.NewEmpty(req.Id.Value), nil
}

func (s *Service) ChangePassword(ctx context.Context, req *services.ChangePasswordRequest) (*resources.Empty, error) {
	if !s.Store.Account.Query().Where(entaccount.IDEQ(req.Id)).ExistX(ctx) {
		return nil, account.ErrAccountNotFound
	}

	if err := s.EventSourcing.Dispatch(ctx, req); err != nil {
		return nil, err
	}

	return internal.NewEmpty(req.Id), nil
}

func (s *Service) GetAccount(ctx context.Context, req *services.GetAccountRequest) (*services.GetAccountResponse, error) {
	item, err := s.Store.Account.Get(ctx, req.Id)
	if err != nil {
		s.Log.Printf("GetAccount err: %v", item)
		if ent.IsNotFound(err) {
			return nil, account.ErrAccountNotFound
		}
		return nil, err
	}

	return internal.NewGetAccountResponse(item), nil
}

func (s *Service) ListAccounts(ctx context.Context, req *services.ListAccountsRequest) (*services.ListAccountsResponse, error) {
	var limit, offset int

	items, err := s.Store.Account.Query().Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		s.Log.Printf("ListAccounts err: %v", err)
		return nil, err
	}

	return internal.NewListAccountsResponse(items), nil
}

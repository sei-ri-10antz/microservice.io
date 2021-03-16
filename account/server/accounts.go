package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/sei-ri/microservice.io/account"
	"github.com/sei-ri/microservice.io/account/ent"
	entaccount "github.com/sei-ri/microservice.io/account/ent/account"
	"github.com/sei-ri/microservice.io/account/server/internal"
	"github.com/sei-ri/microservice.io/api/v1/resources"
	"github.com/sei-ri/microservice.io/api/v1/services"
)

func (s *Service) CreateAccount(ctx context.Context, req *services.CreateAccountRequest) (*resources.Empty, error) {
	if s.Store.Account.Query().Where(entaccount.EmailEQ(req.Email)).ExistX(ctx) {
		return nil, account.ErrEmailAlreadyExists
	}

	id := uuid.New().String()
	if _, err := s.Store.Account.Create().
		SetID(id).SetEmail(req.Email).
		SetPassword(req.Password).
		Save(ctx); err != nil {
		return nil, err
	}
	return internal.NewEmpty(id), nil
}

func (s *Service) ChangePassword(ctx context.Context, req *services.ChangePasswordRequest) (*resources.Empty, error) {
	if !s.Store.Account.Query().Where(entaccount.IDEQ(req.Id)).ExistX(ctx) {
		return nil, account.ErrAccountNotFound
	}

	if _, err := s.Store.Account.UpdateOneID(req.Id).
		SetPassword(req.Password).
		Save(ctx); err != nil {
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

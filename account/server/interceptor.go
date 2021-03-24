package server

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"github.com/sei-ri/microservice.io/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		now := time.Now()
		s.log.Printf("request: %s, %v", info.FullMethod, req)
		resp, err = handler(ctx, req)
		s.log.Printf("reply: %s, %v %v", info.FullMethod, resp, time.Since(now))
		return
	}
}

func (s *Server) Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := recover(); e != nil {
				debug.PrintStack()
				err = status.Errorf(codes.Internal, "Panic err: %v", e)
			}
		}()
		return handler(ctx, req)
	}
}

func (s *Server) Error() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			if errors.Is(err, account.ErrAccountNotFound) {
				err = status.Error(codes.NotFound, err.Error())
			}
			if errors.Is(err, account.ErrEmailAlreadyExists) {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return
	}
}

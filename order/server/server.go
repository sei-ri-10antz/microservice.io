package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/order/ent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Name string `envconfig:"NAME" default:"order"`
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port int    `envconfig:"PORT" default:"10012"`

	DBDriver string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUrl    string `envconfig:"DB_URL" default:"root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"`

	log *log.Logger
}

func (s *Server) Serve(ctx context.Context) {
	s.log = log.New(os.Stdout, fmt.Sprintf("[%s] ", s.Name), log.LstdFlags)

	db, err := s.openDB(ctx)
	if err != nil {
		s.log.Fatal(err)
	}
	defer db.Close()

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			s.Recovery(),
			s.Logging(),
		),
	}
	srv := grpc.NewServer(opts...)

	go func() {
		defer srv.GracefulStop()
		<-ctx.Done()
	}()

	s.openService(ctx, srv, db)

	lis, err := net.Listen("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)))
	if err != nil {
		s.log.Fatal(err)
	}
	defer lis.Close()

	s.log.Println("gRPC server listening at", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		s.log.Fatal(err)
	}
}

func (s *Server) openDB(ctx context.Context) (*ent.Client, error) {
	db, err := ent.Open(s.DBDriver, s.DBUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) openService(ctx context.Context, srv *grpc.Server, db *ent.Client) Service {
	svc := Service{
		Store: db,
		Log:   s.log,
	}

	services.RegisterOrderServiceServer(srv, &svc)
	reflection.Register(srv)

	return svc
}

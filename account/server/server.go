package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sei-ri/microservice.io/account"
	"github.com/sei-ri/microservice.io/account/ent"
	"github.com/sei-ri/microservice.io/api/v1/events"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
	"github.com/sei-ri/microservice.io/pkg/pigeon/storage/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Name string `envconfig:"NAME" default:"account"`
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port int    `envconfig:"PORT" default:"10010"`

	DBDriver string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUrl    string `envconfig:"DB_URL" default:"root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"`

	PigeonStorageDriver string `envconfig:"PIGEON_STORAGE_DRIVER" default:"mysql"`
	PigeonStorageURL    string `envconfig:"PIGEON_STORAGE_URL" default:"root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"`
	PigeonBrokerURL     string `envconfig:"PIGEON_BROKER_URL" default:"root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"`

	log *log.Logger
}

func (s *Server) Serve(ctx context.Context) {
	s.log = log.New(os.Stdout, fmt.Sprintf("[%s] ", s.Name), log.LstdFlags)

	// database
	db, err := s.openDB(ctx)
	if err != nil {
		s.log.Fatal(err)
	}
	defer db.Close()

	// event sourcing
	es, err := s.openEventSourcing(ctx, db)
	if err != nil {
		s.log.Fatal(err)
	}
	defer es.Close()

	// grpc
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			s.Recovery(),
			s.Logging(),
			s.Error(),
		),
	}

	srv := grpc.NewServer(opts...)

	go func() {
		defer srv.GracefulStop()
		<-ctx.Done()
	}()

	// services
	s.openService(ctx, srv, db, es)

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

func (s *Server) openEventSourcing(ctx context.Context, store *ent.Client) (*pigeon.Client, error) {
	storage, err := sql.New(ctx,
		sql.WithDriver(s.PigeonStorageDriver),
		sql.WithURL(s.PigeonStorageURL),
	)
	if err != nil {
		return nil, err
	}

	pg, err := pigeon.NewClient(ctx, pigeon.WithStorage(storage))
	if err != nil {
		return nil, err
	}

	pg.AddCommands(pigeon.NewSimpleCommandHandler(&account.AccountAggregate{}, pg),
		&services.CreateAccountRequest{},
		&services.ChangePasswordRequest{},
	)

	pg.AddEvents(&events.AccountCreated{}, func(e pigeon.Event) error {
		s.log.Println("event received: ", e)

		msg, ok := e.Data.(*events.AccountCreated)
		if !ok {
			return fmt.Errorf("not found event: %v", msg)
		}

		_, err := store.Account.Create().
			SetID(msg.Id).
			SetEmail(msg.Email).
			SetPassword(msg.Password).
			Save(ctx)

		return err
	})
	pg.AddEvents(&events.AccountPasswordChanged{}, func(e pigeon.Event) error {
		s.log.Println("event received: ", e)

		msg, ok := e.Data.(*events.AccountPasswordChanged)
		if !ok {
			return fmt.Errorf("not found event: %v", msg)
		}

		_, err := store.Account.UpdateOneID(msg.Id).SetPassword(msg.Password).Save(ctx)
		return err
	})

	return pg, nil
}

func (s *Server) openService(ctx context.Context, srv *grpc.Server, db *ent.Client, es *pigeon.Client) Service {
	svc := Service{
		Store:         db,
		EventSourcing: es,
		Log:           s.log,
	}

	services.RegisterAccountServiceServer(srv, &svc)
	reflection.Register(srv)

	return svc
}

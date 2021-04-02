package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sei-ri/microservice.io/api/v1/events"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
	"github.com/sei-ri/microservice.io/pkg/pigeon/broker/nats"
	"github.com/sei-ri/microservice.io/pkg/pigeon/storage/sql"
	"github.com/sei-ri/microservice.io/product"
	"github.com/sei-ri/microservice.io/product/ent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Name string `envconfig:"NAME" default:"product"`
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port int    `envconfig:"PORT" default:"10011"`

	DBDriver string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUrl    string `envconfig:"DB_URL" default:"root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"`

	PigeonStorageDriver string `envconfig:"PIGEON_STORAGE_DRIVER" default:"mysql"`
	PigeonStorageURL    string `envconfig:"PIGEON_STORAGE_URL" default:"root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"`
	PigeonBrokerURL     string `envconfig:"PIGEON_BROKER_URL" default:"nats://127.0.0.1:4222"`

	log *log.Logger
}

func (s *Server) Serve(ctx context.Context) {
	s.log = log.New(os.Stdout, fmt.Sprintf("[%s] ", s.Name), log.LstdFlags)

	db, err := s.openDB(ctx)
	if err != nil {
		s.log.Fatal(err)
	}
	defer db.Close()

	es, err := s.openEventSourcing(ctx, db)
	if err != nil {
		s.log.Fatal(err)
	}
	defer es.Close()

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
	broker, err := nats.New(ctx,
		nats.WithClientID(s.Name),
		nats.WithNatsURL(s.PigeonBrokerURL),
	)
	if err != nil {
		return nil, err
	}

	storage, err := sql.New(ctx,
		sql.WithDriver(s.PigeonStorageDriver),
		sql.WithURL(s.PigeonStorageURL),
	)
	if err != nil {
		return nil, err
	}

	pg, err := pigeon.NewClient(ctx,
		pigeon.WithStorage(storage),
		pigeon.WithBroker(broker),
	)
	if err != nil {
		return nil, err
	}

	pg.AddCommands(pigeon.NewSimpleCommandHandler(&product.ProductAggregate{}, pg),
		&services.CreateProductRequest{},
		&services.DeductProductQtyRequest{},
	)

	pg.AddEvents(&events.ProductCreated{}, func(e pigeon.Event) error {
		s.log.Println("recv event: ", e)

		msg, ok := e.Data.(*events.ProductCreated)
		if !ok {
			return fmt.Errorf("not found event: %v", msg)
		}

		_, err := store.Product.Create().Save(ctx)
		return err
	})
	pg.AddEvents(&events.ProductQtyDeducted{}, func(e pigeon.Event) error {
		s.log.Println("recv event: ", e)

		msg, ok := e.Data.(*events.ProductQtyDeducted)
		if !ok {
			return fmt.Errorf("not found event: %v", msg)
		}

		_, err := store.Product.UpdateOneID(msg.Id).AddQty(-1 * int(msg.Qty)).Save(ctx)
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

	services.RegisterProductServiceServer(srv, &svc)
	reflection.Register(srv)

	return svc
}

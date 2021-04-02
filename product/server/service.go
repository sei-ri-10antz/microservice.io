package server

import (
	"log"

	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
	"github.com/sei-ri/microservice.io/product/ent"
)

var _ services.ProductServiceServer = &Service{}

type Service struct {
	Store         *ent.Client
	EventSourcing *pigeon.Client
	Log           *log.Logger
}

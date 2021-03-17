package server

import (
	"log"

	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/order/ent"
)

var _ services.OrderServiceServer = &Service{}

type Service struct {
	Store *ent.Client
	Log   *log.Logger
}

package server

import (
	"log"

	"github.com/sei-ri/microservice.io/account/ent"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
)

type Service struct {
	Store         *ent.Client
	EventSourcing *pigeon.Client
	Log           *log.Logger
}

package server

import (
	"log"

	"github.com/sei-ri/microservice.io/account/ent"
)

type Service struct {
	Store *ent.Client
	Log   *log.Logger
}

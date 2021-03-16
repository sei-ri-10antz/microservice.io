package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/sei-ri/microservice.io/account/ent"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DefaultDriver = "mysql"
	DefaultURL    = "root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=true"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1000, "Make fake data number")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	db, err := ent.Open(DefaultDriver, DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// clean
	if _, err := db.Account.Delete().Exec(ctx); err != nil {
		log.Fatal(err)
	}

	accounts := make([]*ent.AccountCreate, n)
	for i := 0; i < n; i++ {
		accounts[i] = db.Account.Create().
			SetID(uuid.New().String()).
			SetEmail(faker.Email()).
			SetPassword(faker.Password())
	}
	if _, err := db.Account.CreateBulk(accounts...).Save(ctx); err != nil {
		log.Fatal(err)
	}
}

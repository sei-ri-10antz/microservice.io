package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/sei-ri/microservice.io/product/ent"

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

	if _, err := db.Product.Delete().Exec(ctx); err != nil {
		log.Fatal(err)
	}

	products := make([]*ent.ProductCreate, n)
	for i := 0; i < n; i++ {
		products[i] = db.Product.Create().
			SetName(faker.Name()).
			SetImageURL(faker.URL()).
			SetPrice(RangeRandom(10, 30)).
			SetQty(RangeRandom(1, 10))
	}

	if _, err := db.Product.CreateBulk(products...).Save(ctx); err != nil {
		log.Fatal(err)
	}
}

func RangeRandom(min, max int) int {
	return rand.Intn(max-min) + min
}

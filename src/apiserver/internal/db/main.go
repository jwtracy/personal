package main

import (
	"context"
	"log"

	"github.com/johnwtracy/personal/src/apiserver/internal/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open(
		"postgres",
		"host=localhost port=5432 user=johnwtracy dbname=personaldb password=password123 sslmode=disable",
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

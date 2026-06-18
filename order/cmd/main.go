package main

import (
	"log"

	"github.com/leonardoaraujodf/e-commerce/order/config"
	"github.com/leonardoaraujodf/e-commerce/order/internal/adapters/db"
	"github.com/leonardoaraujodf/e-commerce/order/internal/adapters/grpc"
	"github.com/leonardoaraujodf/e-commerce/order/internal/application/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}

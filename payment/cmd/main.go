package main

import (
	"github.com/yifeistudio-developer/patrol/payment/config"
	"github.com/yifeistudio-developer/patrol/payment/internal/adapters/db"
	"github.com/yifeistudio-developer/patrol/payment/internal/adapters/grpc"
	"github.com/yifeistudio-developer/patrol/payment/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}

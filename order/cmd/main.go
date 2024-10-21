package main

import (
	"github.com/yifeistudio-developer/patrol/order/config"
	"github.com/yifeistudio-developer/patrol/order/internal/adapters/db"
	"github.com/yifeistudio-developer/patrol/order/internal/adapters/grpc"
	"github.com/yifeistudio-developer/patrol/order/internal/adapters/payment"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}

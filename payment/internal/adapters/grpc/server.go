package grpc

import (
	"fmt"
	"github.com/yifeistudio-developer/patrol/payment/config"
	"github.com/yifeistudio-developer/patrol/payment/internal/ports"
	"github.com/yifeistudio-developer/wharf/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api  ports.APIPort
	port int
	payment.UnimplementedPaymentServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServer(grpcServer, a)
	// configuration
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

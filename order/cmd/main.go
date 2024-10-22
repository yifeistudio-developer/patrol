package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/sony/gobreaker"
	"github.com/yifeistudio-developer/patrol/order/config"
	"github.com/yifeistudio-developer/patrol/order/internal/adapters/db"
	"github.com/yifeistudio-developer/patrol/order/internal/adapters/grpc"
	"github.com/yifeistudio-developer/patrol/order/internal/adapters/payment"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/api"
	"google.golang.org/grpc/credentials"
	"log"
	"math/rand"
	"os"
	"time"
)

var cb *gobreaker.CircuitBreaker

func circuitBreaker() {
	cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "demo",
		MaxRequests: 3,
		Timeout:     4 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("State change %s from %s to %s", name, from, to)
		},
	})
	execute, err := cb.Execute(func() (interface{}, error) {
		res, isErr := isError()
		if isErr {
			return nil, errors.New("error")
		}
		return res, nil
	})
	if err != nil {
		log.Fatalf("Circuit breaker error %v", err)
	} else {
		log.Printf("Circuirt breaker result: %v", execute)
	}
}

func isError() (int, bool) {
	min := 10
	max := 30
	result := rand.Intn(max-min) + min
	return result, result != 20
}

func getTlsCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("", "")
	if err != nil {
		// handle load error
	}
	certPool := x509.NewCertPool()
	caCert, err := os.ReadFile("")
	if err != nil {
		// handle read file error
	}
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to append ca cert")
	}
	return credentials.NewTLS(
		&tls.Config{
			ClientAuth:   tls.RequireAnyClientCert,
			Certificates: []tls.Certificate{serverCert},
			ClientCAs:    certPool,
		}), nil
}

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

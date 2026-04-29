package main

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-merchant_award/apps"
	"github.com/MamangRust/microservice-ecommerce-pkg/server"
)

func main() {
	srv, err := apps.NewServer(&server.Config{
		ServiceName:    "merchant_award-service",
		ServiceVersion: "1.0.0",
		Environment:    "production",
		OtelEndpoint:   "otel-collector:4317",
		Port:           50065,
		DBCluster:      "MERCHANT_AWARD",
		MigrationPath:  "./database/migrations",
	})

	if err != nil {
		panic(err)
	}

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
